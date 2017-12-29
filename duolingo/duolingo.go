package duolingo

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type (
	Client struct {
		Id         string `json:"user_id"`
		Username   string `json:"username"`
		httpClient http.Client
	}
)

func Login(login string, password string) Client {
	var client Client
	_url := "https://www.duolingo.com/login"

	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	httpClient := http.Client{Jar: jar}
	resp, _ := httpClient.PostForm(_url, url.Values{
		"login":    {login},
		"password": {password},
	})
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&client)
	client.httpClient = httpClient

	return client
}

func (c Client) GetActivity() {
	_url := fmt.Sprintf("%s", c.Id)

	resp, _ := c.httpClient.Get(_url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
