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
