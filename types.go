package privatbank

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"strings"

	//"log"
)

type Merchant struct {
	ID        int    `xml:"id"`
	Signature string `xml:"signature"`
}

type Prop struct {
	XMLName xml.Name `xml:"prop"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Payment struct {
	XMLName    xml.Name `xml:"payment"`
	ID         string   `xml:"id,attr"`
	Properties []Prop
}

type Info struct {
	XMLName xml.Name `xml:"info"`
	Items   []interface{}
}

type OrdersRequestXML struct {
	XMLName  xml.Name `xml:"request"`
	Version  string   `xml:"version,attr"`
	Merchant Merchant `xml:"merchant"`
	Data     struct {
		XMLName xml.Name `xml:"data"`
		Oper    string   `xml:"oper"`
		Wait    int      `xml:"wait"`
		Test    int      `xml:"test"`
		Payment Payment
	}
}

type OrdersResponseXML struct {
	XMLName  xml.Name `xml:"response"`
	Version  string   `xml:"version,attr"`
	Merchant Merchant `xml:"merchant"`
	Data     struct {
		XMLName xml.Name `xml:"data"`
		Oper    string   `xml:"oper"`
		Info    struct {
			XMLName    xml.Name `xml:"info"`
			Statements struct {
				XMLName   xml.Name    `xml:"statements"`
				Status    string      `xml:"status,attr"`
				Credit    string      `xml:"credit,attr"`
				Debet     string      `xml:"debet,attr"`
				Statement []Statement `xml:"statement"`
			}
		}
	}
}

type Statement struct {
	XMLName     xml.Name `xml:"statement"`
	Card        string   `xml:"card,attr"`
	AppCode     string   `xml:"appcode,attr"`
	TranDate    string   `xml:"trandate,attr"`
	TranTime    string   `xml:"trantime,attr"`
	Amount      string   `xml:"amount,attr"`
	CardAmount  string   `xml:"cardamount,attr"`
	Rest        string   `xml:"rest,attr"`
	Terminal    string   `xml:"terminal,attr"`
	Description string   `xml:"description,attr"`
}

//Выписки по счёту мерчанта - физлица


// установить prop для просмотра баланса
// cardnum Номер карты
// country Страна
func (p *Payment) SetBalanceProperties(cardnum string, country string) {
	paymentProp := make([]Prop, 2)
	paymentProp[0] = Prop{Name: "cardnum", Value: cardnum}
	paymentProp[1] = Prop{Name: "country", Value: country}
	p.Properties = paymentProp
}

func (p *Payment) FizAccountStatmentProperties(cardnum string, startDate string, endDate string) {
	paymentProp := make([]Prop, 3)
	paymentProp[0] = Prop{Name: "sd", Value: startDate}
	paymentProp[1] = Prop{Name: "ed", Value: endDate}
	paymentProp[2] = Prop{Name: "card", Value: cardnum}
	p.Properties = paymentProp
}

// prop для платежа на карту приват банка
// bCardOrAcc Карта или счёт получателя
// amt Сумма Напр.: 23.05
// ccy Валюта (UAH, EUR, USD)
// details Назначение платежа
func (p *Payment) SetPrivatPaymentProperties(bCardOrAcc string, amt string, ccy string, details string) {
	paymentProp := make([]Prop, 4)
	paymentProp[0] = Prop{Name: "b_card_or_acc", Value: bCardOrAcc}
	paymentProp[1] = Prop{Name: "amt", Value: amt}
	paymentProp[2] = Prop{Name: "ccy", Value: ccy}
	paymentProp[3] = Prop{Name: "amt", Value: details}
	p.Properties = paymentProp
}

func GetMD5Hash(text string) string {
	text = strings.ReplaceAll(text,"<data>","")
	text = strings.ReplaceAll(text,"</data>","")
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}



func SHA1(text string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
