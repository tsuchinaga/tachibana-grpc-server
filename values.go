package tachibana_grpc_server

import (
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
	"time"

	tachibana "gitlab.com/tsuchinaga/go-tachibanaapi"
)

type accountSession struct {
	Date         time.Time
	Token        string
	Session      *tachibana.Session
	BaseResponse *pb.LoginResponse
}

func (v *accountSession) getLoginResponse(token string) *pb.LoginResponse {
	return &pb.LoginResponse{
		CommonResponse:            v.BaseResponse.CommonResponse,
		ResultCode:                v.BaseResponse.ResultCode,
		ResultText:                v.BaseResponse.ResultText,
		AccountType:               v.BaseResponse.AccountType,
		SecondPasswordOmit:        v.BaseResponse.SecondPasswordOmit,
		LastLoginDatetime:         v.BaseResponse.LastLoginDatetime,
		GeneralAccount:            v.BaseResponse.GeneralAccount,
		SafekeepingAccount:        v.BaseResponse.SafekeepingAccount,
		TransferAccount:           v.BaseResponse.TransferAccount,
		ForeignAccount:            v.BaseResponse.ForeignAccount,
		MrfAccount:                v.BaseResponse.MrfAccount,
		StockSpecificAccount:      v.BaseResponse.StockSpecificAccount,
		MarginSpecificAccount:     v.BaseResponse.MarginSpecificAccount,
		InvestmentSpecificAccount: v.BaseResponse.InvestmentSpecificAccount,
		DividendAccount:           v.BaseResponse.DividendAccount,
		SpecificAccount:           v.BaseResponse.SpecificAccount,
		MarginAccount:             v.BaseResponse.MarginAccount,
		FutureOptionAccount:       v.BaseResponse.FutureOptionAccount,
		MmfAccount:                v.BaseResponse.MmfAccount,
		ChinaForeignAccount:       v.BaseResponse.ChinaForeignAccount,
		FxAccount:                 v.BaseResponse.FxAccount,
		NisaAccount:               v.BaseResponse.NisaAccount,
		UnreadDocument:            v.BaseResponse.UnreadDocument,
		Token:                     token,
	}
}
