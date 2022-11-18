package apiclients

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	Text string
	Authore string
}

type QuoteList []Quote

type QuotesAPIClient struct {
	HttpClient *http.Client
}

func NewQuotesAPIClient() *QuotesAPIClient {
	return &QuotesAPIClient{
		HttpClient: &http.Client{},
	}
}

func (c *QuotesAPIClient) GetQuotes() (QuoteList, error) {
	resp, err := http.Get("https://type.fit/api/quotes")
	if err != nil {
		return QuoteList{}, err
	}
	defer resp.Body.Close()
	rawBody, err := ioutil.ReadAll(resp.Body)
	var quotes QuoteList
	if err := json.Unmarshal(rawBody, &quotes); err != nil {
		return QuoteList{}, err
	}

	return quotes, nil
}
