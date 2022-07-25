package tachibana_grpc_server

import (
	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (t *tachibanaApi) toLoginRequest(req *pb.LoginRequest) *tachibana.LoginRequest {
	return &tachibana.LoginRequest{
		UserId:   req.UserId,
		Password: req.Password,
	}
}

func (t *tachibanaApi) fromLoginResponse(res *tachibana.LoginResponse) *pb.LoginResponse {
	return &pb.LoginResponse{
		CommonResponse:            t.fromCommonResponse(&res.CommonResponse),
		ResultCode:                res.ResultCode,
		ResultText:                res.ResultText,
		AccountType:               t.fromAccountType(res.AccountType),
		SecondPasswordOmit:        res.SecondPasswordOmit,
		LastLoginDatetime:         timestamppb.New(res.LastLoginDateTime),
		GeneralAccount:            res.GeneralAccount,
		SafekeepingAccount:        res.SafekeepingAccount,
		TransferAccount:           res.TransferAccount,
		ForeignAccount:            res.ForeignAccount,
		MrfAccount:                res.MRFAccount,
		StockSpecificAccount:      t.fromSpecificAccountType(res.StockSpecificAccount),
		MarginSpecificAccount:     t.fromSpecificAccountType(res.MarginSpecificAccount),
		InvestmentSpecificAccount: t.fromSpecificAccountType(res.InvestmentSpecificAccount),
		DividendAccount:           res.DividendAccount,
		SpecificAccount:           res.SpecificAccount,
		MarginAccount:             res.MarginAccount,
		FutureOptionAccount:       res.FutureOptionAccount,
		MmfAccount:                res.MMFAccount,
		ChinaForeignAccount:       res.ChinaForeignAccount,
		FxAccount:                 res.FXAccount,
		NisaAccount:               res.NISAAccount,
		UnreadDocument:            res.UnreadDocument,
		Token:                     "",
	}
}

func (t *tachibanaApi) fromCommonResponse(res *tachibana.CommonResponse) *pb.CommonResponse {
	return &pb.CommonResponse{
		No:           res.No,
		SendDate:     timestamppb.New(res.SendDate),
		ReceiveDate:  timestamppb.New(res.ReceiveDate),
		ErrorNo:      t.fromErrorNo(res.ErrorNo),
		ErrorMessage: res.ErrorMessage,
		MessageType:  t.fromMessageType(res.MessageType),
	}
}

func (t *tachibanaApi) toNewOrderRequest(req *pb.NewOrderRequest) *tachibana.NewOrderRequest {
	return &tachibana.NewOrderRequest{
		AccountType:         t.toAccountType(req.AccountType),
		DeliveryAccountType: t.toDeliveryAccountType(req.DeliveryAccountType),
		IssueCode:           req.IssueCode,
		Exchange:            t.toExchange(req.Exchange),
		Side:                t.toSide(req.Side),
		ExecutionTiming:     t.toExecutionTiming(req.ExecutionTiming),
		OrderPrice:          req.OrderPrice,
		OrderQuantity:       req.OrderQuantity,
		TradeType:           t.toTradeType(req.TradeType),
		ExpireDate:          req.ExpireDate.AsTime().In(time.Local),
		ExpireDateIsToday:   req.ExpireDateIsToday,
		StopOrderType:       t.toStopOrderType(req.StopOrderType),
		TriggerPrice:        req.TriggerPrice,
		StopOrderPrice:      req.StopOrderPrice,
		ExitPositionType:    t.toExitPositionType(req.ExitPositionType),
		SecondPassword:      req.SecondPassword,
		ExitPositions:       t.toExitPositions(req.ExitPositions),
	}
}

func (t *tachibanaApi) toExitPositions(exitPositions []*pb.ExitPosition) []tachibana.ExitPosition {
	res := make([]tachibana.ExitPosition, len(exitPositions))
	for i, p := range exitPositions {
		res[i] = tachibana.ExitPosition{
			PositionNumber: p.PositionNumber,
			SequenceNumber: int(p.SequenceNumber),
			OrderQuantity:  p.OrderQuantity,
		}
	}
	return res
}

func (t *tachibanaApi) fromNewOrderResponse(res *tachibana.NewOrderResponse) *pb.NewOrderResponse {
	return &pb.NewOrderResponse{
		CommonResponse: t.fromCommonResponse(&res.CommonResponse),
		ResultCode:     res.ResultCode,
		ResultText:     res.ResultText,
		WarningCode:    res.WarningCode,
		WarningText:    res.WarningText,
		OrderNumber:    res.OrderNumber,
		ExecutionDate:  timestamppb.New(res.ExecutionDate),
		DeliveryAmount: res.DeliveryAmount,
		Commission:     res.Commission,
		CommissionTax:  res.CommissionTax,
		Interest:       res.Interest,
		OrderDatetime:  timestamppb.New(res.OrderDateTime),
	}
}

func (t *tachibanaApi) toCancelOrderRequest(req *pb.CancelOrderRequest) *tachibana.CancelOrderRequest {
	return &tachibana.CancelOrderRequest{
		OrderNumber:    req.OrderNumber,
		ExecutionDate:  req.ExecutionDate.AsTime().In(time.Local),
		SecondPassword: req.SecondPassword,
	}
}

func (t *tachibanaApi) fromCancelOrderResponse(res *tachibana.CancelOrderResponse) *pb.CancelOrderResponse {
	return &pb.CancelOrderResponse{
		CommonResponse: t.fromCommonResponse(&res.CommonResponse),
		ResultCode:     res.ResultCode,
		ResultText:     res.ResultText,
		OrderNumber:    res.OrderNumber,
		ExecutionDate:  timestamppb.New(res.ExecutionDate),
		DeliveryAmount: res.DeliveryAmount,
		OrderDatetime:  timestamppb.New(res.OrderDateTime),
	}
}

func (t *tachibanaApi) toOrderListRequest(req *pb.OrderListRequest) *tachibana.OrderListRequest {
	return &tachibana.OrderListRequest{
		IssueCode:          req.IssueCode,
		ExecutionDate:      req.ExecutionDate.AsTime().In(time.Local),
		OrderInquiryStatus: t.toOrderInquiryStatus(req.OrderInquiryStatus),
	}
}

func (t *tachibanaApi) fromOrderListResponse(res *tachibana.OrderListResponse) *pb.OrderListResponse {
	return &pb.OrderListResponse{
		CommonResponse:     t.fromCommonResponse(&res.CommonResponse),
		IssueCode:          res.IssueCode,
		ExecutionDate:      timestamppb.New(res.ExecutionDate),
		OrderInquiryStatus: t.fromOrderInquiryStatus(res.OrderInquiryStatus),
		ResultCode:         res.ResultCode,
		ResultText:         res.ResultText,
		WarningCode:        res.WarningCode,
		WarningText:        res.WarningText,
		Orders:             t.fromOrders(res.Orders),
	}
}

func (t *tachibanaApi) fromOrders(orders []tachibana.Order) []*pb.Order {
	res := make([]*pb.Order, len(orders))
	for i, o := range orders {
		res[i] = &pb.Order{
			WarningCode:            o.WarningCode,
			WarningText:            o.WarningText,
			OrderNumber:            o.OrderNumber,
			IssueCode:              o.IssueCode,
			Exchange:               t.fromExchange(o.Exchange),
			AccountType:            t.fromAccountType(o.AccountType),
			TradeType:              t.fromTradeType(o.TradeType),
			ExitTermType:           t.fromExitTermType(o.ExitTermType),
			Side:                   t.fromSide(o.Side),
			OrderQuantity:          o.OrderQuantity,
			CurrentQuantity:        o.CurrentQuantity,
			Price:                  o.Price,
			ExecutionTiming:        t.fromExecutionTiming(o.ExecutionTiming),
			ExecutionType:          t.fromExecutionType(o.ExecutionType),
			StopOrderType:          t.fromStopOrderType(o.StopOrderType),
			StopTriggerPrice:       o.StopTriggerPrice,
			StopOrderExecutionType: t.fromExecutionType(o.StopOrderExecutionType),
			StopOrderPrice:         o.StopOrderPrice,
			TriggerType:            t.fromTriggerType(o.TriggerType),
			ExitPositionType:       t.fromExitPositionType(o.ExitPositionType),
			ContractQuantity:       o.ContractQuantity,
			ContractPrice:          o.ContractPrice,
			PartContractType:       t.fromPartContractType(o.PartContractType),
			ExecutionDate:          timestamppb.New(o.ExecutionDate),
			OrderStatus:            t.fromOrderStatus(o.OrderStatus),
			OrderStatusText:        o.OrderStatusText,
			ContractStatus:         t.fromContractStatus(o.ContractStatus),
			OrderDatetime:          timestamppb.New(o.OrderDateTime),
			ExpireDate:             timestamppb.New(o.ExpireDate),
			CarryOverType:          t.fromCarryOverType(o.CarryOverType),
			CorrectCancelType:      t.fromCorrectCancelType(o.CorrectCancelType),
			EstimationAmount:       o.EstimationAmount,
		}
	}
	return res
}

func (t *tachibanaApi) toOrderDetailRequest(req *pb.OrderDetailRequest) *tachibana.OrderDetailRequest {
	return &tachibana.OrderDetailRequest{
		OrderNumber:   req.OrderNumber,
		ExecutionDate: req.ExecutionDate.AsTime().In(time.Local),
	}
}

func (t *tachibanaApi) fromOrderDetailResponse(res *tachibana.OrderDetailResponse) *pb.OrderDetailResponse {
	return &pb.OrderDetailResponse{
		CommonResponse:         t.fromCommonResponse(&res.CommonResponse),
		OrderNumber:            res.OrderNumber,
		ExecutionDate:          timestamppb.New(res.ExecutionDate),
		ResultCode:             res.ResultCode,
		ResultText:             res.ResultText,
		WarningCode:            res.WarningCode,
		WarningText:            res.WarningText,
		IssueCode:              res.IssueCode,
		Exchange:               t.fromExchange(res.Exchange),
		Side:                   t.fromSide(res.Side),
		TradeType:              t.fromTradeType(res.TradeType),
		ExitTermType:           t.fromExitTermType(res.ExitTermType),
		ExecutionTiming:        t.fromExecutionTiming(res.ExecutionTiming),
		ExecutionType:          t.fromExecutionType(res.ExecutionType),
		Price:                  res.Price,
		OrderQuantity:          res.OrderQuantity,
		CurrentQuantity:        res.CurrentQuantity,
		OrderStatus:            t.fromOrderStatus(res.OrderStatus),
		OrderStatusText:        res.OrderStatusText,
		OrderDatetime:          timestamppb.New(res.OrderDateTime),
		ExpireDate:             timestamppb.New(res.ExpireDate),
		Channel:                t.fromChannel(res.Channel),
		StockAccountType:       t.fromAccountType(res.StockAccountType),
		MarginAccountType:      t.fromAccountType(res.MarginAccountType),
		StopOrderType:          t.fromStopOrderType(res.StopOrderType),
		StopTriggerPrice:       res.StopTriggerPrice,
		StopOrderExecutionType: t.fromExecutionType(res.StopOrderExecutionType),
		StopOrderPrice:         res.StopOrderPrice,
		TriggerType:            t.fromTriggerType(res.TriggerType),
		TriggerDatetime:        timestamppb.New(res.TriggerDateTime),
		DeliveryDate:           timestamppb.New(res.DeliveryDate),
		ContractPrice:          res.ContractPrice,
		ContractQuantity:       res.ContractQuantity,
		TradingAmount:          res.TradingAmount,
		PartContractType:       t.fromPartContractType(res.PartContractType),
		EstimationAmount:       res.EstimationAmount,
		Commission:             res.Commission,
		CommissionTax:          res.CommissionTax,
		ExitPositionType:       t.fromExitPositionType(res.ExitPositionType),
		ExchangeErrorCode:      res.ExchangeErrorCode,
		ExchangeOrderDatetime:  timestamppb.New(res.ExchangeOrderDateTime),
		Contracts:              t.fromContracts(res.Contracts),
		HoldPositions:          t.fromHoldPositions(res.HoldPositions),
	}
}

func (t *tachibanaApi) fromContracts(contracts []tachibana.Contract) []*pb.Contract {
	res := make([]*pb.Contract, len(contracts))
	for i, c := range contracts {
		res[i] = &pb.Contract{
			WarningCode: c.WarningCode,
			WarningText: c.WarningText,
			Quantity:    c.Quantity,
			Price:       c.Price,
			Datetime:    timestamppb.New(c.DateTime),
		}
	}
	return res
}

func (t *tachibanaApi) fromHoldPositions(holdPositions []tachibana.HoldPosition) []*pb.HoldPosition {
	res := make([]*pb.HoldPosition, len(holdPositions))
	for i, hp := range holdPositions {
		res[i] = &pb.HoldPosition{
			WarningCode:   hp.WarningCode,
			WarningText:   hp.WarningText,
			SortOrder:     int32(hp.SortOrder),
			ContractDate:  timestamppb.New(hp.ContractDate),
			EntryPrice:    hp.EntryPrice,
			HoldQuantity:  hp.HoldQuantity,
			ExitQuantity:  hp.ExitQuantity,
			ExitPrice:     hp.ExitPrice,
			Commission:    hp.Commission,
			Interest:      hp.Interest,
			Premiums:      hp.Premiums,
			RewritingFee:  hp.RewritingFee,
			ManagementFee: hp.ManagementFee,
			LendingFee:    hp.LendingFee,
			OtherFee:      hp.OtherFee,
			Profit:        hp.Profit,
		}
	}
	return res
}

func (t *tachibanaApi) toStockMasterRequest(req *pb.StockMasterRequest) *tachibana.StockMasterRequest {
	return &tachibana.StockMasterRequest{
		Columns: t.toStockMasterColumns(req.Columns),
	}
}

func (t *tachibanaApi) toStockMasterColumns(columns []pb.StockMasterColumn) []tachibana.StockMasterColumn {
	res := make([]tachibana.StockMasterColumn, len(columns))
	for i, c := range columns {
		res[i] = t.toStockMasterColumn(c)
	}
	return res
}

func (t *tachibanaApi) fromStockMasterResponse(res *tachibana.StockMasterResponse) *pb.StockMasterResponse {
	return &pb.StockMasterResponse{
		CommonResponse: t.fromCommonResponse(&res.CommonResponse),
		StockMasters:   t.fromStockMasters(res.StockMasters),
	}
}

func (t *tachibanaApi) fromStockMasters(stockMasters []tachibana.StockMaster) []*pb.StockMaster {
	res := make([]*pb.StockMaster, len(stockMasters))
	for i, s := range stockMasters {
		res[i] = &pb.StockMaster{
			IssueCode:            s.IssueCode,
			Name:                 s.Name,
			ShortName:            s.ShortName,
			Kana:                 s.Kana,
			Alphabet:             s.Alphabet,
			SpecificTarget:       s.SpecificTarget,
			TaxFree:              t.fromTaxFree(s.TaxFree),
			SharedStocks:         s.SharedStocks,
			ExRightType:          t.fromExRightType(s.ExRight),
			LastRightDay:         timestamppb.New(s.LastRightDay),
			ListingType:          t.fromListingType(s.ListingType),
			ReleaseTradingDate:   timestamppb.New(s.ReleaseTradingDate),
			TradingDate:          timestamppb.New(s.TradingDate),
			TradingUnit:          s.TradingUnit,
			NextTradingUnit:      s.NextTradingUnit,
			StopTradingType:      t.fromStopTradingType(s.StopTradingType),
			StartPublicationDate: timestamppb.New(s.StartPublicationDate),
			LastPublicationDate:  timestamppb.New(s.LastPublicationDate),
			SettlementType:       t.fromSettlementType(s.SettlementType),
			SettlementDate:       timestamppb.New(s.SettlementDate),
			ListingDate:          timestamppb.New(s.ListingDate),
			ExpireDate_2Type:     s.ExpireDate2Type,
			LargeUnit:            s.LargeUnit,
			LargeAmount:          s.LargeAmount,
			OutputTicketType:     s.OutputTicketType,
			DepositAmount:        s.DepositAmount,
			DepositValuation:     s.DepositValuation,
			OrganizationType:     s.OrganizationType,
			ProvisionalType:      s.ProvisionalType,
			PrimaryExchange:      t.fromExchange(s.PrimaryExchange),
			IndefinitePeriodType: s.IndefinitePeriodType,
			IndustryCode:         s.IndustryCode,
			IndustryName:         s.IndustryName,
			SorTargetType:        s.SORTargetType,
			CreateDatetime:       timestamppb.New(s.CreateDateTime),
			UpdateDatetime:       timestamppb.New(s.UpdateDateTime),
			UpdateNumber:         s.UpdateNumber,
		}
	}
	return res
}

func (t *tachibanaApi) toStockExchangeMasterRequest(req *pb.StockExchangeMasterRequest) *tachibana.StockExchangeMasterRequest {
	return &tachibana.StockExchangeMasterRequest{
		Columns: t.toStockExchangeMasterColumns(req.Columns),
	}
}

func (t *tachibanaApi) toStockExchangeMasterColumns(columns []pb.StockExchangeMasterColumn) []tachibana.StockExchangeMasterColumn {
	res := make([]tachibana.StockExchangeMasterColumn, len(columns))
	for i, c := range columns {
		res[i] = t.toStockExchangeMasterColumn(c)
	}
	return res
}

func (t *tachibanaApi) fromStockExchangeMasterResponse(res *tachibana.StockExchangeMasterResponse) *pb.StockExchangeMasterResponse {
	return &pb.StockExchangeMasterResponse{
		CommonResponse:       t.fromCommonResponse(&res.CommonResponse),
		StockExchangeMasters: t.fromStockExchangeMasters(res.StockExchangeMasters),
	}
}

func (t *tachibanaApi) fromStockExchangeMasters(stockExchangeMasters []tachibana.StockExchangeMaster) []*pb.StockExchangeMaster {
	res := make([]*pb.StockExchangeMaster, len(stockExchangeMasters))
	for i, r := range stockExchangeMasters {
		res[i] = &pb.StockExchangeMaster{
			IssueCode:                   r.IssueCode,
			Exchange:                    t.fromExchange(r.Exchange),
			StockSystemType:             r.StockSystemType,
			UnderLimitPrice:             r.UnderLimitPrice,
			UpperLimitPrice:             r.UpperLimitPrice,
			SymbolCategory:              r.SymbolCategory,
			LimitPriceExchange:          t.fromExchange(r.LimitPriceExchange),
			MarginType:                  t.fromMarginType(r.MarginType),
			ListingDate:                 timestamppb.New(r.ListingDate),
			LimitPriceDate:              timestamppb.New(r.LimitPriceDate),
			LimitPriceCategory:          r.LimitPriceCategory,
			LimitPriceValue:             r.LimitPriceValue,
			ConfirmLimitPrice:           r.ConfirmLimitPrice,
			Section:                     r.Section,
			PrevClosePrice:              r.PrevClosePrice,
			CalculateLimitPriceExchange: t.fromExchange(r.CalculateLimitPriceExchange),
			Regulation1:                 r.Regulation1,
			Regulation2:                 r.Regulation2,
			SectionType:                 r.SectionType,
			DelistingDate:               timestamppb.New(r.DelistingDate),
			TradingUnit:                 r.TradingUnit,
			NextTradingUnit:             r.NextTradingUnit,
			TickGroupType:               t.fromTickGroupType(r.TickGroupType),
			NextTickGroupType:           t.fromTickGroupType(r.NextTickGroupType),
			InformationSource:           r.InformationSource,
			InformationCode:             r.InformationCode,
			OfferPrice:                  r.OfferPrice,
			CreateDatetime:              timestamppb.New(r.CreateDateTime),
			UpdateDatetime:              timestamppb.New(r.UpdateDateTime),
			UpdateNumber:                r.UpdateNumber,
		}
	}
	return res
}

func (t *tachibanaApi) toMarketPriceRequest(req *pb.MarketPriceRequest) *tachibana.MarketPriceRequest {
	return &tachibana.MarketPriceRequest{
		IssueCodes: req.IssueCodes,
		Columns:    t.fromMarketPriceColumns(req.Columns),
	}
}

func (t *tachibanaApi) fromMarketPriceColumns(columns []pb.MarketPriceColumn) []tachibana.MarketPriceColumn {
	res := make([]tachibana.MarketPriceColumn, len(columns))
	for i, c := range columns {
		res[i] = t.fromMarketPriceColumn(c)
	}
	return res
}

func (t *tachibanaApi) fromMarketPriceResponse(res *tachibana.MarketPriceResponse) *pb.MarketPriceResponse {
	return &pb.MarketPriceResponse{
		CommonResponse: t.fromCommonResponse(&res.CommonResponse),
		MarketPrices:   t.fromMarketPrices(res.MarketPrices),
	}
}

func (t *tachibanaApi) fromMarketPrices(marketPrices []tachibana.MarketPrice) []*pb.MarketPrice {
	res := make([]*pb.MarketPrice, len(marketPrices))
	for i, p := range marketPrices {
		res[i] = &pb.MarketPrice{
			IssueCode:         p.IssueCode,
			Section:           p.Section,
			CurrentPrice:      p.CurrentPrice,
			CurrentPriceTime:  timestamppb.New(p.CurrentPriceTime),
			ChangePriceType:   t.fromChangePriceType(p.ChangePriceType),
			PrevDayRatio:      p.PrevDayRatio,
			PrevDayPercent:    p.PrevDayPercent,
			OpenPrice:         p.OpenPrice,
			OpenPriceTime:     timestamppb.New(p.OpenPriceTime),
			HighPrice:         p.HighPrice,
			HighPriceTime:     timestamppb.New(p.HighPriceTime),
			LowPrice:          p.LowPrice,
			LowPriceTime:      timestamppb.New(p.LowPriceTime),
			Volume:            p.Volume,
			AskSign:           t.fromIndicationPriceType(p.AskSign),
			AskPrice:          p.AskPrice,
			AskQuantity:       p.AskQuantity,
			BidSign:           t.fromIndicationPriceType(p.BidSign),
			BidPrice:          p.BidPrice,
			BidQuantity:       p.BidQuantity,
			ExRightType:       p.ExRightType,
			DiscontinuityType: p.DiscontinuityType,
			StopHigh:          t.fromCurrentPriceType(p.StopHigh),
			StopLow:           t.fromCurrentPriceType(p.StopLow),
			TradingAmount:     p.TradingAmount,
			AskQuantityMarket: p.AskQuantityMarket,
			BidQuantityMarket: p.BidQuantityMarket,
			AskQuantityOver:   p.AskQuantityOver,
			AskQuantity10:     p.AskQuantity10,
			AskPrice10:        p.AskPrice10,
			AskQuantity9:      p.AskQuantity9,
			AskPrice9:         p.AskPrice9,
			AskQuantity8:      p.AskQuantity8,
			AskPrice8:         p.AskPrice8,
			AskQuantity7:      p.AskQuantity7,
			AskPrice7:         p.AskPrice7,
			AskQuantity6:      p.AskQuantity6,
			AskPrice6:         p.AskPrice6,
			AskQuantity5:      p.AskQuantity5,
			AskPrice5:         p.AskPrice5,
			AskQuantity4:      p.AskQuantity4,
			AskPrice4:         p.AskPrice4,
			AskQuantity3:      p.AskQuantity3,
			AskPrice3:         p.AskPrice3,
			AskQuantity2:      p.AskQuantity2,
			AskPrice2:         p.AskPrice2,
			AskQuantity1:      p.AskQuantity1,
			AskPrice1:         p.AskPrice1,
			BidQuantity1:      p.BidQuantity1,
			BidPrice1:         p.BidPrice1,
			BidQuantity2:      p.BidQuantity2,
			BidPrice2:         p.BidPrice2,
			BidQuantity3:      p.BidQuantity3,
			BidPrice3:         p.BidPrice3,
			BidQuantity4:      p.BidQuantity4,
			BidPrice4:         p.BidPrice4,
			BidQuantity5:      p.BidQuantity5,
			BidPrice5:         p.BidPrice5,
			BidQuantity6:      p.BidQuantity6,
			BidPrice6:         p.BidPrice6,
			BidQuantity7:      p.BidQuantity7,
			BidPrice7:         p.BidPrice7,
			BidQuantity8:      p.BidQuantity8,
			BidPrice8:         p.BidPrice8,
			BidQuantity9:      p.BidQuantity9,
			BidPrice9:         p.BidPrice9,
			BidQuantity10:     p.BidQuantity10,
			BidPrice10:        p.BidPrice10,
			BidQuantityUnder:  p.BidQuantityUnder,
			Vwap:              p.VWAP,
			Prp:               p.PRP,
		}
	}
	return res
}

func (t *tachibanaApi) toBusinessDayRequest(_ *pb.BusinessDayRequest) *tachibana.BusinessDayRequest {
	return &tachibana.BusinessDayRequest{}
}

func (t *tachibanaApi) fromBusinessDayResponse(res []*tachibana.BusinessDayResponse) *pb.BusinessDayResponse {
	var commonResponse *pb.CommonResponse
	var businessDays []*pb.BusinessDay

	if len(res) > 0 {
		commonResponse = t.fromCommonResponse(&res[len(res)-1].CommonResponse)
		businessDays = make([]*pb.BusinessDay, len(res))

		for i, r := range res {
			businessDays[i] = t.fromBusinessDay(r)
		}
	}

	return &pb.BusinessDayResponse{
		CommonResponse: commonResponse,
		BusinessDays:   businessDays,
	}
}

func (t *tachibanaApi) fromBusinessDay(businessDay *tachibana.BusinessDayResponse) *pb.BusinessDay {
	return &pb.BusinessDay{
		DayKey:                 t.fromDayKey(businessDay.DayKey),
		PrevDay1:               timestamppb.New(businessDay.PrevDay1),
		PrevDay2:               timestamppb.New(businessDay.PrevDay2),
		PrevDay3:               timestamppb.New(businessDay.PrevDay3),
		Today:                  timestamppb.New(businessDay.Today),
		NextDay1:               timestamppb.New(businessDay.NextDay1),
		NextDay2:               timestamppb.New(businessDay.NextDay2),
		NextDay3:               timestamppb.New(businessDay.NextDay3),
		NextDay4:               timestamppb.New(businessDay.NextDay4),
		NextDay5:               timestamppb.New(businessDay.NextDay5),
		NextDay6:               timestamppb.New(businessDay.NextDay6),
		NextDay7:               timestamppb.New(businessDay.NextDay7),
		NextDay8:               timestamppb.New(businessDay.NextDay8),
		NextDay9:               timestamppb.New(businessDay.NextDay9),
		NextDay10:              timestamppb.New(businessDay.NextDay10),
		DeliveryDay:            timestamppb.New(businessDay.DeliveryDay),
		ProvisionalDeliveryDay: timestamppb.New(businessDay.ProvisionalDeliveryDay),
		BondDeliveryDay:        timestamppb.New(businessDay.BondDeliveryDay),
	}
}

func (t *tachibanaApi) toTickGroupRequest(_ *pb.TickGroupRequest) *tachibana.TickGroupRequest {
	return &tachibana.TickGroupRequest{}
}

func (t *tachibanaApi) fromTickGroupResponse(res []*tachibana.TickGroupResponse) *pb.TickGroupResponse {
	var commonResponse *pb.CommonResponse
	var tickGroups []*pb.TickGroup

	if len(res) > 0 {
		commonResponse = t.fromCommonResponse(&res[len(res)-1].CommonResponse)
		tickGroups = make([]*pb.TickGroup, len(res))

		for i, tg := range res {
			tickGroups[i] = t.fromTickGroup(tg)
		}
	}

	return &pb.TickGroupResponse{
		CommonResponse: commonResponse,
		TickGroups:     tickGroups,
	}
}

func (t *tachibanaApi) fromTickGroup(res *tachibana.TickGroupResponse) *pb.TickGroup {
	return &pb.TickGroup{
		TickGroupType: t.fromTickGroupType(res.TickGroupType),
		StartDate:     timestamppb.New(res.StartDate),
		TickGroupList: t.fromTickGroups(res.TickGroups),
		CreateDate:    timestamppb.New(res.CreateDate),
		UpdateDate:    timestamppb.New(res.UpdateDate),
	}
}

func (t *tachibanaApi) fromTickGroups(tickGroups [20]tachibana.TickGroup) []*pb.TickGroupPrice {
	res := make([]*pb.TickGroupPrice, len(tickGroups))
	for i, tg := range tickGroups {
		res[i] = t.fromTickGroupPrice(tg)
	}
	return res
}

func (t *tachibanaApi) fromTickGroupPrice(tickGroup tachibana.TickGroup) *pb.TickGroupPrice {
	return &pb.TickGroupPrice{
		Number:    int32(tickGroup.Number),
		BasePrice: tickGroup.BasePrice,
		UnitPrice: tickGroup.UnitPrice,
		Digits:    int32(tickGroup.Digits),
	}
}

func (t *tachibanaApi) toStreamRequest(req *pb.StreamRequest) *tachibana.StreamRequest {
	r := &tachibana.StreamRequest{
		ColumnNumber:      []int{},
		IssueCodes:        []string{},
		MarketCodes:       []tachibana.Exchange{},
		StartStreamNumber: 0,
		StreamEventTypes:  []tachibana.EventType{},
	}

	// event type
	for _, e := range req.EventTypes {
		r.StreamEventTypes = append(r.StreamEventTypes, t.toEventType(e))
	}

	// issues
	for i, issue := range req.StreamIssues {
		r.ColumnNumber = append(r.ColumnNumber, i+1)
		r.IssueCodes = append(r.IssueCodes, issue.IssueCode)
		r.MarketCodes = append(r.MarketCodes, t.toExchange(issue.Exchange))
	}

	return r
}

func (t *tachibanaApi) fromStreamResponse(res tachibana.StreamResponse) *pb.StreamResponse {
	switch v := res.(type) {
	case *tachibana.CommonStreamResponse:
		return t.fromCommonStreamResponse(v)
	case *tachibana.ContractStreamResponse:
		return t.fromContractStreamResponse(v)
	case *tachibana.NewsStreamResponse:
		return t.fromNewsStreamResponse(v)
	case *tachibana.SystemStatusStreamResponse:
		return t.fromSystemStatusStreamResponse(v)
	case *tachibana.OperationStatusStreamResponse:
		return t.fromOperationStatusStreamResponse(v)
	}
	return &pb.StreamResponse{
		EventType:    t.fromEventType(res.GetEventType()),
		ErrorNo:      t.fromErrorNo(res.GetErrorNo()),
		ErrorMessage: res.GetErrorText(),
		Body:         nil,
		IsFirstTime:  true,
	}
}

func (t *tachibanaApi) fromCommonStreamResponse(res *tachibana.CommonStreamResponse) *pb.StreamResponse {
	return &pb.StreamResponse{
		EventType:                     t.fromEventType(res.EventType),
		StreamNumber:                  res.StreamNumber,
		StreamDateTime:                timestamppb.New(res.StreamDateTime),
		ErrorNo:                       t.fromErrorNo(res.ErrorNo),
		ErrorMessage:                  res.ErrorText,
		Body:                          res.Body,
		IsFirstTime:                   true,
		ContractStreamResponse:        nil,
		NewsStreamResponse:            nil,
		SystemStatusStreamResponse:    nil,
		OperationStatusStreamResponse: nil,
	}
}

func (t *tachibanaApi) fromContractStreamResponse(res *tachibana.ContractStreamResponse) *pb.StreamResponse {
	return &pb.StreamResponse{
		EventType:      t.fromEventType(res.EventType),
		StreamNumber:   res.StreamNumber,
		StreamDateTime: timestamppb.New(res.StreamDateTime),
		ErrorNo:        t.fromErrorNo(res.ErrorNo),
		ErrorMessage:   res.ErrorText,
		Body:           res.Body,
		IsFirstTime:    res.FirstTime,
		ContractStreamResponse: &pb.ContractStreamResponse{
			Provider:                 res.Provider,
			EventNo:                  res.EventNo,
			StreamOrderType:          t.fromStreamOrderType(res.StreamOrderType),
			OrderNumber:              res.OrderNumber,
			ExecutionDate:            timestamppb.New(res.ExecutionDate),
			ParentOrderNumber:        res.ParentOrderNumber,
			ParentOrder:              res.ParentOrder,
			ProductType:              t.fromProductType(res.ProductType),
			IssueCode:                res.IssueCode,
			Exchange:                 t.fromExchange(res.Exchange),
			Side:                     t.fromSide(res.Side),
			TradeType:                t.fromTradeType(res.TradeType),
			ExecutionTiming:          t.fromExecutionTiming(res.ExecutionTiming),
			ExecutionType:            t.fromExecutionType(res.ExecutionType),
			Price:                    res.Price,
			Quantity:                 res.Quantity,
			CancelQuantity:           res.CancelQuantity,
			ExpireQuantity:           res.ExpireQuantity,
			ContractQuantity:         res.ContractQuantity,
			StreamOrderStatus:        t.fromStreamOrderStatus(res.StreamOrderStatus),
			CarryOverType:            t.fromCarryOverType(res.CarryOverType),
			CancelOrderStatus:        t.fromCancelOrderStatus(res.CancelOrderStatus),
			ContractStatus:           t.fromContractStatus(res.ContractStatus),
			ExpireDate:               timestamppb.New(res.ExpireDate),
			SecurityExpireReason:     res.SecurityExpireReason,
			SecurityContractPrice:    res.SecurityContractPrice,
			SecurityContractQuantity: res.SecurityContractQuantity,
			SecurityError:            res.SecurityError,
			NotifyDatetime:           timestamppb.New(res.NotifyDateTime),
			IssueName:                res.IssueName,
			CorrectExecutionTiming:   t.fromExecutionTiming(res.CorrectExecutionTiming),
			CorrectContractQuantity:  res.CorrectContractQuantity,
			CorrectExecutionType:     t.fromExecutionType(res.CorrectExecutionType),
			CorrectPrice:             res.CorrectPrice,
			CorrectQuantity:          res.CorrectQuantity,
			CorrectExpireDate:        timestamppb.New(res.CorrectExpireDate),
			CorrectStopOrderType:     t.fromStopOrderType(res.CorrectStopOrderType),
			CorrectTriggerPrice:      res.CorrectTriggerPrice,
			CorrectStopOrderPrice:    res.CorrectStopOrderPrice,
		},
	}
}

func (t *tachibanaApi) fromNewsStreamResponse(res *tachibana.NewsStreamResponse) *pb.StreamResponse {
	return &pb.StreamResponse{
		EventType:      t.fromEventType(res.EventType),
		StreamNumber:   res.StreamNumber,
		StreamDateTime: timestamppb.New(res.StreamDateTime),
		ErrorNo:        t.fromErrorNo(res.ErrorNo),
		ErrorMessage:   res.ErrorText,
		Body:           res.Body,
		IsFirstTime:    res.FirstTime,
		NewsStreamResponse: &pb.NewsStreamResponse{
			Provider:      res.Provider,
			EventNo:       res.EventNo,
			NewsId:        res.NewsId,
			NewsDatetime:  timestamppb.New(res.NewsDateTime),
			NumOfCategory: int64(res.NumOfCategory),
			Categories:    res.Categories,
			NumOfGenre:    int64(res.NumOfGenre),
			Genres:        res.Genres,
			NumOfIssue:    int64(res.NumOfIssue),
			Issues:        res.Issues,
			Title:         res.Title,
			Content:       res.Content,
		},
	}
}

func (t *tachibanaApi) fromSystemStatusStreamResponse(res *tachibana.SystemStatusStreamResponse) *pb.StreamResponse {
	return &pb.StreamResponse{
		EventType:      t.fromEventType(res.EventType),
		StreamNumber:   res.StreamNumber,
		StreamDateTime: timestamppb.New(res.StreamDateTime),
		ErrorNo:        t.fromErrorNo(res.ErrorNo),
		ErrorMessage:   res.ErrorText,
		Body:           res.Body,
		IsFirstTime:    res.FirstTime,
		SystemStatusStreamResponse: &pb.SystemStatusStreamResponse{
			Provider:       res.Provider,
			EventNo:        res.EventNo,
			UpdateDatetime: timestamppb.New(res.UpdateDateTime),
			ApprovalLogin:  t.fromApprovalLogin(res.ApprovalLogin),
			SystemStatus:   t.fromSystemStatus(res.SystemStatus),
		},
	}
}

func (t *tachibanaApi) fromOperationStatusStreamResponse(res *tachibana.OperationStatusStreamResponse) *pb.StreamResponse {
	return &pb.StreamResponse{
		EventType:      t.fromEventType(res.EventType),
		StreamNumber:   res.StreamNumber,
		StreamDateTime: timestamppb.New(res.StreamDateTime),
		ErrorNo:        t.fromErrorNo(res.ErrorNo),
		ErrorMessage:   res.ErrorText,
		Body:           res.Body,
		IsFirstTime:    res.FirstTime,
		OperationStatusStreamResponse: &pb.OperationStatusStreamResponse{
			Provider:          res.Provider,
			EventNo:           res.EventNo,
			UpdateDatetime:    timestamppb.New(res.UpdateDateTime),
			Exchange:          t.fromExchange(res.Exchange),
			AssetCode:         res.AssetCode,
			ProductType:       res.ProductType,
			OperationCategory: res.OperationCategory,
			OperationUnit:     res.OperationUnit,
			BusinessDayType:   res.BusinessDayType,
			OperationStatus:   res.OperationStatus,
		},
	}
}

func (t *tachibanaApi) fromErrorNo(errorNo tachibana.ErrorNo) pb.ErrorNo {
	switch errorNo {
	case tachibana.ErrorNoProblem:
		return pb.ErrorNo_ERROR_NO_NO_PROBLEM
	case tachibana.ErrorNoData:
		return pb.ErrorNo_ERROR_NO_NO_DATA
	case tachibana.ErrorSessionInactive:
		return pb.ErrorNo_ERROR_NO_SESSION_INACTIVE
	case tachibana.ErrorProgressedNumber:
		return pb.ErrorNo_ERROR_NO_PROGRESSED_NUMBER
	case tachibana.ErrorExceedLimitTime:
		return pb.ErrorNo_ERROR_NO_EXCEED_LIMIT_TIME
	case tachibana.ErrorServiceOffline:
		return pb.ErrorNo_ERROR_NO_SERVICE_OFFLINE
	case tachibana.ErrorBadRequest:
		return pb.ErrorNo_ERROR_NO_BAD_REQUEST
	case tachibana.ErrorDatabaseAccess:
		return pb.ErrorNo_ERROR_NO_DATABASE_ACCESS
	case tachibana.ErrorServerAccess:
		return pb.ErrorNo_ERROR_NO_SERVER_ACCESS
	case tachibana.ErrorSystemOffline:
		return pb.ErrorNo_ERROR_NO_SYSTEM_OFFLINE
	case tachibana.ErrorOffHours:
		return pb.ErrorNo_ERROR_NO_OFF_HOURS
	}
	return pb.ErrorNo_ERROR_NO_UNSPECIFIED
}

func (t *tachibanaApi) fromMessageType(messageType tachibana.MessageType) pb.MessageType {
	switch messageType {
	case tachibana.MessageTypeLoginRequest:
		return pb.MessageType_MESSAGE_TYPE_LOGIN_REQUEST
	case tachibana.MessageTypeLoginResponse:
		return pb.MessageType_MESSAGE_TYPE_LOGIN_RESPONSE
	case tachibana.MessageTypeLogoutRequest:
		return pb.MessageType_MESSAGE_TYPE_LOGOUT_REQUEST
	case tachibana.MessageTypeLogoutResponse:
		return pb.MessageType_MESSAGE_TYPE_LOGOUT_RESPONSE
	case tachibana.MessageTypeNewOrder:
		return pb.MessageType_MESSAGE_TYPE_NEW_ORDER
	case tachibana.MessageTypeCorrectOrder:
		return pb.MessageType_MESSAGE_TYPE_CORRECT_ORDER
	case tachibana.MessageTypeCancelOrder:
		return pb.MessageType_MESSAGE_TYPE_CANCEL_ORDER
	case tachibana.MessageTypeStockPositionList:
		return pb.MessageType_MESSAGE_TYPE_STOCK_POSITION_LIST
	case tachibana.MessageTypeMarginPositionList:
		return pb.MessageType_MESSAGE_TYPE_MARGIN_POSITION_LIST
	case tachibana.MessageTypeStockWallet:
		return pb.MessageType_MESSAGE_TYPE_STOCK_WALLET
	case tachibana.MessageTypeMarginWallet:
		return pb.MessageType_MESSAGE_TYPE_MARGIN_WALLET
	case tachibana.MessageTypeStockSellable:
		return pb.MessageType_MESSAGE_TYPE_STOCK_SELLABLE
	case tachibana.MessageTypeOrderList:
		return pb.MessageType_MESSAGE_TYPE_ORDER_LIST
	case tachibana.MessageTypeOrderDetail:
		return pb.MessageType_MESSAGE_TYPE_ORDER_DETAIL
	case tachibana.MessageTypeSummary:
		return pb.MessageType_MESSAGE_TYPE_SUMMARY
	case tachibana.MessageTypeSummaryRecord:
		return pb.MessageType_MESSAGE_TYPE_SUMMARY_RECORD
	case tachibana.MessageTypeStockEntryDetail:
		return pb.MessageType_MESSAGE_TYPE_STOCK_ENTRY_DETAIL
	case tachibana.MessageTypeMarginEntryDetail:
		return pb.MessageType_MESSAGE_TYPE_MARGIN_ENTRY_DETAIL
	case tachibana.MessageTypeDepositRate:
		return pb.MessageType_MESSAGE_TYPE_DEPOSIT_RATE
	case tachibana.MessageTypeEventDownload:
		return pb.MessageType_MESSAGE_TYPE_EVENT_DOWNLOAD
	case tachibana.MessageTypeEventSystemStatus:
		return pb.MessageType_MESSAGE_TYPE_EVENT_SYSTEM_STATUS
	case tachibana.MessageTypeBusinessDay:
		return pb.MessageType_MESSAGE_TYPE_BUSINESS_DAY
	case tachibana.MessageTypeTickGroup:
		return pb.MessageType_MESSAGE_TYPE_TICK_GROUP
	case tachibana.MessageTypeEventOperationStatus:
		return pb.MessageType_MESSAGE_TYPE_EVENT_OPERATION_STATUS
	case tachibana.MessageTypeEventStockOperationStatus:
		return pb.MessageType_MESSAGE_TYPE_EVENT_STOCK_OPERATION_STATUS
	case tachibana.MessageTypeEventProductOperationStatus:
		return pb.MessageType_MESSAGE_TYPE_EVENT_PRODUCT_OPERATION_STATUS
	case tachibana.MessageTypeStockMaster:
		return pb.MessageType_MESSAGE_TYPE_STOCK_MASTER
	case tachibana.MessageTypeStockExchangeMaster:
		return pb.MessageType_MESSAGE_TYPE_STOCK_EXCHANGE_MASTER
	case tachibana.MessageTypeEventStockRegulation:
		return pb.MessageType_MESSAGE_TYPE_EVENT_STOCK_REGULATION
	case tachibana.MessageTypeEventFutureMaster:
		return pb.MessageType_MESSAGE_TYPE_EVENT_FUTURE_MASTER
	case tachibana.MessageTypeEventOptionMaster:
		return pb.MessageType_MESSAGE_TYPE_EVENT_OPTION_MASTER
	case tachibana.MessageTypeEventExchangeRegulation:
		return pb.MessageType_MESSAGE_TYPE_EVENT_EXCHANGE_REGULATION
	case tachibana.MessageTypeEventSubstitute:
		return pb.MessageType_MESSAGE_TYPE_EVENT_SUBSTITUTE
	case tachibana.MessageTypeEventDepositMaster:
		return pb.MessageType_MESSAGE_TYPE_EVENT_DEPOSIT_MASTER
	case tachibana.MessageTypeEventErrorReason:
		return pb.MessageType_MESSAGE_TYPE_EVENT_ERROR_REASON
	case tachibana.MessageTypeEventDownloadComplete:
		return pb.MessageType_MESSAGE_TYPE_EVENT_DOWNLOAD_COMPLETE
	case tachibana.MessageTypeMasterData:
		return pb.MessageType_MESSAGE_TYPE_MASTER_DATA
	case tachibana.MessageTypeMarketPrice:
		return pb.MessageType_MESSAGE_TYPE_MARKET_PRICE
	}
	return pb.MessageType_MESSAGE_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromAccountType(accountType tachibana.AccountType) pb.AccountType {
	switch accountType {
	case tachibana.AccountTypeSpecific:
		return pb.AccountType_ACCOUNT_TYPE_SPECIFIC
	case tachibana.AccountTypeGeneral:
		return pb.AccountType_ACCOUNT_TYPE_GENERAL
	case tachibana.AccountTypeNISA:
		return pb.AccountType_ACCOUNT_TYPE_NISA
	}
	return pb.AccountType_ACCOUNT_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromSpecificAccountType(specificAccountType tachibana.SpecificAccountType) pb.SpecificAccountType {
	switch specificAccountType {
	case tachibana.SpecificAccountTypeGeneral:
		return pb.SpecificAccountType_SPECIFIC_ACCOUNT_TYPE_GENERAL
	case tachibana.SpecificAccountTypeNothing:
		return pb.SpecificAccountType_SPECIFIC_ACCOUNT_TYPE_NOTHING
	case tachibana.SpecificAccountTypeWithholding:
		return pb.SpecificAccountType_SPECIFIC_ACCOUNT_TYPE_WITHHOLDING
	}
	return pb.SpecificAccountType_SPECIFIC_ACCOUNT_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) toAccountType(accountType pb.AccountType) tachibana.AccountType {
	switch accountType {
	case pb.AccountType_ACCOUNT_TYPE_SPECIFIC:
		return tachibana.AccountTypeSpecific
	case pb.AccountType_ACCOUNT_TYPE_GENERAL:
		return tachibana.AccountTypeGeneral
	case pb.AccountType_ACCOUNT_TYPE_NISA:
		return tachibana.AccountTypeNISA
	}
	return tachibana.AccountTypeUnspecified
}

func (t *tachibanaApi) toDeliveryAccountType(deliveryAccountType pb.DeliveryAccountType) tachibana.DeliveryAccountType {
	switch deliveryAccountType {
	case pb.DeliveryAccountType_DELIVERY_ACCOUNT_TYPE_UNUSED:
		return tachibana.DeliveryAccountTypeUnused
	case pb.DeliveryAccountType_DELIVERY_ACCOUNT_TYPE_SPECIFIC:
		return tachibana.DeliveryAccountTypeSpecific
	case pb.DeliveryAccountType_DELIVERY_ACCOUNT_TYPE_GENERAL:
		return tachibana.DeliveryAccountTypeGeneral
	case pb.DeliveryAccountType_DELIVERY_ACCOUNT_TYPE_NISA:
		return tachibana.DeliveryAccountTypeNISA
	}
	return tachibana.DeliveryAccountTypeUnspecified
}

func (t *tachibanaApi) toExchange(exchange pb.Exchange) tachibana.Exchange {
	switch exchange {
	case pb.Exchange_EXCHANGE_TOUSHOU:
		return tachibana.ExchangeToushou
	case pb.Exchange_EXCHANGE_MEISHOU:
		return tachibana.ExchangeMeishou
	case pb.Exchange_EXCHANGE_FUKUSHOU:
		return tachibana.ExchangeFukushou
	case pb.Exchange_EXCHANGE_SATSUSHOU:
		return tachibana.ExchangeSatsushou
	case pb.Exchange_EXCHANGE_STOPPING:
		return tachibana.ExchangeStopping
	}
	return tachibana.ExchangeUnspecified
}

func (t *tachibanaApi) toSide(side pb.Side) tachibana.Side {
	switch side {
	case pb.Side_SIDE_SELL:
		return tachibana.SideSell
	case pb.Side_SIDE_BUY:
		return tachibana.SideBuy
	case pb.Side_SIDE_DELIVERY:
		return tachibana.SideDelivery
	case pb.Side_SIDE_RECEIPT:
		return tachibana.SideReceipt
	}
	return tachibana.SideUnspecified
}

func (t *tachibanaApi) toExecutionTiming(executionTiming pb.ExecutionTiming) tachibana.ExecutionTiming {
	switch executionTiming {
	case pb.ExecutionTiming_EXECUTION_TIMING_NO_CHANGE:
		return tachibana.ExecutionTimingNoChange
	case pb.ExecutionTiming_EXECUTION_TIMING_NORMAL:
		return tachibana.ExecutionTimingNormal
	case pb.ExecutionTiming_EXECUTION_TIMING_OPENING:
		return tachibana.ExecutionTimingOpening
	case pb.ExecutionTiming_EXECUTION_TIMING_CLOSING:
		return tachibana.ExecutionTimingClosing
	case pb.ExecutionTiming_EXECUTION_TIMING_FUNARI:
		return tachibana.ExecutionTimingFunari
	}
	return tachibana.ExecutionTimingUnspecified
}

func (t *tachibanaApi) toTradeType(tradeType pb.TradeType) tachibana.TradeType {
	switch tradeType {
	case pb.TradeType_TRADE_TYPE_STOCK:
		return tachibana.TradeTypeStock
	case pb.TradeType_TRADE_TYPE_STANDARD_ENTRY:
		return tachibana.TradeTypeStandardEntry
	case pb.TradeType_TRADE_TYPE_STANDARD_EXIT:
		return tachibana.TradeTypeStandardExit
	case pb.TradeType_TRADE_TYPE_NEGOTIATE_ENTRY:
		return tachibana.TradeTypeNegotiateEntry
	case pb.TradeType_TRADE_TYPE_NEGOTIATE_EXIT:
		return tachibana.TradeTypeNegotiateExit
	}
	return tachibana.TradeTypeUnspecified
}

func (t *tachibanaApi) toStopOrderType(stopOrderType pb.StopOrderType) tachibana.StopOrderType {
	switch stopOrderType {
	case pb.StopOrderType_STOP_ORDER_TYPE_NORMAL:
		return tachibana.StopOrderTypeNormal
	case pb.StopOrderType_STOP_ORDER_TYPE_STOP:
		return tachibana.StopOrderTypeStop
	case pb.StopOrderType_STOP_ORDER_TYPE_OCO:
		return tachibana.StopOrderTypeOCO
	}
	return tachibana.StopOrderTypeUnspecified
}

func (t *tachibanaApi) toExitPositionType(exitPositionType pb.ExitPositionType) tachibana.ExitPositionType {
	switch exitPositionType {
	case pb.ExitPositionType_EXIT_POSITION_TYPE_NO_SELECTED:
		return tachibana.ExitPositionTypeNoSelected
	case pb.ExitPositionType_EXIT_POSITION_TYPE_UNUSED:
		return tachibana.ExitPositionTypeUnused
	case pb.ExitPositionType_EXIT_POSITION_TYPE_POSITION_NUMBER:
		return tachibana.ExitPositionTypePositionNumber
	case pb.ExitPositionType_EXIT_POSITION_TYPE_DAY_ASC:
		return tachibana.ExitPositionTypeDayAsc
	case pb.ExitPositionType_EXIT_POSITION_TYPE_PROFIT_DESC:
		return tachibana.ExitPositionTypeProfitDesc
	case pb.ExitPositionType_EXIT_POSITION_TYPE_PROFIT_ASC:
		return tachibana.ExitPositionTypeProfitAsc
	}
	return tachibana.ExitPositionTypeUnspecified
}

func (t *tachibanaApi) toOrderInquiryStatus(orderInquiryStatus pb.OrderInquiryStatus) tachibana.OrderInquiryStatus {
	switch orderInquiryStatus {
	case pb.OrderInquiryStatus_ORDER_INQUIRY_STATUS_IN_ORDER:
		return tachibana.OrderInquiryStatusInOrder
	case pb.OrderInquiryStatus_ORDER_INQUIRY_STATUS_DONE:
		return tachibana.OrderInquiryStatusDone
	case pb.OrderInquiryStatus_ORDER_INQUIRY_STATUS_PART:
		return tachibana.OrderInquiryStatusPart
	case pb.OrderInquiryStatus_ORDER_INQUIRY_STATUS_EDITABLE:
		return tachibana.OrderInquiryStatusEditable
	case pb.OrderInquiryStatus_ORDER_INQUIRY_STATUS_PART_IN_ORDER:
		return tachibana.OrderInquiryStatusInOrder
	}
	return tachibana.OrderInquiryStatusUnspecified
}

func (t *tachibanaApi) fromOrderInquiryStatus(orderInquiryStatus tachibana.OrderInquiryStatus) pb.OrderInquiryStatus {
	switch orderInquiryStatus {
	case tachibana.OrderInquiryStatusInOrder:
		return pb.OrderInquiryStatus_ORDER_INQUIRY_STATUS_IN_ORDER
	case tachibana.OrderInquiryStatusDone:
		return pb.OrderInquiryStatus_ORDER_INQUIRY_STATUS_DONE
	case tachibana.OrderInquiryStatusPart:
		return pb.OrderInquiryStatus_ORDER_INQUIRY_STATUS_PART
	case tachibana.OrderInquiryStatusEditable:
		return pb.OrderInquiryStatus_ORDER_INQUIRY_STATUS_EDITABLE
	case tachibana.OrderInquiryStatusPartInOrder:
		return pb.OrderInquiryStatus_ORDER_INQUIRY_STATUS_PART_IN_ORDER
	}
	return pb.OrderInquiryStatus_ORDER_INQUIRY_STATUS_UNSPECIFIED
}

func (t *tachibanaApi) fromExchange(exchange tachibana.Exchange) pb.Exchange {
	switch exchange {
	case tachibana.ExchangeToushou:
		return pb.Exchange_EXCHANGE_TOUSHOU
	case tachibana.ExchangeMeishou:
		return pb.Exchange_EXCHANGE_MEISHOU
	case tachibana.ExchangeFukushou:
		return pb.Exchange_EXCHANGE_FUKUSHOU
	case tachibana.ExchangeSatsushou:
		return pb.Exchange_EXCHANGE_SATSUSHOU
	case tachibana.ExchangeStopping:
		return pb.Exchange_EXCHANGE_STOPPING
	}
	return pb.Exchange_EXCHANGE_UNSPECIFIED
}

func (t *tachibanaApi) fromTradeType(tradeType tachibana.TradeType) pb.TradeType {
	switch tradeType {
	case tachibana.TradeTypeStock:
		return pb.TradeType_TRADE_TYPE_STOCK
	case tachibana.TradeTypeStandardEntry:
		return pb.TradeType_TRADE_TYPE_STANDARD_ENTRY
	case tachibana.TradeTypeStandardExit:
		return pb.TradeType_TRADE_TYPE_STANDARD_EXIT
	case tachibana.TradeTypeNegotiateEntry:
		return pb.TradeType_TRADE_TYPE_NEGOTIATE_ENTRY
	case tachibana.TradeTypeNegotiateExit:
		return pb.TradeType_TRADE_TYPE_NEGOTIATE_EXIT
	}
	return pb.TradeType_TRADE_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromExitTermType(exitTermType tachibana.ExitTermType) pb.ExitTermType {
	switch exitTermType {
	case tachibana.ExitTermTypeNoLimit:
		return pb.ExitTermType_EXIT_TERM_TYPE_NO_LIMIT
	case tachibana.ExitTermTypeStandardMargin6m:
		return pb.ExitTermType_EXIT_TERM_TYPE_STANDARD_MARGIN_6M
	case tachibana.ExitTermTypeStandardMarginNoLimit:
		return pb.ExitTermType_EXIT_TERM_TYPE_STANDARD_MARGIN_NO_LIMIT
	case tachibana.ExitTermTypeNegotiateMargin6m:
		return pb.ExitTermType_EXIT_TERM_TYPE_NEGOTIATE_MARGIN_6M
	case tachibana.ExitTermTypeNegotiateMarginNoLimit:
		return pb.ExitTermType_EXIT_TERM_TYPE_NEGOTIATE_MARGIN_NO_LIMIT
	}
	return pb.ExitTermType_EXIT_TERM_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromSide(side tachibana.Side) pb.Side {
	switch side {
	case tachibana.SideSell:
		return pb.Side_SIDE_SELL
	case tachibana.SideBuy:
		return pb.Side_SIDE_BUY
	case tachibana.SideDelivery:
		return pb.Side_SIDE_DELIVERY
	case tachibana.SideReceipt:
		return pb.Side_SIDE_RECEIPT
	}
	return pb.Side_SIDE_UNSPECIFIED
}

func (t *tachibanaApi) fromExecutionTiming(executionTiming tachibana.ExecutionTiming) pb.ExecutionTiming {
	switch executionTiming {
	case tachibana.ExecutionTimingNoChange:
		return pb.ExecutionTiming_EXECUTION_TIMING_NO_CHANGE
	case tachibana.ExecutionTimingNormal:
		return pb.ExecutionTiming_EXECUTION_TIMING_NORMAL
	case tachibana.ExecutionTimingOpening:
		return pb.ExecutionTiming_EXECUTION_TIMING_OPENING
	case tachibana.ExecutionTimingClosing:
		return pb.ExecutionTiming_EXECUTION_TIMING_CLOSING
	case tachibana.ExecutionTimingFunari:
		return pb.ExecutionTiming_EXECUTION_TIMING_FUNARI
	}
	return pb.ExecutionTiming_EXECUTION_TIMING_UNSPECIFIED
}

func (t *tachibanaApi) fromExecutionType(executionType tachibana.ExecutionType) pb.ExecutionType {
	switch executionType {
	case tachibana.ExecutionTypeUnused:
		return pb.ExecutionType_EXECUTION_TYPE_UNUSED
	case tachibana.ExecutionTypeMarket:
		return pb.ExecutionType_EXECUTION_TYPE_MARKET
	case tachibana.ExecutionTypeLimit:
		return pb.ExecutionType_EXECUTION_TYPE_LIMIT
	case tachibana.ExecutionTypeHigher:
		return pb.ExecutionType_EXECUTION_TYPE_HIGHER
	case tachibana.ExecutionTypeLower:
		return pb.ExecutionType_EXECUTION_TYPE_LOWER
	}
	return pb.ExecutionType_EXECUTION_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromStopOrderType(stopOrderType tachibana.StopOrderType) pb.StopOrderType {
	switch stopOrderType {
	case tachibana.StopOrderTypeNormal:
		return pb.StopOrderType_STOP_ORDER_TYPE_NORMAL
	case tachibana.StopOrderTypeStop:
		return pb.StopOrderType_STOP_ORDER_TYPE_STOP
	case tachibana.StopOrderTypeOCO:
		return pb.StopOrderType_STOP_ORDER_TYPE_OCO
	}
	return pb.StopOrderType_STOP_ORDER_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromTriggerType(triggerType tachibana.TriggerType) pb.TriggerType {
	switch triggerType {
	case tachibana.TriggerTypeNoFired:
		return pb.TriggerType_TRIGGER_TYPE_NO_FIRED
	case tachibana.TriggerTypeAuto:
		return pb.TriggerType_TRIGGER_TYPE_AUTO
	case tachibana.TriggerTypeManualOrder:
		return pb.TriggerType_TRIGGER_TYPE_MANUAL_ORDER
	case tachibana.TriggerTypeManualExpired:
		return pb.TriggerType_TRIGGER_TYPE_MANUAL_EXPIRED
	}
	return pb.TriggerType_TRIGGER_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromExitPositionType(exitPositionType tachibana.ExitPositionType) pb.ExitPositionType {
	switch exitPositionType {
	case tachibana.ExitPositionTypeNoSelected:
		return pb.ExitPositionType_EXIT_POSITION_TYPE_NO_SELECTED
	case tachibana.ExitPositionTypeUnused:
		return pb.ExitPositionType_EXIT_POSITION_TYPE_UNUSED
	case tachibana.ExitPositionTypePositionNumber:
		return pb.ExitPositionType_EXIT_POSITION_TYPE_POSITION_NUMBER
	case tachibana.ExitPositionTypeDayAsc:
		return pb.ExitPositionType_EXIT_POSITION_TYPE_DAY_ASC
	case tachibana.ExitPositionTypeProfitDesc:
		return pb.ExitPositionType_EXIT_POSITION_TYPE_PROFIT_DESC
	case tachibana.ExitPositionTypeProfitAsc:
		return pb.ExitPositionType_EXIT_POSITION_TYPE_PROFIT_ASC
	}
	return pb.ExitPositionType_EXIT_POSITION_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromPartContractType(partContractType tachibana.PartContractType) pb.PartContractType {
	switch partContractType {
	case tachibana.PartContractTypeUnused:
		return pb.PartContractType_PART_CONTRACT_TYPE_UNUSED
	case tachibana.PartContractTypePart:
		return pb.PartContractType_PART_CONTRACT_TYPE_PART
	}
	return pb.PartContractType_PART_CONTRACT_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromOrderStatus(orderStatus tachibana.OrderStatus) pb.OrderStatus {
	switch orderStatus {
	case tachibana.OrderStatusReceived:
		return pb.OrderStatus_ORDER_STATUS_RECEIVED
	case tachibana.OrderStatusInOrder:
		return pb.OrderStatus_ORDER_STATUS_IN_ORDER
	case tachibana.OrderStatusError:
		return pb.OrderStatus_ORDER_STATUS_ERROR
	case tachibana.OrderStatusInCorrect:
		return pb.OrderStatus_ORDER_STATUS_IN_CORRECT
	case tachibana.OrderStatusCorrected:
		return pb.OrderStatus_ORDER_STATUS_CORRECTED
	case tachibana.OrderStatusCorrectFailed:
		return pb.OrderStatus_ORDER_STATUS_CORRECT_FAILED
	case tachibana.OrderStatusInCancel:
		return pb.OrderStatus_ORDER_STATUS_IN_CANCEL
	case tachibana.OrderStatusCanceled:
		return pb.OrderStatus_ORDER_STATUS_CANCELED
	case tachibana.OrderStatusCancelFailed:
		return pb.OrderStatus_ORDER_STATUS_CANCEL_FAILED
	case tachibana.OrderStatusPart:
		return pb.OrderStatus_ORDER_STATUS_PART
	case tachibana.OrderStatusDone:
		return pb.OrderStatus_ORDER_STATUS_DONE
	case tachibana.OrderStatusPartExpired:
		return pb.OrderStatus_ORDER_STATUS_PART_EXPIRED
	case tachibana.OrderStatusExpired:
		return pb.OrderStatus_ORDER_STATUS_EXPIRED
	case tachibana.OrderStatusWait:
		return pb.OrderStatus_ORDER_STATUS_WAIT
	case tachibana.OrderStatusInvalid:
		return pb.OrderStatus_ORDER_STATUS_INVALID
	case tachibana.OrderStatusTrigger:
		return pb.OrderStatus_ORDER_STATUS_TRIGGER
	case tachibana.OrderStatusTriggered:
		return pb.OrderStatus_ORDER_STATUS_TRIGGERED
	case tachibana.OrderStatusTriggerFailed:
		return pb.OrderStatus_ORDER_STATUS_TRIGGER_FAILED
	case tachibana.OrderStatusCarryOverFailed:
		return pb.OrderStatus_ORDER_STATUS_CARRYOVER_FAILED
	case tachibana.OrderStatusInOrderStop:
		return pb.OrderStatus_ORDER_STATUS_IN_ORDER_STOP
	}
	return pb.OrderStatus_ORDER_STATUS_UNSPECIFIED
}

func (t *tachibanaApi) fromContractStatus(contractStatus tachibana.ContractStatus) pb.ContractStatus {
	switch contractStatus {
	case tachibana.ContractStatusInOrder:
		return pb.ContractStatus_CONTRACT_STATUS_IN_ORDER
	case tachibana.ContractStatusPart:
		return pb.ContractStatus_CONTRACT_STATUS_PART
	case tachibana.ContractStatusDone:
		return pb.ContractStatus_CONTRACT_STATUS_DONE
	case tachibana.ContractStatusInContract:
		return pb.ContractStatus_CONTRACT_STATUS_IN_CONTRACT
	}
	return pb.ContractStatus_CONTRACT_STATUS_UNSPECIFIED
}

func (t *tachibanaApi) fromCarryOverType(carryOverType tachibana.CarryOverType) pb.CarryOverType {
	switch carryOverType {
	case tachibana.CarryOverTypeToday:
		return pb.CarryOverType_CARRY_OVER_TYPE_TODAY
	case tachibana.CarryOverTypeCarry:
		return pb.CarryOverType_CARRY_OVER_TYPE_CARRY
	case tachibana.CarryOverTypeInvalid:
		return pb.CarryOverType_CARRY_OVER_TYPE_INVALID
	}
	return pb.CarryOverType_CARRY_OVER_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromCorrectCancelType(correctCancelType tachibana.CorrectCancelType) pb.CorrectCancelType {
	switch correctCancelType {
	case tachibana.CorrectCancelTypeCorrectable:
		return pb.CorrectCancelType_CORRECT_CANCEL_TYPE_CORRECTABLE
	case tachibana.CorrectCancelTypeCancelable:
		return pb.CorrectCancelType_CORRECT_CANCEL_TYPE_CANCELABLE
	case tachibana.CorrectCancelTypeInvalid:
		return pb.CorrectCancelType_CORRECT_CANCEL_TYPE_INVALID
	}
	return pb.CorrectCancelType_CORRECT_CANCEL_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromChannel(channel tachibana.Channel) pb.Channel {
	switch channel {
	case tachibana.ChannelMeet:
		return pb.Channel_CHANNEL_MEET
	case tachibana.ChannelPC:
		return pb.Channel_CHANNEL_PC
	case tachibana.ChannelCallCenter:
		return pb.Channel_CHANNEL_CALL_CENTER
	case tachibana.ChannelCallCenter2:
		return pb.Channel_CHANNEL_CALL_CENTER2
	case tachibana.ChannelCallCenter3:
		return pb.Channel_CHANNEL_CALL_CENTER3
	case tachibana.ChannelMobile:
		return pb.Channel_CHANNEL_MOBILE
	case tachibana.ChannelRich:
		return pb.Channel_CHANNEL_RICH
	case tachibana.ChannelSmartPhone:
		return pb.Channel_CHANNEL_SMARTPHONE
	case tachibana.ChannelIPadApp:
		return pb.Channel_CHANNEL_IPAD_APP
	case tachibana.ChannelHost:
		return pb.Channel_CHANNEL_HOST
	}
	return pb.Channel_CHANNEL_UNSPECIFIED
}

func (t *tachibanaApi) toStockMasterColumn(column pb.StockMasterColumn) tachibana.StockMasterColumn {
	switch column {
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_ISSUE_CODE:
		return tachibana.StockMasterColumnIssueCode
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_NAME:
		return tachibana.StockMasterColumnName
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_SHORT_NAME:
		return tachibana.StockMasterColumnShortName
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_KANA:
		return tachibana.StockMasterColumnKana
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_ALPHABET:
		return tachibana.StockMasterColumnAlphabet
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_SPECIFIC_TARGET:
		return tachibana.StockMasterColumnSpecificTarget
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_TAX_FREE:
		return tachibana.StockMasterColumnTaxFree
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_SHARED_STOCKS:
		return tachibana.StockMasterColumnSharedStocks
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_EX_RIGHT:
		return tachibana.StockMasterColumnExRight
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_LAST_RIGHT_DAY:
		return tachibana.StockMasterColumnLastRightDay
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_LISTING_TYPE:
		return tachibana.StockMasterColumnListingType
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_RELEASE_TRADING_DATE:
		return tachibana.StockMasterColumnReleaseTradingDate
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_TRADING_DATE:
		return tachibana.StockMasterColumnTradingDate
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_TRADING_UNIT:
		return tachibana.StockMasterColumnTradingUnit
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_NEXT_TRADING_UNIT:
		return tachibana.StockMasterColumnNextTradingUnit
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_STOP_TRADING_TYPE:
		return tachibana.StockMasterColumnStopTradingType
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_START_PUBLICATION_DATE:
		return tachibana.StockMasterColumnStartPublicationDate
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_LAST_PUBLICATION_DATE:
		return tachibana.StockMasterColumnLastPublicationDate
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_SETTLEMENT_TYPE:
		return tachibana.StockMasterColumnSettlementType
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_SETTLEMENT_DATE:
		return tachibana.StockMasterColumnSettlementDate
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_LISTING_DATE:
		return tachibana.StockMasterColumnListingDate
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_EXPIRE_DATE_2_TYPE:
		return tachibana.StockMasterColumnExpireDate2Type
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_LARGE_UNIT:
		return tachibana.StockMasterColumnLargeUnit
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_LARGE_AMOUNT:
		return tachibana.StockMasterColumnLargeAmount
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_OUTPUT_TICKET_TYPE:
		return tachibana.StockMasterColumnOutputTicketType
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_DEPOSIT_AMOUNT:
		return tachibana.StockMasterColumnDepositAmount
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_DEPOSIT_VALUATION:
		return tachibana.StockMasterColumnDepositValuation
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_ORGANIZATION_TYPE:
		return tachibana.StockMasterColumnOrganizationType
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_PROVISIONAL_TYPE:
		return tachibana.StockMasterColumnProvisionalType
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_PRIMARY_EXCHANGE:
		return tachibana.StockMasterColumnPrimaryExchange
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_INDEFINITE_PERIOD_TYPE:
		return tachibana.StockMasterColumnIndefinitePeriodType
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_INDUSTRY_CODE:
		return tachibana.StockMasterColumnIndustryCode
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_INDUSTRY_NAME:
		return tachibana.StockMasterColumnIndustryName
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_SOR_TARGET_TYPE:
		return tachibana.StockMasterColumnSORTargetType
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_CREATE_DATETIME:
		return tachibana.StockMasterColumnCreateDateTime
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_UPDATE_DATETIME:
		return tachibana.StockMasterColumnUpdateDateTime
	case pb.StockMasterColumn_STOCK_MASTER_COLUMN_UPDATE_NUMBER:
		return tachibana.StockMasterColumnUpdateNumber
	}
	return ""
}

func (t *tachibanaApi) fromTaxFree(taxFree tachibana.TaxFree) pb.TaxFree {
	switch taxFree {
	case tachibana.TaxFreeUnUsed:
		return pb.TaxFree_TAX_FREE_UN_USED
	case tachibana.TaxFreeValid:
		return pb.TaxFree_TAX_FREE_VALID
	}
	return pb.TaxFree_TAX_FREE_UNSPECIFIED
}

func (t *tachibanaApi) fromExRightType(exRightType tachibana.ExRightType) pb.ExRightType {
	switch exRightType {
	case tachibana.ExRightTypeNothing:
		return pb.ExRightType_EX_RIGHT_TYPE_NOTHING
	case tachibana.ExRightTypeStockSplit:
		return pb.ExRightType_EX_RIGHT_TYPE_STOCK_SPLIT
	case tachibana.ExRightTypeDividend:
		return pb.ExRightType_EX_RIGHT_TYPE_DIVIDEND
	case tachibana.ExRightTypeOther:
		return pb.ExRightType_EX_RIGHT_TYPE_OTHER
	case tachibana.ExRightTypeDividendAndOther:
		return pb.ExRightType_EX_RIGHT_TYPE_DIVIDEND_AND_OTHER
	case tachibana.ExRightTypeStockSplitAndOther:
		return pb.ExRightType_EX_RIGHT_TYPE_STOCK_SPLIT_AND_OTHER
	case tachibana.ExRightTypeStockSplitAndOtherMiddle:
		return pb.ExRightType_EX_RIGHT_TYPE_STOCK_SPLIT_AND_OTHER_MIDDLE
	}
	return pb.ExRightType_EX_RIGHT_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromListingType(listingType tachibana.ListingType) pb.ListingType {
	switch listingType {
	case tachibana.ListingTypeUnUsed:
		return pb.ListingType_LISTING_TYPE_UN_USED
	case tachibana.ListingTypeNewest:
		return pb.ListingType_LISTING_TYPE_NEWEST
	case tachibana.ListingTypeGeneral:
		return pb.ListingType_LISTING_TYPE_GENERAL
	case tachibana.ListingTypeRight:
		return pb.ListingType_LISTING_TYPE_RIGHT
	case tachibana.ListingTypeOffer:
		return pb.ListingType_LISTING_TYPE_OFFER
	case tachibana.ListingTypeSelling:
		return pb.ListingType_LISTING_TYPE_SELLING
	case tachibana.ListingTypeOpenBuy:
		return pb.ListingType_LISTING_TYPE_OPEN_BUY
	case tachibana.ListingTypeTransmission:
		return pb.ListingType_LISTING_TYPE_TRANSMISSION
	}
	return pb.ListingType_LISTING_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromStopTradingType(stopTradingType tachibana.StopTradingType) pb.StopTradingType {
	switch stopTradingType {
	case tachibana.StopTradingTypeUnUsed:
		return pb.StopTradingType_STOP_TRADING_TYPE_UN_USED
	case tachibana.StopTradingTypeRelease:
		return pb.StopTradingType_STOP_TRADING_TYPE_RELEASE
	case tachibana.StopTradingTypeStopping:
		return pb.StopTradingType_STOP_TRADING_TYPE_STOPPING
	}
	return pb.StopTradingType_STOP_TRADING_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromSettlementType(settlementType tachibana.SettlementType) pb.SettlementType {
	switch settlementType {
	case tachibana.SettlementTypeCapitalIncrease:
		return pb.SettlementType_SETTLEMENT_TYPE_CAPITAL_INCREASE
	case tachibana.SettlementTypeSplit:
		return pb.SettlementType_SETTLEMENT_TYPE_SPLIT
	case tachibana.SettlementTypeAssignment:
		return pb.SettlementType_SETTLEMENT_TYPE_ASSIGNMENT
	}
	return pb.SettlementType_SETTLEMENT_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) toStockExchangeMasterColumn(column pb.StockExchangeMasterColumn) tachibana.StockExchangeMasterColumn {
	switch column {
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_ISSUE_CODE:
		return tachibana.StockExchangeMasterColumnIssueCode
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_EXCHANGE:
		return tachibana.StockExchangeMasterColumnExchange
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_STOCK_SYSTEM_TYPE:
		return tachibana.StockExchangeMasterColumnStockSystemType
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_UNDER_LIMIT_PRICE:
		return tachibana.StockExchangeMasterColumnUnderLimitPrice
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_UPPER_LIMIT_PRICE:
		return tachibana.StockExchangeMasterColumnUpperLimitPrice
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_SYMBOL_CATEGORY:
		return tachibana.StockExchangeMasterColumnSymbolCategory
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_LIMIT_PRICE_EXCHANGE:
		return tachibana.StockExchangeMasterColumnLimitPriceExchange
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_MARGIN_TYPE:
		return tachibana.StockExchangeMasterColumnMarginType
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_LISTING_DATE:
		return tachibana.StockExchangeMasterColumnListingDate
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_LIMIT_PRICE_DATE:
		return tachibana.StockExchangeMasterColumnLimitPriceDate
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_LIMIT_PRICE_CATEGORY:
		return tachibana.StockExchangeMasterColumnLimitPriceCategory
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_LIMIT_PRICE_VALUE:
		return tachibana.StockExchangeMasterColumnLimitPriceValue
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_CONFIRM_LIMIT_PRICE:
		return tachibana.StockExchangeMasterColumnConfirmLimitPrice
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_SECTION:
		return tachibana.StockExchangeMasterColumnSection
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_PREV_CLOSE_PRICE:
		return tachibana.StockExchangeMasterColumnPrevClosePrice
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_CALCULATE_LIMIT_PRICE_EXCHANGE:
		return tachibana.StockExchangeMasterColumnLimitPriceExchange
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_REGULATION1:
		return tachibana.StockExchangeMasterColumnRegulation1
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_REGULATION2:
		return tachibana.StockExchangeMasterColumnRegulation2
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_SECTION_TYPE:
		return tachibana.StockExchangeMasterColumnSectionType
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_DELISTING_DATE:
		return tachibana.StockExchangeMasterColumnDelistingDate
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_TRADING_UNIT:
		return tachibana.StockExchangeMasterColumnTradingUnit
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_NEXT_TRADING_UNIT:
		return tachibana.StockExchangeMasterColumnNextTradingUnit
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_TICK_GROUP_TYPE:
		return tachibana.StockExchangeMasterColumnTickGroupType
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_NEXT_TICK_GROUP_TYPE:
		return tachibana.StockExchangeMasterColumnNextTickGroupType
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_INFORMATION_SOURCE:
		return tachibana.StockExchangeMasterColumnInformationSource
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_INFORMATION_CODE:
		return tachibana.StockExchangeMasterColumnInformationCode
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_OFFER_PRICE:
		return tachibana.StockExchangeMasterColumnOfferPrice
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_CREATE_DATETIME:
		return tachibana.StockExchangeMasterColumnCreateDateTime
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_UPDATE_DATETIME:
		return tachibana.StockExchangeMasterColumnUpdateDateTime
	case pb.StockExchangeMasterColumn_STOCK_EXCHANGE_MASTER_COLUMN_UPDATE_NUMBER:
		return tachibana.StockExchangeMasterColumnUpdateNumber
	}
	return ""
}

func (t *tachibanaApi) fromMarginType(marginType tachibana.MarginType) pb.MarginType {
	switch marginType {
	case tachibana.MarginTypeMarginTrading:
		return pb.MarginType_MARGIN_TYPE_MARGIN_TRADING
	case tachibana.MarginTypeStandard:
		return pb.MarginType_MARGIN_TYPE_STANDARD
	case tachibana.MarginTypeNegotiate:
		return pb.MarginType_MARGIN_TYPE_NEGOTIATE
	}
	return pb.MarginType_MARGIN_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromTickGroupType(tickGroupType tachibana.TickGroupType) pb.TickGroupType {
	switch tickGroupType {
	case tachibana.TickGroupTypeStock1:
		return pb.TickGroupType_TICK_GROUP_TYPE_STOCK1
	case tachibana.TickGroupTypeStock2:
		return pb.TickGroupType_TICK_GROUP_TYPE_STOCK2
	case tachibana.TickGroupTypeStock3:
		return pb.TickGroupType_TICK_GROUP_TYPE_STOCK3
	case tachibana.TickGroupTypeBond1:
		return pb.TickGroupType_TICK_GROUP_TYPE_BOND1
	case tachibana.TickGroupTypeBond2:
		return pb.TickGroupType_TICK_GROUP_TYPE_BOND2
	case tachibana.TickGroupTypeNK225:
		return pb.TickGroupType_TICK_GROUP_TYPE_NK225
	case tachibana.TickGroupTypeNK225Mini:
		return pb.TickGroupType_TICK_GROUP_TYPE_NK225_MINI
	case tachibana.TickGroupTypeNK225OP:
		return pb.TickGroupType_TICK_GROUP_TYPE_NK225_OP
	}
	return pb.TickGroupType_TICK_GROUP_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromMarketPriceColumn(column pb.MarketPriceColumn) tachibana.MarketPriceColumn {
	switch column {
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_SECTION:
		return tachibana.MarketPriceColumnSection
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_CURRENT_PRICE:
		return tachibana.MarketPriceColumnCurrentPrice
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_CURRENT_PRICE_TIME:
		return tachibana.MarketPriceColumnCurrentPriceTime
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_CHANGE_PRICE_TYPE:
		return tachibana.MarketPriceColumnChangePriceType
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_PREV_DAY_RATIO:
		return tachibana.MarketPriceColumnPrevDayRatio
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_PREV_DAY_PERCENT:
		return tachibana.MarketPriceColumnPrevDayPercent
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_OPEN_PRICE:
		return tachibana.MarketPriceColumnOpenPrice
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_OPEN_PRICE_TIME:
		return tachibana.MarketPriceColumnOpenPriceTime
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_HIGH_PRICE:
		return tachibana.MarketPriceColumnHighPrice
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_HIGH_PRICE_TIME:
		return tachibana.MarketPriceColumnHighPriceTime
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_LOW_PRICE:
		return tachibana.MarketPriceColumnLowPrice
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_LOW_PRICE_TIME:
		return tachibana.MarketPriceColumnLowPriceTime
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_VOLUME:
		return tachibana.MarketPriceColumnVolume
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_SIGN:
		return tachibana.MarketPriceColumnAskSign
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_PRICE:
		return tachibana.MarketPriceColumnAskPrice
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY:
		return tachibana.MarketPriceColumnAskQuantity
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_SIGN:
		return tachibana.MarketPriceColumnBidSign
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_PRICE:
		return tachibana.MarketPriceColumnBidPrice
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY:
		return tachibana.MarketPriceColumnBidQuantity
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_EX_RIGHT_TYPE:
		return tachibana.MarketPriceColumnExRightType
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_DISCONTINUITY_TYPE:
		return tachibana.MarketPriceColumnDiscontinuityType
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_STOP_HIGH:
		return tachibana.MarketPriceColumnStopHigh
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_STOP_LOW:
		return tachibana.MarketPriceColumnStopLow
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_TRADING_AMOUNT:
		return tachibana.MarketPriceColumnTradingAmount
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY_MARKET:
		return tachibana.MarketPriceColumnAskQuantityMarket
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY_MARKET:
		return tachibana.MarketPriceColumnBidQuantityMarket
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY_OVER:
		return tachibana.MarketPriceColumnAskQuantityOver
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY10:
		return tachibana.MarketPriceColumnAskQuantity10
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_PRICE10:
		return tachibana.MarketPriceColumnAskPrice10
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY9:
		return tachibana.MarketPriceColumnAskQuantity9
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_PRICE9:
		return tachibana.MarketPriceColumnAskPrice9
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY8:
		return tachibana.MarketPriceColumnAskQuantity8
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_PRICE8:
		return tachibana.MarketPriceColumnAskPrice8
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY7:
		return tachibana.MarketPriceColumnAskQuantity7
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_PRICE7:
		return tachibana.MarketPriceColumnAskPrice7
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY6:
		return tachibana.MarketPriceColumnAskQuantity6
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_PRICE6:
		return tachibana.MarketPriceColumnAskPrice6
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY5:
		return tachibana.MarketPriceColumnAskQuantity5
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_PRICE5:
		return tachibana.MarketPriceColumnAskPrice5
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY4:
		return tachibana.MarketPriceColumnAskQuantity4
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_PRICE4:
		return tachibana.MarketPriceColumnAskPrice4
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY3:
		return tachibana.MarketPriceColumnAskQuantity3
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_PRICE3:
		return tachibana.MarketPriceColumnAskPrice3
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY2:
		return tachibana.MarketPriceColumnAskQuantity2
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_PRICE2:
		return tachibana.MarketPriceColumnAskPrice2
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_QUANTITY1:
		return tachibana.MarketPriceColumnAskQuantity1
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_ASK_PRICE1:
		return tachibana.MarketPriceColumnAskPrice1
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY1:
		return tachibana.MarketPriceColumnBidQuantity1
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_PRICE1:
		return tachibana.MarketPriceColumnBidPrice1
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY2:
		return tachibana.MarketPriceColumnBidQuantity2
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_PRICE2:
		return tachibana.MarketPriceColumnBidPrice2
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY3:
		return tachibana.MarketPriceColumnBidQuantity3
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_PRICE3:
		return tachibana.MarketPriceColumnBidPrice3
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY4:
		return tachibana.MarketPriceColumnBidQuantity4
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_PRICE4:
		return tachibana.MarketPriceColumnBidPrice4
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY5:
		return tachibana.MarketPriceColumnBidQuantity5
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_PRICE5:
		return tachibana.MarketPriceColumnBidPrice5
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY6:
		return tachibana.MarketPriceColumnBidQuantity6
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_PRICE6:
		return tachibana.MarketPriceColumnBidPrice6
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY7:
		return tachibana.MarketPriceColumnBidQuantity7
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_PRICE7:
		return tachibana.MarketPriceColumnBidPrice7
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY8:
		return tachibana.MarketPriceColumnBidQuantity8
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_PRICE8:
		return tachibana.MarketPriceColumnBidPrice8
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY9:
		return tachibana.MarketPriceColumnBidQuantity9
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_PRICE9:
		return tachibana.MarketPriceColumnBidPrice9
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY10:
		return tachibana.MarketPriceColumnBidQuantity10
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_PRICE10:
		return tachibana.MarketPriceColumnBidPrice10
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_BID_QUANTITY_UNDER:
		return tachibana.MarketPriceColumnBidQuantityUnder
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_VWAP:
		return tachibana.MarketPriceColumnVWAP
	case pb.MarketPriceColumn_MARKET_PRICE_COLUMN_PRP:
		return tachibana.MarketPriceColumnPRP
	}
	return ""
}

func (t *tachibanaApi) fromChangePriceType(changePriceType tachibana.ChangePriceType) pb.ChangePriceType {
	switch changePriceType {
	case tachibana.ChangePriceTypeNoChange:
		return pb.ChangePriceType_CHANGE_PRICE_TYPE_NO_CHANGE
	case tachibana.ChangePriceTypeEqual:
		return pb.ChangePriceType_CHANGE_PRICE_TYPE_EQUAL
	case tachibana.ChangePriceTypeRise:
		return pb.ChangePriceType_CHANGE_PRICE_TYPE_RISE
	case tachibana.ChangePriceTypeDown:
		return pb.ChangePriceType_CHANGE_PRICE_TYPE_DOWN
	case tachibana.ChangePriceTypeOpenAfterStopping:
		return pb.ChangePriceType_CHANGE_PRICE_TYPE_OPEN_AFTER_STOPPING
	case tachibana.ChangePriceTypeZaraba:
		return pb.ChangePriceType_CHANGE_PRICE_TYPE_ZARABA
	case tachibana.ChangePriceTypeClose:
		return pb.ChangePriceType_CHANGE_PRICE_TYPE_CLOSE
	case tachibana.ChangePriceTypeCloseAtStopping:
		return pb.ChangePriceType_CHANGE_PRICE_TYPE_CLOSE_AT_STOPPING
	case tachibana.ChangePriceTypeStopping:
		return pb.ChangePriceType_CHANGE_PRICE_TYPE_STOPPING
	}
	return pb.ChangePriceType_CHANGE_PRICE_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromIndicationPriceType(indicationPriceType tachibana.IndicationPriceType) pb.IndicationPriceType {
	switch indicationPriceType {
	case tachibana.IndicationPriceTypeNoChange:
		return pb.IndicationPriceType_INDICATION_PRICE_TYPE_NO_CHANGE
	case tachibana.IndicationPriceTypeGeneral:
		return pb.IndicationPriceType_INDICATION_PRICE_TYPE_GENERAL
	case tachibana.IndicationPriceTypeSpecific:
		return pb.IndicationPriceType_INDICATION_PRICE_TYPE_SPECIFIC
	case tachibana.IndicationPriceTypeBeforeOpening:
		return pb.IndicationPriceType_INDICATION_PRICE_TYPE_BEFORE_OPENING
	case tachibana.IndicationPriceTypeBeforeClosing:
		return pb.IndicationPriceType_INDICATION_PRICE_TYPE_BEFORE_CLOSING
	case tachibana.IndicationPriceTypeContinuance:
		return pb.IndicationPriceType_INDICATION_PRICE_TYPE_CONTINUANCE
	case tachibana.IndicationPriceTypeContinuanceBeforeClosing:
		return pb.IndicationPriceType_INDICATION_PRICE_TYPE_CONTINUANCE_BEFORE_CLOSING
	case tachibana.IndicationPriceTypeMoving:
		return pb.IndicationPriceType_INDICATION_PRICE_TYPE_MOVING
	}
	return pb.IndicationPriceType_INDICATION_PRICE_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromCurrentPriceType(currentPriceType tachibana.CurrentPriceType) pb.CurrentPriceType {
	switch currentPriceType {
	case tachibana.CurrentPriceTypeNoChange:
		return pb.CurrentPriceType_CURRENT_PRICE_TYPE_NO_CHANGE
	case tachibana.CurrentPriceTypeStopHigh:
		return pb.CurrentPriceType_CURRENT_PRICE_TYPE_STOP_HIGH
	case tachibana.CurrentPriceTypeStopLow:
		return pb.CurrentPriceType_CURRENT_PRICE_TYPE_STOP_LOW
	}
	return pb.CurrentPriceType_CURRENT_PRICE_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromDayKey(dayKey tachibana.DayKey) pb.DayKey {
	switch dayKey {
	case tachibana.DayKeyToday:
		return pb.DayKey_DAY_KEY_TODAY
	case tachibana.DayKeyNextDay:
		return pb.DayKey_DAY_KEY_NEXT_DAY
	}
	return pb.DayKey_DAY_KEY_UNSPECIFIED
}

func (t *tachibanaApi) toEventType(eventType pb.EventType) tachibana.EventType {
	switch eventType {
	case pb.EventType_EVENT_TYPE_ERROR_STATUS:
		return tachibana.EventTypeErrorStatus
	case pb.EventType_EVENT_TYPE_KEEPALIVE:
		return tachibana.EventTypeKeepAlive
	case pb.EventType_EVENT_TYPE_MARKET_PRICE:
		return tachibana.EventTypeMarketPrice
	case pb.EventType_EVENT_TYPE_CONTRACT:
		return tachibana.EventTypeContract
	case pb.EventType_EVENT_TYPE_NEWS:
		return tachibana.EventTypeNews
	case pb.EventType_EVENT_TYPE_SYSTEM_STATUS:
		return tachibana.EventTypeSystemStatus
	case pb.EventType_EVENT_TYPE_OPERATION_STATUS:
		return tachibana.EventTypeOperationStatus
	}
	return tachibana.EventTypeUnspecified
}

func (t *tachibanaApi) fromEventType(eventType tachibana.EventType) pb.EventType {
	switch eventType {
	case tachibana.EventTypeErrorStatus:
		return pb.EventType_EVENT_TYPE_ERROR_STATUS
	case tachibana.EventTypeMarketPrice:
		return pb.EventType_EVENT_TYPE_MARKET_PRICE
	case tachibana.EventTypeContract:
		return pb.EventType_EVENT_TYPE_CONTRACT
	case tachibana.EventTypeNews:
		return pb.EventType_EVENT_TYPE_NEWS
	case tachibana.EventTypeSystemStatus:
		return pb.EventType_EVENT_TYPE_SYSTEM_STATUS
	case tachibana.EventTypeOperationStatus:
		return pb.EventType_EVENT_TYPE_OPERATION_STATUS
	}
	return pb.EventType_EVENT_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromStreamOrderType(streamOrderType tachibana.StreamOrderType) pb.StreamOrderType {
	switch streamOrderType {
	case tachibana.StreamOrderTypeReceiveOrder:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_RECEIVE_ORDER
	case tachibana.StreamOrderTypeReceiveCorrect:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_RECEIVE_CORRECT
	case tachibana.StreamOrderTypeReceiveCancel:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_RECEIVE_CANCEL
	case tachibana.StreamOrderTypeReceiveError:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_RECEIVE_ERROR
	case tachibana.StreamOrderTypeReceiveCorrectError:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_RECEIVE_CORRECT_ERROR
	case tachibana.StreamOrderTypeReceiveCancelError:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_RECEIVE_CANCEL_ERROR
	case tachibana.StreamOrderTypeOrderError:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_ORDER_ERROR
	case tachibana.StreamOrderTypeCorrectError:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_CORRECT_ERROR
	case tachibana.StreamOrderTypeCancelError:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_CANCEL_ERROR
	case tachibana.StreamOrderTypeCorrected:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_CORRECTED
	case tachibana.StreamOrderTypeCanceled:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_CANCELED
	case tachibana.StreamOrderTypeContract:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_CONTRACT
	case tachibana.StreamOrderTypeExpire:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_EXPIRE
	case tachibana.StreamOrderTypeExpireContinue:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_EXPIRE_CONTINUE
	case tachibana.StreamOrderTypeCancelContract:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_CANCEL_CONTRACT
	case tachibana.StreamOrderTypeCarryOver:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_CARRYOVER
	case tachibana.StreamOrderTypeReceived:
		return pb.StreamOrderType_STREAM_ORDER_TYPE_RECEIVED
	}
	return pb.StreamOrderType_STREAM_ORDER_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromProductType(productType tachibana.ProductType) pb.ProductType {
	switch productType {
	case tachibana.ProductTypeStock:
		return pb.ProductType_PRODUCT_TYPE_STOCK
	case tachibana.ProductTypeFuture:
		return pb.ProductType_PRODUCT_TYPE_FUTURE
	case tachibana.ProductTypeOption:
		return pb.ProductType_PRODUCT_TYPE_OPTION
	}
	return pb.ProductType_PRODUCT_TYPE_UNSPECIFIED
}

func (t *tachibanaApi) fromStreamOrderStatus(streamOrderStatus tachibana.StreamOrderStatus) pb.StreamOrderStatus {
	switch streamOrderStatus {
	case tachibana.StreamOrderStatusNew:
		return pb.StreamOrderStatus_STREAM_ORDER_STATUS_NEW
	case tachibana.StreamOrderStatusReceived:
		return pb.StreamOrderStatus_STREAM_ORDER_STATUS_RECEIVED
	case tachibana.StreamOrderStatusError:
		return pb.StreamOrderStatus_STREAM_ORDER_STATUS_ERROR
	case tachibana.StreamOrderStatusPartExpired:
		return pb.StreamOrderStatus_STREAM_ORDER_STATUS_PART_EXPIRED
	case tachibana.StreamOrderStatusExpired:
		return pb.StreamOrderStatus_STREAM_ORDER_STATUS_EXPIRED
	case tachibana.StreamOrderStatusCarryOverExpired:
		return pb.StreamOrderStatus_STREAM_ORDER_STATUS_CARRY_OVER_EXPIRED
	}
	return pb.StreamOrderStatus_STREAM_ORDER_STATUS_UNSPECIFIED
}

func (t *tachibanaApi) fromCancelOrderStatus(cancelOrderStatus tachibana.CancelOrderStatus) pb.CancelOrderStatus {
	switch cancelOrderStatus {
	case tachibana.CancelOrderStatusNoCorrect:
		return pb.CancelOrderStatus_CANCEL_ORDER_STATUS_NO_CORRECT
	case tachibana.CancelOrderStatusInCorrect:
		return pb.CancelOrderStatus_CANCEL_ORDER_STATUS_INCORRECT
	case tachibana.CancelOrderStatusInCancel:
		return pb.CancelOrderStatus_CANCEL_ORDER_STATUS_IN_CANCEL
	case tachibana.CancelOrderStatusCorrected:
		return pb.CancelOrderStatus_CANCEL_ORDER_STATUS_CORRECTED
	case tachibana.CancelOrderStatusCanceled:
		return pb.CancelOrderStatus_CANCEL_ORDER_STATUS_CANCELED
	case tachibana.CancelOrderStatusCorrectFailed:
		return pb.CancelOrderStatus_CANCEL_ORDER_STATUS_CORRECT_FAILED
	case tachibana.CancelOrderStatusCancelFailed:
		return pb.CancelOrderStatus_CANCEL_ORDER_STATUS_CANCEL_FAILED
	case tachibana.CancelOrderStatusSwitch:
		return pb.CancelOrderStatus_CANCEL_ORDER_STATUS_SWITCH
	case tachibana.CancelOrderStatusSwitched:
		return pb.CancelOrderStatus_CANCEL_ORDER_STATUS_SWITCHED
	case tachibana.CancelOrderStatusSwitchFailed:
		return pb.CancelOrderStatus_CANCEL_ORDER_STATUS_SWITCH_FAILED
	}
	return pb.CancelOrderStatus_CANCEL_ORDER_STATUS_UNSPECIFIED
}

func (t *tachibanaApi) fromApprovalLogin(approvalLogin tachibana.ApprovalLogin) pb.ApprovalLogin {
	switch approvalLogin {
	case tachibana.ApprovalLoginUnApproval:
		return pb.ApprovalLogin_APPROVAL_LOGIN_UN_APPROVAL
	case tachibana.ApprovalLoginApproval:
		return pb.ApprovalLogin_APPROVAL_LOGIN_APPROVAL
	case tachibana.ApprovalLoginOutOfService:
		return pb.ApprovalLogin_APPROVAL_LOGIN_OUT_OF_SERVICE
	case tachibana.ApprovalLoginTesting:
		return pb.ApprovalLogin_APPROVAL_LOGIN_TESTING
	}
	return pb.ApprovalLogin_APPROVAL_LOGIN_UNSPECIFIED
}

func (t *tachibanaApi) fromSystemStatus(systemStatus tachibana.SystemStatus) pb.SystemStatus {
	switch systemStatus {
	case tachibana.SystemStatusClosing:
		return pb.SystemStatus_SYSTEM_STATUS_CLOSING
	case tachibana.SystemStatusOpening:
		return pb.SystemStatus_SYSTEM_STATUS_OPENING
	case tachibana.SystemStatusPause:
		return pb.SystemStatus_SYSTEM_STATUS_PAUSE
	}
	return pb.SystemStatus_SYSTEM_STATUS_UNSPECIFIED
}
