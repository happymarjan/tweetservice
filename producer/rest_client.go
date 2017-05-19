package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func restClientMain(tweeData *TweetData) {
	//var data []byte
	var resp *http.Response
	var req *http.Request

	type TweetData struct {
		StatusText      string
		StatusTimestamp string
		StatusAuthor    string
		StatusId        string
	}

	/*fmt.Println("In REST client")
	fmt.Println(tweeData.StatusId)
	fmt.Println(tweeData.StatusAuthor)
	fmt.Println(tweeData.StatusTimestamp)
	fmt.Println(tweeData.StatusText)*/

	toSendStr := fmt.Sprintf("{\"NewsID\":\"%s\",\"NewsAuthor\":\"%s\", \"NewsDate\": \"%s\", \"NewsText\":\"%s\"}", tweeData.StatusId, tweeData.StatusAuthor, tweeData.StatusTimestamp, tweeData.StatusText)
	body := strings.NewReader(toSendStr)
	//body := strings.NewReader(`{"NewsID":"1","NewsAuthor":"testauthor", "NewsDate": "2011-07-14T19:43:37+0100", "NewsText":"Happy day"}`)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8082/News/", body)
	if err == nil {
		req.Header.Set("Content-Type", "Content-Type: application/json")
		req.Header.Set("Accept", "application/json")
		debug(httputil.DumpRequestOut(req, true))
		resp, err = http.DefaultClient.Do(req)

		//resp, err = (&http.Client{}).Do(req)
	}
	if err == nil {
		defer resp.Body.Close()
		//debug(httputil.DumpResponse(resp, true))
		//data, err = ioutil.ReadAll(resp.Body)
	}

	if err == nil {
		//fmt.Println(string(data))
		fmt.Println("resp status: ", resp.Status)
	} else {
		log.Fatalf("Error: %s", err)
	}
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n", data)
	} else {
		log.Fatalf("%s\n", err)
	}
}
