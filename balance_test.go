package privatbank

import (
	"encoding/xml"
	"testing"
)

func TestPrivat24Api_prepareAccountStatementRequest(t *testing.T) {
	type fields struct {
		merchantID       int
		merchantPassword string
	}
	type args struct {
		cardNumber string
		beginDate  string
		endDate    string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantBalanceRequest string
	}{

		{name: "realexample",
			fields:             fields{131270, "0UU014SHF000000091yym3W2mNL4qA0Q"},
			args:               args{"5363000000007422", "12.07.2019", "01.08.2019"},
			wantBalanceRequest: "<request version=\"1.0\"><merchant><id>131270</id><signature>192df5c087cb0b3fe40b2461dd59ee0b0f91131f</signature></merchant><data><oper>cmt</oper><wait>0</wait><payment id=\"\"><prop name=\"sd\" value=\"12.07.2019\"></prop><prop name=\"ed\" value=\"01.08.2019\"></prop><prop name=\"card\" value=\"5363000000007422\"></prop></payment></data></request>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &Privat24Api{
				MerchantID:       tt.fields.merchantID,
				MerchantPassword: tt.fields.merchantPassword,
			}
			gotBalanceRequest := api.prepareAccountStatementRequest(tt.args.cardNumber, tt.args.beginDate, tt.args.endDate)
			gotBalanceRequestXML, _ := xml.Marshal(gotBalanceRequest)

			if string(gotBalanceRequestXML) != tt.wantBalanceRequest {
				//fmt.Println(string(gotBalanceRequestXML))
				t.Errorf("Privat24Api.prepareAccountStatementRequest() = %v, \n\nwant %v", string(gotBalanceRequestXML), tt.wantBalanceRequest)
			}
		})
	}
}
