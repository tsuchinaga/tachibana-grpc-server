package tachibana_grpc_server

import (
	pb "gitlab.com/tsuchinaga/tachibana-grpc-server/tachibanapb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"testing"
)

func Test_accountSession_getLoginResponse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		session *accountSession
		arg1    string
		want1   *pb.LoginResponse
	}{
		{name: "引数のtokenと、baseResponseを使ってgRPCで返すためのLoginResponseが作れる",
			session: &accountSession{BaseResponse: &pb.LoginResponse{
				CommonResponse: &pb.CommonResponse{
					No:           1,
					SendDate:     &timestamppb.Timestamp{Seconds: 1656366658, Nanos: 470000000},
					ReceiveDate:  &timestamppb.Timestamp{Seconds: 1656366658, Nanos: 413000000},
					ErrorNo:      pb.ErrorNo_ERROR_NO_NO_PROBLEM,
					ErrorMessage: "",
					MessageType:  pb.MessageType_MESSAGE_TYPE_LOGIN_RESPONSE,
				},
				ResultCode:                "0",
				ResultText:                "",
				AccountType:               pb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				SecondPasswordOmit:        false,
				LastLoginDatetime:         &timestamppb.Timestamp{Seconds: 1656366551, Nanos: 0},
				GeneralAccount:            true,
				SafekeepingAccount:        true,
				TransferAccount:           true,
				ForeignAccount:            true,
				MrfAccount:                false,
				StockSpecificAccount:      pb.SpecificAccountType_SPECIFIC_ACCOUNT_TYPE_NOTHING,
				MarginSpecificAccount:     pb.SpecificAccountType_SPECIFIC_ACCOUNT_TYPE_NOTHING,
				InvestmentSpecificAccount: pb.SpecificAccountType_SPECIFIC_ACCOUNT_TYPE_NOTHING,
				DividendAccount:           false,
				SpecificAccount:           true,
				MarginAccount:             true,
				FutureOptionAccount:       false,
				MmfAccount:                false,
				ChinaForeignAccount:       false,
				FxAccount:                 false,
				NisaAccount:               false,
				UnreadDocument:            false,
				Token:                     "",
			}},
			arg1: "c63454cf231f3a8cd967035d578cf166d50a646842de86090cf76fcda1b52e30",
			want1: &pb.LoginResponse{
				CommonResponse: &pb.CommonResponse{
					No:           1,
					SendDate:     &timestamppb.Timestamp{Seconds: 1656366658, Nanos: 470000000},
					ReceiveDate:  &timestamppb.Timestamp{Seconds: 1656366658, Nanos: 413000000},
					ErrorNo:      pb.ErrorNo_ERROR_NO_NO_PROBLEM,
					ErrorMessage: "",
					MessageType:  pb.MessageType_MESSAGE_TYPE_LOGIN_RESPONSE,
				},
				ResultCode:                "0",
				ResultText:                "",
				AccountType:               pb.AccountType_ACCOUNT_TYPE_SPECIFIC,
				SecondPasswordOmit:        false,
				LastLoginDatetime:         &timestamppb.Timestamp{Seconds: 1656366551, Nanos: 0},
				GeneralAccount:            true,
				SafekeepingAccount:        true,
				TransferAccount:           true,
				ForeignAccount:            true,
				MrfAccount:                false,
				StockSpecificAccount:      pb.SpecificAccountType_SPECIFIC_ACCOUNT_TYPE_NOTHING,
				MarginSpecificAccount:     pb.SpecificAccountType_SPECIFIC_ACCOUNT_TYPE_NOTHING,
				InvestmentSpecificAccount: pb.SpecificAccountType_SPECIFIC_ACCOUNT_TYPE_NOTHING,
				DividendAccount:           false,
				SpecificAccount:           true,
				MarginAccount:             true,
				FutureOptionAccount:       false,
				MmfAccount:                false,
				ChinaForeignAccount:       false,
				FxAccount:                 false,
				NisaAccount:               false,
				UnreadDocument:            false,
				Token:                     "c63454cf231f3a8cd967035d578cf166d50a646842de86090cf76fcda1b52e30",
			}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got1 := test.session.getLoginResponse(test.arg1)
			if !reflect.DeepEqual(test.want1, got1) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want1, got1)
			}
		})
	}
}
