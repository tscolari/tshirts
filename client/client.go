package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/tscolari/tshirts/challenge"
)

func New(baseUrl, accessToken string) *Client {
	return &Client{
		BaseUrl:     baseUrl,
		AccessToken: accessToken,
	}
}

type Client struct {
	BaseUrl     string
	AccessToken string
}

func (c Client) FetchInks() (challenge.Inks, error) {
	req := c.newRequest("GET", "inks", nil)
	resp, err := c.httpClient().Do(req)
	if err != nil {
		return challenge.Inks{}, err
	}

	var inksResponse challenge.InksResponse
	err = c.parseResponse(resp, &inksResponse)
	if err != nil {
		return challenge.Inks{}, err
	}

	return inksResponse.Inks, err
}

func (c Client) FetchQuestion() (challenge.Scenario, error) {
	req := c.newRequest("GET", "question/practice", nil)
	resp, err := c.httpClient().Do(req)
	if err != nil {
		return challenge.Scenario{}, err
	}

	var respScenario challenge.Scenario
	err = c.parseResponse(resp, &respScenario)

	return respScenario, err
}

func (c Client) PostAnswer(solution challenge.Solution) (string, error) {
	jsonBody, err := json.Marshal(solution)
	if err != nil {
		return "", err
	}

	req := c.newRequest("POST", "answer/practice", bytes.NewBuffer(jsonBody))
	resp, err := c.httpClient().Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func (c Client) url(url string) string {
	return fmt.Sprintf("%s/%s", c.BaseUrl, url)
}

func (c Client) httpClient() *http.Client {
	return &http.Client{}
}

func (c Client) newRequest(verb, path string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(verb, c.url(path), body)
	req.Header.Add("Auth-Token", c.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	return req
}

func (c Client) parseResponse(resp *http.Response, object interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, object)
	return err
}
