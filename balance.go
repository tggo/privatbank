package privatbank

import (
	"bytes"
	"encoding/xml"
	"github.com/sirupsen/logrus"
	"time"
)

type BalanceRequestXML struct {
	XMLName  xml.Name `xml:"request"`
	Version  string   `xml:"version,attr"`
	Merchant Merchant `xml:"merchant"`
	Data     struct {
		XMLName   xml.Name `xml:"data"`
		Operation string   `xml:"oper"`
		Wait      int      `xml:"wait"`
		//Test      int      `xml:"test"`
		Payment Payment
	}
}

type BalanceResponseXML struct {
	XMLName  xml.Name `xml:"response"`
	Version  float64  `xml:"version,attr"`
	Merchant Merchant `xml:"merchant"`
	Data     struct {
		XMLName   xml.Name `xml:"data"`
		Operation string   `xml:"oper"`
		Info      struct {
			XMLName     xml.Name `xml:"info"`
			CardBalance CardBalance
		}
	}
}

type CardBalance struct {
	XMLName    xml.Name `xml:"cardbalance"`
	Card       Card
	AvBalance  float64   `xml:"av_balance"`
	BalDate    time.Time `xml:"bal_date"`
	BalDyn     string    `xml:"bal_dyn"`
	Balance    float64   `xml:"balance"`
	FinLimit   float64   `xml:"fin_limit"`
	TradeLimit float64   `xml:"trade_limit"`
}

type Card struct {
	XMLName        xml.Name `xml:"card"`
	Account        int      `xml:"account"`
	CardNumber     int      `xml:"card_number"`
	AccName        string   `xml:"acc_name"`
	AccType        string   `xml:"acc_type"`
	Currency       string   `xml:"currency"`
	CardType       string   `xml:"card_type"`
	MainCardNumber int      `xml:"main_card_number"`
	CardStat       string   `xml:"card_stat"`
	Src            string   `xml:"src"`
}

func (api *Privat24Api) Balance(cardNumber string) BalanceResponseXML {
	url := api.APIUrl + "/balance"

	payment := new(Payment)
	payment.SetBalanceProperties(cardNumber, "UA")

	balanceRequest := new(BalanceRequestXML)
	balanceRequest.Version = "1.0"
	balanceRequest.Data.Operation = "cmt"
	balanceRequest.Data.Wait = 0
	//balanceRequest.Data.Test = 0
	balanceRequest.Data.Payment = *payment
	balanceRequest.Merchant = api.getMerchantStructure(balanceRequest.Data)

	byteXML, err := xml.Marshal(balanceRequest)
	if err != nil {
		logrus.Println(err.Error())
	}

	reqBody := append([]byte{}, []byte(xml.Header)...)
	reqBody = append(reqBody, byteXML...)

	response, err := api.requestXML(url, bytes.NewBuffer(byteXML), "POST")
	if err != nil {
		logrus.Println(err.Error())
	}

	balanceResponse := new(BalanceResponseXML)
	logrus.Infof("%+v", response)

	err = xml.Unmarshal(response, &balanceResponse)
	if err != nil {
		logrus.Println(err.Error())
	}

	return *balanceResponse
}

func (api *Privat24Api) AccountStatement(cardNumber string, startDate string, endDate string) OrdersResponseXML {
	url := api.APIUrl + "/rest_fiz"

	balanceRequest := api.prepareAccountStatementRequest(cardNumber, startDate, endDate)
	byteXML, err := xml.Marshal(balanceRequest)
	if err != nil {
		logrus.Println(err.Error())
	}

	reqBody := append([]byte{}, []byte(xml.Header)...)
	reqBody = append(reqBody, byteXML...)

	response, err := api.requestXML(url, bytes.NewBuffer(byteXML), "POST")
	if err != nil {
		logrus.Errorf("error make request %s", err.Error())
	}

	balanceResponse := new(OrdersResponseXML)

	err = xml.Unmarshal(response, &balanceResponse)
	if err != nil {
		logrus.Errorf("error Unmarshal %s", err.Error())
		logrus.Errorf("body: %s", string(response))

	}

	return *balanceResponse
}

func (api *Privat24Api) prepareAccountStatementRequest(cardNumber string, startDate string, endDate string) (balanceRequest BalanceRequestXML) {
	payment := new(Payment)
	payment.FizAccountStatmentProperties(cardNumber, startDate, endDate)

	balanceRequest.Version = "1.0"
	balanceRequest.Data.Operation = "cmt"
	balanceRequest.Data.Wait = 0
	//balanceRequest.Data.Test = 0
	balanceRequest.Data.Payment = *payment
	balanceRequest.Merchant = api.getMerchantStructure(balanceRequest.Data)

	return balanceRequest
}
