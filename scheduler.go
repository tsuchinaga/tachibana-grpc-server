package tachibana_grpc_server

import "time"

func (s *server) StartScheduler() {
	go s.clearSessionScheduler()
	select {}
}

func (s *server) clearSessionScheduler() {
	for {
		time.Sleep(s.clock.nextDateTimeDuration(time.Date(0, 1, 1, 6, 0, 0, 0, time.Local), s.clock.now()))
		s.logger.scheduler("start clearSessionScheduler")
		s.sessionStore.clear()
		s.logger.scheduler("end clearSessionScheduler")
	}
}
