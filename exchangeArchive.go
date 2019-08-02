package privatbank

import (
	"encoding/xml"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type ExchangeArchiveResponse struct {
	XMLName         xml.Name              `xml:"exchangerates"`
	Date            string                `xml:"date,attr"`
	Bank            string                `xml:"bank,attr"`
	BaseCurrency    string                `xml:"BaseCurrency,attr"`
	BaseCurrencyLit string                `xml:"BaseCurrencyLit,attr"`
	ExchangeRate    []ExchangeRateArchive `xml:"exchangerate"`
}

type ExchangeRateArchive struct {
	XMLName        xml.Name `xml:"exchangerate"`
	BaseCurrency   string   `xml:"baseCurrency,attr"`   //Базовая валюта
	Currency       string   `xml:"currency,attr"`       //Валюта сделки
	SaleRateNB     string   `xml:"saleRateNB,attr"`     //Курс продажи НБУ
	PurchaseRateNB string   `xml:"purchaseRateNB,attr"` //Курс покупки НБУ
	SaleRate       string   `xml:"saleRate,attr"`       //Курс продажи ПриватБанка
	PurchaseRate   string   `xml:"purchaseRate,attr"`   //Курс покупки ПриватБанка
}

//GetExchangeArchive - Архив курсов валют ПриватБанка, НБУ
//API позволяет получить информацию о наличных курсах валют ПриватБанка и НБУ на выбранную дату.
//Архив хранит данные за последние 4 года
//01.12.2014
func (api *Privat24Api) GetExchangeArchive(date time.Time) ExchangeArchiveResponse {
	url := "https://api.privatbank.ua/p24api/exchange_rates?date=" + date.Format("02.01.2006")
	logrus.Println(date.Format("02.01.2006"))
	response, err := api.requestXML(url, nil, "GET")
	if err != nil {
		log.Println(err.Error())
	}

	var exchangeArchiveResponse ExchangeArchiveResponse
	err = xml.Unmarshal(response, &exchangeArchiveResponse)
	if err != nil {
		logrus.Println(err.Error())
	}

	return exchangeArchiveResponse
}
