package main

import (
	"encoding/json"
	"fmt"
	//	"log"
	"net/http"
)

type SocialAccountsRequests struct {
	Accounts []SocialAccountRequest
}

type SocialAccountRequest struct {
	AccountName string
	URL         string
	Service     string
}

type SocialAccountResponses struct {
	Responses []SocialAccountResponse
}

func (s *SocialAccountResponses) add(result SocialAccountResponse) {
	s.Responses = append(s.Responses, result)
}

type SocialAccountResponse struct {
	Available   bool
	AccountName string
	URL         string
	Service     string
	StatusCode  int
}

var dummyResponse = SocialAccountsRequests{
	Accounts: []SocialAccountRequest{
		{
			AccountName: "blah",
			URL:         "blah",
			Service:     "twitter",
		},
	},
}

func buildTwitterRequest(accountNameQuery string) SocialAccountRequest {
	//TODO case sensitive?
	return SocialAccountRequest{
		AccountName: accountNameQuery,
		URL:         fmt.Sprintf("https://twitter.com/%v", accountNameQuery),
		Service:     "twitter",
	}
}

func buildInstagramRequest(accountNameQuery string) SocialAccountRequest {
	//TODO case sensitive?
	return SocialAccountRequest{
		AccountName: accountNameQuery,
		URL:         fmt.Sprintf("https://www.instagram.com/%v", accountNameQuery),
		Service:     "instagram",
	}
}

func buildRequests(accountNameQuery string) SocialAccountsRequests {
	var requests SocialAccountsRequests
	requests.Accounts = append(requests.Accounts, buildTwitterRequest(accountNameQuery))
	requests.Accounts = append(requests.Accounts, buildInstagramRequest(accountNameQuery))
	return requests
}

// upload logic
func checkSocialHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	name := "goeyespylkjlkjlkjlkj"

	accountRequests := buildRequests(name)
	var accountResponses SocialAccountResponses
	for _, accountRequest := range accountRequests.Accounts {
		result := tryAccount(accountRequest)
		accountResponses.add(result)
	}

	replyJSON(w, accountResponses)
}

func replyJSON(w http.ResponseWriter, accounts SocialAccountResponses) {
	js, err := json.Marshal(accounts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func tryAccount(account SocialAccountRequest) SocialAccountResponse {
	var result SocialAccountResponse
	response, err := http.Get(account.URL) //HACK this line can panic if not net connection is available?
	defer response.Body.Close()

	result.StatusCode = response.StatusCode
	result.AccountName = account.AccountName
	result.Available = false
	result.Service = account.Service
	result.URL = account.URL

	if (err == nil) && (result.StatusCode == 200) {
		result.Available = false
	} else if err == nil {
		result.Available = true
	}
	return result
}
