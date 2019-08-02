package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/tggo/privatbank"
)

func main() {
	merchantPtr := flag.Int("merchant", 0, "merchant ID")
	merchantPasswordPtr := flag.String("password", "", "merchant password")

	cardNumPtr := flag.String("card", "", "card number")
	startDatePtr := flag.String("start", "01.07.2019", "start date")
	endDatePtr := flag.String("end", "31.12.2019", "end date")
	flag.Parse()

	if *cardNumPtr == "" {
		panic("need set -card")
	}

	if *merchantPasswordPtr == "" {
		panic("need set -password")
	}

	if *merchantPtr == 0 {
		panic("need set -merchant")
	}

	api := privatbank.NewApi(*merchantPtr, *merchantPasswordPtr)
	gotBalanceRequest := api.AccountStatement(*cardNumPtr, *startDatePtr, *endDatePtr)
	logrus.Infof("%+v", gotBalanceRequest)
	//data := api.GetExchangeArchive(time.Date(2018, 1, 18, 0,0,0,0, time.Local))
	//logrus.Infof("%+v", data)
}
