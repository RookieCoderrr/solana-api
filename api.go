package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	solanaApi = "https://public-api.solscan.io/account/transactions"
	account = "7aMFk2CEo33Sb9R1a1W6q2HxZ1Z97hbLedDy3ZgHJ4i6"
)

func getTransactionsByAccount(fileName string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET",solanaApi, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := url.Values{}
	q.Add("account",account)
	q.Add("limit","50")
	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	err = ioutil.WriteFile(fileName, respBody, 0644)
	if err != nil {
		fmt.Println(err)
	}

}
