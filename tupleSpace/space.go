package tupleSpace

import (
	//"code.google.com/p/go-uuid/uuid"
)
import (
	"fmt"
	"net/http"
	"log"
	"net/url"
	"encoding/json"
	"bytes"
)

var restURL = "http://128.233.173.24:8080/LindaRestServer/tuple"

type Item struct{
	Key string `json:"key"`
	Data string `json:"data"`
}

func Take(key string) Item {
	safeKey := url.QueryEscape(key)
	url := fmt.Sprintf(restURL+"/%s",safeKey)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		//return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		//return
	}
	defer resp.Body.Close()
	var record Item
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	if record.Data=="0"{
		record.Key=""
	}
	return record

}

func Write(item Item) {
	url := fmt.Sprintf(restURL)
	itemByte,_:=json.Marshal(item)

	// Build the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(itemByte))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()
}
