package main

import (
	"fmt"
	"testing"
)

var testDummyResponse = SocialAccountsRequests{
	Accounts: []SocialAccountRequest{
		{
			AccountName: "goeyespy",
			URL:         "https://twitter.com/goeyejkjkjspy",
			Service:     "twitter",
		},
	},
}

func TestSocialHandler(t *testing.T) {

}

func TestTryTwitter(t *testing.T) {
	accountName := "goeyespy"
	result := buildTwitterRequest(accountName)
	if result.URL != "https://twitter.com/goeyespy" {
		t.Error("Expected https://twitter.com/goeyespy but got", result.URL)
	}
}

func TestTryAccount(t *testing.T) {
	account := testDummyResponse.Accounts[0]
	response := tryAccount(account)
	fmt.Printf("result = %v and statusCode = %v \n", response.Available, response.StatusCode)
}
