package tachibana_grpc_server

import (
	"context"
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
	"sort"
	"sync"
)

type iStreamService interface {
	connect(ctx context.Context, session *accountSession, clientToken string, req *pb.StreamRequest, stream pb.TachibanaService_StreamServer) error
}

type streamService struct {
	tachibanaApi   iTachibanaApi
	sessionStreams map[string]*sessionStream
}

func (s *streamService) connect(ctx context.Context, session *accountSession, clientToken string, req *pb.StreamRequest, stream pb.TachibanaService_StreamServer) error {
	if _, ok := s.sessionStreams[session.Token]; !ok {
		cCtx, cf := context.WithCancel(context.Background())
		s.sessionStreams[session.Token] = &sessionStream{
			sessionToken: session.Token,
			streams:      map[string]iClientStream{},
			ctx:          cCtx,
			cf:           cf,
		}
	}
	ss := s.sessionStreams[session.Token]

	cErrCh := make(chan error)
	sReq := ss.addClient(ctx, clientToken, req, stream, cErrCh)

	// 切断・接続
	cCtx, cf := context.WithCancel(ctx)
	resCh, sErrCh := s.tachibanaApi.stream(cCtx, session.Session, sReq)
	ss.start(cCtx, cf, resCh, sErrCh)

	return <-cErrCh
}

type sessionStream struct {
	sessionToken string
	streams      map[string]iClientStream
	isConnected  bool
	resCh        <-chan *pb.StreamResponse
	errCh        <-chan error
	ctx          context.Context
	cf           context.CancelFunc
	mtx          sync.RWMutex
}

func (s *sessionStream) addClient(ctx context.Context, clientToken string, req *pb.StreamRequest, stream pb.TachibanaService_StreamServer, errCh chan error) *pb.StreamRequest {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.streams == nil {
		s.streams = map[string]iClientStream{}
	}
	if cli, ok := s.streams[clientToken]; ok {
		cli.disconnect(newStreamRequestErr)
	}
	s.streams[clientToken] = &clientStream{
		clientToken: clientToken,
		request:     req,
		stream:      stream,
		ctx:         ctx,
		errCh:       errCh,
	}
	s.streams[clientToken].start()

	return s.getRequest()
}

func (s *sessionStream) getRequest() *pb.StreamRequest {
	req := &pb.StreamRequest{
		EventTypes:   []pb.EventType{},
		StreamIssues: []*pb.StreamIssue{},
	}

	for _, cs := range s.streams {
		req = req.Union(cs.getRequest())
	}

	// 冪等性のため並び替えておく
	sort.Slice(req.EventTypes, func(i, j int) bool {
		return req.EventTypes[i] < req.EventTypes[j]
	})
	sort.Slice(req.StreamIssues, func(i, j int) bool {
		return req.StreamIssues[i].IssueCode < req.StreamIssues[j].IssueCode ||
			(req.StreamIssues[i].IssueCode == req.StreamIssues[j].IssueCode && req.StreamIssues[i].Exchange < req.StreamIssues[j].Exchange)
	})

	return req
}

func (s *sessionStream) start(ctx context.Context, cf context.CancelFunc, resCh <-chan *pb.StreamResponse, errCh <-chan error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.isConnected {
		s.disconnect()
	}

	s.ctx = ctx
	s.cf = cf
	s.resCh = resCh
	s.errCh = errCh
	s.isConnected = true
	go func() {
		for {
			select {
			case res, ok := <-s.resCh:
				if !ok {
					return
				}
				s.send(res)
			case err, ok := <-s.errCh:
				if err != nil && ok {
					s.sendError(err)
				}
			case <-s.ctx.Done():
				return
			}
		}
	}()
}

func (s *sessionStream) send(res *pb.StreamResponse) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	for _, cs := range s.streams {
		go cs.send(res)
	}
}

func (s *sessionStream) sendError(err error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	for _, cs := range s.streams {
		go cs.disconnect(err)
	}
	defer s.disconnect()
}

func (s *sessionStream) disconnect() {
	if s.isConnected {
		s.cf()
		s.isConnected = false
	}
}

type iClientStream interface {
	getRequest() *pb.StreamRequest
	start()
	send(res *pb.StreamResponse)
	disconnect(err error)
}

type clientStream struct {
	clientToken string
	request     *pb.StreamRequest
	stream      pb.TachibanaService_StreamServer
	ctx         context.Context
	errCh       chan<- error
	isConnected bool
}

func (s *clientStream) getRequest() *pb.StreamRequest {
	return s.request
}

func (s *clientStream) start() {
	s.isConnected = true
	go func() {
		<-s.ctx.Done()
		s.disconnect(s.ctx.Err())
	}()
}

func (s *clientStream) send(res *pb.StreamResponse) {
	if !s.isConnected {
		return
	}

	if !s.request.Sendable(res) {
		return
	}
	if err := s.stream.Send(res); err != nil {
		s.disconnect(err)
	}
}

func (s *clientStream) disconnect(err error) {
	if !s.isConnected {
		return
	}

	s.isConnected = false
	s.errCh <- err
}
