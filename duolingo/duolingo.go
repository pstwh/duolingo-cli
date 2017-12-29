package duolingo

import (
	"fmt"

	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func Login() {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	client := http.Client{Jar: jar}
	resp, _ := client.PostForm("https://www.duolingo.com/login", url.Values{
		"login":    {""},
		"password": {""},
	})

	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Println(string(data))
}
