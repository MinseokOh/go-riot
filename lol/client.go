package lol

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type Client struct {
	apiKey string
	region string

	client *http.Client

	Summoner        *SummonerAPI
	Match           *MatchAPI
	Champion        *ChampionAPI
	ChampionMastery *ChampionMasteryAPI
	Clash           *ClashAPI
}

func NewClient(region string, apiKey string, httpClient *http.Client) (client *Client) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client = &Client{
		region: region,
		apiKey: apiKey,
		client: httpClient,
	}

	client.Summoner = &SummonerAPI{client: client}
	client.Champion = &ChampionAPI{client: client}
	client.ChampionMastery = &ChampionMasteryAPI{client: client}
	client.Match = &MatchAPI{client: client}
	client.Clash = &ClashAPI{client: client}
	return
}

func (c *Client) GetPath(api interface{}, params ...string) (path string) {
	var sb strings.Builder

	pkg := strings.Split(
		reflect.ValueOf(api).Elem().Type().PkgPath(),
		"/")[1]

	sb.WriteString("https://")
	sb.WriteString(c.region + "/")
	sb.WriteString(pkg + "/")
	sb.WriteString(api.(API).APIName() + "/")
	sb.WriteString(api.(API).APIVersion() + "/")

	for _, param := range params {
		sb.WriteString(param + "/")
	}

	path = sb.String()
	fmt.Println("Path : ", sb.String())

	return
}

func (c *Client) NewRequest(method string, url string) (req *http.Request, err error) {
	if len(c.apiKey) == 0 {
		return
	}

	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	req.Header.Add("X-Riot-Token", c.apiKey)

	return
}

func (c *Client) Do(req *http.Request, v interface{}) (err error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if v != nil {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(body, v)
	}

	return
}
