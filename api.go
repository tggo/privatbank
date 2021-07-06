package privatbank

import (
	"encoding/xml"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	"fmt"
)

const apiUrl = "https://api.privatbank.ua/p24api"

type Privat24Api struct {
	MerchantID       int
	MerchantPassword string
	Client           *http.Client
	APIUrl           string
}

func NewPublicApi() *Privat24Api {
	client := &http.Client{
		Timeout: time.Second * 360,
	}
	api := &Privat24Api{Client: client, APIUrl: apiUrl}

	return api
}

func NewApi(merchantID int, merchantPassword string) *Privat24Api {
	client := &http.Client{
		Timeout: time.Second * 360,
	}
	api := &Privat24Api{MerchantID: merchantID, MerchantPassword: merchantPassword, Client: client, APIUrl: apiUrl}

	return api
}

func (api *Privat24Api) requestXML(url string, body io.Reader, method string) ([]byte, error) {
	req, err := http.NewRequest(http.Request{}.Method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	response, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		logrus.Println("[REQUEST ERROR]: ", response.Status, string(result))
		err = errors.New(response.Status + " : " + string(result))
	}

	return result, err
}

func (api *Privat24Api) getMerchantStructure(data interface{}) Merchant {
	res, err := xml.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	merchant := new(Merchant)
	merchant.ID = api.MerchantID
	merchant.Signature = SHA1(GetMD5Hash(string(res) + api.MerchantPassword))
	fmt.Printf("merchant: %+v\n",merchant)
	
	return *merchant
}
