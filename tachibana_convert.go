package tachibana_grpc_server

import (
	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
