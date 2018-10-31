package jenga

type authResponse struct {
	AccessToken string `json:"access_token"`
}

type Balance struct {
	Currency string `json:"currency"`
	Balances []struct {
		Amount string `json:"amount"`
		Type   string `json:"type"`
	} `json:"balances"`
}

type MobileWallets struct {
	Source struct {
		CountryCode   string `json:"countryCode"`
		Name          string `json:"name"`
		AccountNumber string `json:"accountNumber"`
	} `json:"source"`
	Destination struct {
		Type         string `json:"type"`
		CountryCode  string `json:"countryCode"`
		Name         string `json:"name"`
		MobileNumber string `json:"mobileNumber"`
		WalletName   string `json:"walletName"`
	} `json:"destination"`
	Transfer struct {
		Type         string `json:"type"`
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currencyCode"`
		Reference    string `json:"reference"`
		Date         string `json:"date"`
		Description  string `json:"description"`
	} `json:"transfer"`
}

type RTGS struct {
	Source struct {
		CountryCode   string `json:"countryCode"`
		Name          string `json:"name"`
		AccountNumber string `json:"accountNumber"`
	} `json:"source"`
	Destination struct {
		Type          string `json:"type"`
		CountryCode   string `json:"countryCode"`
		Name          string `json:"name"`
		BankCode      string `json:"bankCode"`
		AccountNumber string `json:"accountNumber"`
	} `json:"destination"`
	Transfer struct {
		Type         string `json:"type"`
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currencyCode"`
		Reference    string `json:"reference"`
		Date         string `json:"date"`
		Description  string `json:"description"`
	} `json:"transfer"`
}