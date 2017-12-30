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

	Courses struct {
		Course []struct {
			FromLanguage     string `json:"fromLanguage"`
			Title            string `json:"title"`
			LearningLanguage string `json:"learningLanguage"`
			HealthEnabled    bool   `json:"healthEnabled"`
			AuthorID         string `json:"authorId"`
			Xp               int    `json:"xp"`
			ID               string `json:"id"`
		} `json:"courses"`
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

func (c Client) GetLeaderboard() {
	_url := "https://www.duolingo.com/friendships/leaderboard_activity"

	resp, _ := c.httpClient.Get(_url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func (c Client) GetCourses() Courses {
	var courses Courses
	_url := fmt.Sprintf("https://www.duolingo.com/2017-06-30/users/%s?fields=courses", c.Id)

	resp, _ := c.httpClient.Get(_url)
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&courses)

	fmt.Println(courses.Course[0])

	return courses
}

func (c Client) GetFromLanguage() {
	_url := fmt.Sprintf("https://www.duolingo.com/2017-06-30/users/%s?fields=fromLanguage", c.Id)

	resp, _ := c.httpClient.Get(_url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func (c Client) GetLearningLanguage() {
	_url := fmt.Sprintf("https://www.duolingo.com/2017-06-30/users/%s?fields=learningLanguage", c.Id)

	resp, _ := c.httpClient.Get(_url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
