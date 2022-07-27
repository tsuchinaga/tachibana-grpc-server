package tachibanapb

func (x *StreamRequest) Sendable(res *StreamResponse) bool {
	// 再送を受けないリクエストに対して、初回配信でないイベントは送信しない
	if !x.ReceiveResend && !res.IsFirstTime {
		return false
	}

	var contain bool
	for _, xe := range x.EventTypes {
		if xe == res.EventType {
			contain = true
			break
		}
	}

	if !contain {
		return false
	}

	// 市場価格情報なら参照している銘柄かのチェック
	if res.EventType == EventType_EVENT_TYPE_MARKET_PRICE {
		return res.MarketPriceStreamResponse != nil && x.hasIssue(res.MarketPriceStreamResponse.IssueCode, res.MarketPriceStreamResponse.Exchange)
	}

	return true
}

func (x *StreamRequest) hasIssue(issueCode string, exchange Exchange) bool {
	for _, issue := range x.StreamIssues {
		if issue.IssueCode == issueCode && issue.Exchange == exchange {
			return true
		}
	}
	return false
}

func (x *StreamRequest) Union(res *StreamRequest) *StreamRequest {
	if x.EventTypes == nil {
		x.EventTypes = make([]EventType, 0)
	}
	if x.StreamIssues == nil {
		x.StreamIssues = make([]*StreamIssue, 0)
	}

	// event types の 結合
	for _, e := range res.EventTypes {
		var contain bool
		for _, ne := range x.EventTypes {
			if e == ne {
				contain = true
				break
			}
		}
		if !contain {
			x.EventTypes = append(x.EventTypes, e)
		}
	}

	// issue の 結合
	for _, issue := range res.StreamIssues {
		var contain bool
		for _, ni := range x.StreamIssues {
			if issue.IssueCode == ni.IssueCode && issue.Exchange == ni.Exchange {
				contain = true
				break
			}
		}
		if !contain {
			x.StreamIssues = append(x.StreamIssues, issue)
		}
	}

	return x
}
