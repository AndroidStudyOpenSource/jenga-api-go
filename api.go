package jenga

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Env is the environment type
type Env string

const (
	// DEV is the development env tag

	// SANDBOX is the sandbox env tag
	SANDBOX = iota
	// PRODUCTION is the production env tag
	PRODUCTION
)

// Service is an Jenga Service
type Service struct {
	Username string
	Password string
	Env      int
}

var signature = "vtU9bsRz0WSBrjLYY4dbqloUd0bk7mE6rYa80jxJls7R++YA5hStZIh8mZkMFwQ4UEfXIwQQES8DpP0H9lhyt62ftLf3i6M4WcI31KV4VK2w2Wqf7ZVouw1pYbitWuMcoEQc0YUHBUPMFVmuO8N82ns72914Oms3iOlxg9/pkC1W/FWCHQAOq8RWNGFpmsufEtEnKUOUKAsj0+yVrJ1fpUEpqG2I5hVipz0/c0RVAhuHnTH+/YY6n7jCraSUMMGSfgUDPwY7WgaVfMVv30UTKsq6a0JEdsvOeUVr4jDao+WLK4W6cv3S2vJSDex5lmnQykFptWeVZn0u0PsPu1aTfw=="

// New return a new Jenga Service
func New(username, password string, env int) (Service, error) {
	return Service{username, password, env}, nil
}

//Generate the Jenga Access Token
func (s Service) auth() (string, error) {

	data := url.Values{}
	data.Set("username", s.Username)
	data.Add("password", s.Password)

	reqUrl := s.baseURL() + "identity-test/v2/token"
	req, err := http.NewRequest(http.MethodPost, reqUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Basic QWEzY0dxRVpWOVo5dDFVSnhTR3BwZUF4WEZrcFFyYUk6cU5BYWs5YXcyVDNtSW1uMg==")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return "", fmt.Errorf("could not send auth request: %v", err)
	}

	var authResponse authResponse
	err = json.NewDecoder(res.Body).Decode(&authResponse)
	if err != nil {
		return "", fmt.Errorf("could not decode auth response: %v", err)
	}

	accessToken := authResponse.AccessToken
	return accessToken, nil
}

// BalanceInquiry sends a balance inquiry
func (s Service) BalanceInquiry() string {

	auth, err := s.auth()
	if err != nil {
		return ""
	}

	reqUrl := s.baseURL() + "account-test/v2/accounts/balances/KE/0011547896523"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	req.Header.Set("Authorization", "Bearer "+auth)
	req.Header.Set("signature", signature)

	res, err := client.Do(req)
	if err != nil {
		defer res.Body.Close()
	}

	stringBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}

	return string(stringBody)

}

// MobileWalletRequest sends a new request
func (s Service) MobileWalletRequest(mobileWallets MobileWallets) (string, error) {
	body, err := json.Marshal(mobileWallets)
	if err != nil {
		return "", nil
	}
	auth, err := s.auth()
	if err != nil {
		return "", nil
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["signature"] = signature

	reqUrl := s.baseURL() + "transaction-test/v2/remittance"
	return s.newReq(reqUrl, body, headers)
}

func (s Service) newReq(url string, body []byte, headers map[string]string) (string, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", nil
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(request)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return "", err
	}

	stringBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(stringBody), nil
}

func (s Service) baseURL() string {
	if s.Env == PRODUCTION {
		return "https://sandbox.jengahq.io/"
	}
	return "https://sandbox.jengahq.io/"
}
