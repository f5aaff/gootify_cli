package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
)

var (
	supportedControls = []string{"play", "pause", "next", "previous"}
)

func playControls(args []*string) {
	if slices.Contains(supportedControls, *args[0]) {
		switch args[1] {
		case nil:
			res, err := http.Get(baseURL + "devices/player" + *args[0])
			if err != nil {
				fmt.Println("error executing " + *args[0] + " function: " + err.Error())
			}
			if res.StatusCode == http.StatusOK {
				fmt.Println("success")
			}
		default:
			if *args[1] != "" {
				res, err := http.Get(baseURL + "devices/player/playCustom/" + *args[1])
				if err != nil {
					fmt.Println("error executing " + *args[0] + " function: " + err.Error())
				}
				if res.StatusCode == http.StatusOK {
					fmt.Println("success")
				}
			}

		}
		_, err := http.Get(baseURL + "devices/player/")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func deviceList() {
	res, err := http.Get(baseURL + "devices/list")
	if err != nil {
		fmt.Println("error getting devices" + err.Error())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {

		fmt.Println("error reading device list" + err.Error())
	}
	fmt.Printf(string(body))
}

type SearchRequest struct {
	Query  string            `json:"query"`
	Tags   map[string]string `json:"tags"`
	Types  []string          `json:"types"`
	Market string            `json:"market"`
	Limit  int               `json:"limit"`
}

var (
	query string
	tags  map[string]string
	types []string
	limit int
)

func search() {
	data := &SearchRequest{}
	data.Query = query
	data.Tags = nil
	data.Limit = limit
	data.Types = types
	data.Market = "ES"

	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error marshalling query payload", err.Error())
	}
	res, err := http.Post(baseURL+"search", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("error sending search request", err.Error())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response ", err.Error())
	}
	fmt.Println(string(body))
	var resList *searchResponse
	err = json.Unmarshal(body, resList)
	if err != nil {
		fmt.Println("error unmarshalling result ", err.Error())
	}

	//r := *resList
	//tableprinter.Print(os.Stdout, r[0].Tracks.Items)
}
