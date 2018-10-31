package main

import (
	"encoding/json"
	"github.com/AndroidStudyOpenSource/jenga-api-go"
	"log"
)

const (
	username = "8485200649"                        // sandbox --> change to yours
	password = "Hb8jahNZDnPjCw0T9RXxWH8KGvwQJweZK" // sandbox --> change to yours
)

func main() {

	jeng, err := jenga.New(username, password, jenga.SANDBOX)
	if err != nil {
		panic(err)
	}

	bal := jeng.BalanceInquiry()
	log.Println(bal)

	//config := jenga.MobileWallets{
	//	Source: struct {
	//		CountryCode   string
	//		Name          string
	//		AccountNumber string
	//	}{CountryCode: "KES", Name: "Tom Doe", AccountNumber: "0011547896523"},
	//	Destination: struct {
	//		Type         string
	//		CountryCode  string
	//		Name         string
	//		MobileNumber string
	//		WalletName   string
	//	}{Type: "mobile", CountryCode: "KE", Name: "John Doe", MobileNumber: "0763555619", WalletName: "Equitel"},
	//	Transfer: struct {
	//		Type         string
	//		Amount       string
	//		CurrencyCode string
	//		Reference    string
	//		Date         string
	//		Description  string
	//	}{Type: "MobileWallet", Amount: "20", CurrencyCode: "KES", Reference: "692194625798", Date: "2018-06-14", Description: "Some remarks here"},
	//}
	//
	//mobile, err := jeng.MobileWalletRequest(config)
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(mobile)



	jsonConfig := []byte(`{
   "source": {
      "countryCode": "KE",
      "name": "Tom Doe",
      "accountNumber": "0011547896523"
   },
   "destination": {
      "type": "mobile",
      "countryCode": "KE",
      "name": "John Doe",
      "mobileNumber": "0763555619",
      "walletName": "Equitel"
   },
   "transfer": {
      "type": "MobileWallet",
      "amount": "20",
      "currencyCode": "KES",
      "reference": "692194625798",
      "date": "2018-06-14",
      "description": "Some remarks here"
   }}`)
	var config jenga.MobileWallets
	err = json.Unmarshal(jsonConfig, &config)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Config: %+v\n", config)
	mobile, err := jeng.MobileWalletRequest(config)

	if err != nil {
		log.Println(err)
	}
	log.Println(mobile)


}
