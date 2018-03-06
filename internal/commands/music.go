package commands

import (
	"encoding/json"
	"fmt"
	"github.com/codeclimate/hestia/internal/config"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Music struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

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

func (c Music) Run() {
	if c.Input.Args == "playlists" {
		c.listPlaylists()
	} else if c.Input.Args == "state" {
		c.getState()
	} else {
		c.invokeCommand()
	}
}

func (c Music) getState() {
	req := c.buildRequest()
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	response := StateResponse{}
	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		log.Fatal(err)
	}

	message := fmt.Sprintf("<@%s>: [%s] %s by %s on %s", c.User, response.Status, response.Title, response.Artist, response.Album)

	c.Notifier.Log(message)
}

func (c Music) listPlaylists() {
	req := c.buildRequest()
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	response := []string{}
	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		log.Fatal(err)
	}

	message := fmt.Sprintf("<@%s>:\n%s", c.User, strings.Join(response, "\n"))

	c.Notifier.Log(message)
}

func (c Music) invokeCommand() {
	req := c.buildRequest()
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	response := CommandResponse{}
	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		log.Fatal(err)
	}

	var message string

	if response.Error != "" {
		message = fmt.Sprintf("<@%s>: %s", c.User, response.Error)
	} else {
		message = fmt.Sprintf("<@%s>: %s", c.User, response.Response)
	}

	c.Notifier.Log(message)
}

func (c Music) buildRequest() *http.Request {
	username := config.Fetch("music_username")
	password := config.Fetch("music_password")
	domain := config.Fetch("music_domain")

	var url string

	if c.Input.Args == "playlists" {
		url = fmt.Sprintf("%s/%s", domain, "api/listplaylists")
	} else if c.Input.Args == "state" {
		url = fmt.Sprintf("%s/%s", domain, "api/getState")
	} else {
		url = fmt.Sprintf("%s/%s", domain, "api/commands")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(username, password)

	if strings.HasSuffix(url, "/commands") {
		parts := strings.Split(c.Input.Args, " ")
		subcommand := parts[0]

		q := req.URL.Query()
		q.Add("cmd", subcommand)

		switch subcommand {
		case "pause":
			if parts[1] != "" {
				q.Add("N", parts[1])
			}
		case "playplaylist":
			q.Add("name", parts[1])
		case "volume":
			q.Add("volume", parts[1])
		}

		req.URL.RawQuery = q.Encode()
	}

	return req
}
