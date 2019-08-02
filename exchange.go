package privatbank

import (
	"encoding/xml"
	"log"
)

type ExchangeRates struct {
	XMLName xml.Name `xml:"exchangerates"`
	Rows    []Row    `xml:"row"`
}

type Row struct {
	XMLName      xml.Name     `xml:"row"`
	ExchangeRate ExchangeRate `xml:"exchangerate"`
}

type ExchangeRate struct {
	XMLName xml.Name `xml:"exchangerate"`
	Ccy     string   `xml:"ccy,attr"`      //Код валюты (справочник кодов валют: https://ru.wikipedia.org/wiki/Коды_валют)
	BaseCcy string   `xml:"base_ccy,attr"` //Код национальной валюты
	Buy     string   `xml:"buy,attr"`      //Курс покупки
	Sale    string   `xml:"sale,attr"`     //Курс продажи
}

//Наличный курс ПриватБанка (в отделениях):
func (api *Privat24Api) GetExchangeRatesCash() ExchangeRates {
	url := api.APIUrl + "/p24api/pubinfo?exchange&coursid=5"

	response, err := api.requestXML(url, nil, "GET")
	if err != nil {
		log.Println(err.Error())
	}

	var exchangerates ExchangeRates

	err = xml.Unmarshal(response, &exchangerates)
	if err != nil {
		log.Println(err.Error())
	}

	return exchangerates
}

//Безналичный курс ПриватБанка (конвертация по картам, Приват24, пополнение вкладов):
func (api *Privat24Api) GetExchangeRatesCard() ExchangeRates {
	url := api.APIUrl + "/p24api/pubinfo?exchange&coursid=11"

	response, err := api.requestXML(url, nil, "GET")
	if err != nil {
		log.Println(err.Error())
	}

	var exchangerates ExchangeRates

	err = xml.Unmarshal(response, &exchangerates)
	if err != nil {
		log.Println(err.Error())
	}

	return exchangerates
}
