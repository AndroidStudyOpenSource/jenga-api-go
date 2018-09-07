package jenga

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
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

//Generate the Jenga
// Access Token
func (s Service) Auth() (string, error) {

	data := url.Values{}
	data.Set("username", s.Username)
	data.Add("password", s.Password)

	url := s.baseURL() + "identity-test/v2/token"
	req, err := http.NewRequest(http.MethodPost, url,  strings.NewReader(data.Encode()))
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

	str := spew.Sdump(authResponse)
	log.Printf(str)

	accessToken := authResponse.AccessToken
	log.Println("Username is ", s.Username)
	log.Println("Password is ", s.Password)
	log.Println("Token is ", accessToken)
	return accessToken, nil
}

// New return a new Jenga Service
func New(username, password string, env int) (Service, error) {
	return Service{username, password, env}, nil
}

func (s Service) baseURL() string {
	if s.Env == PRODUCTION {
		return "https://sandbox.jengahq.io/"
	}
	return "https://sandbox.jengahq.io/"
}
