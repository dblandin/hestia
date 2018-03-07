package music

import (
	"encoding/json"
	"fmt"
	"github.com/codeclimate/hestia/internal/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CommandResponse struct {
	Error    string `json:"Error"`
	Response string `json:"response"`
}

type StateResponse struct {
	Album    string `json:"album"`
	AlbumArt string `json:"albumart"`
	Artist   string `json:"artist"`
	Mute     bool   `json:"mute"`
	Random   bool   `json:"random"`
	Repeat   bool   `json:"repeat"`
	Status   string `json:"status"`
	Title    string `json:"title"`
	Uri      string `json:"uri"`
	Volume   int    `json:"volume"`
}

func ListPlaylists() []string {
	req := buildRequest("api/listplaylists")
	resp := doRequest(req)

	response := []string{}
	err := json.Unmarshal(resp, &response)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func InvokeCommand(command string, arg string) CommandResponse {
	params := map[string]string{"cmd": command}

	if command == "play" && arg != "" {
		params["N"] = arg
	} else if command == "playplaylist" {
		params["name"] = arg
	} else if command == "volume" {
		params["volume"] = arg
	}

	req := buildRequestWithParams("api/commands", params)
	resp := doRequest(req)

	response := CommandResponse{}
	err := json.Unmarshal(resp, &response)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func GetState() StateResponse {
	req := buildRequest("api/getState")
	resp := doRequest(req)

	response := StateResponse{}
	err := json.Unmarshal(resp, &response)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func buildRequest(path string) *http.Request {
	params := map[string]string{}
	return buildRequestWithParams(path, params)
}

func buildRequestWithParams(path string, params map[string]string) *http.Request {
	username := config.Fetch("music_username")
	password := config.Fetch("music_password")
	domain := config.Fetch("music_domain")

	url := fmt.Sprintf("%s/%s", domain, path)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(username, password)

	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	return req
}

func doRequest(req *http.Request) []byte {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bodyBytes
}
