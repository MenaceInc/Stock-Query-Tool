package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const stockEndpoint = "https://api.iextrading.com/1.0/stock/"

func main() {
	for {
		print("Please input a stock code: ")
		var stock string
		fmt.Scanln(&stock)

		if stock == "exit" || stock == "quit" {
			break
		}

		println(requestCompanyData(stock))
		println("")
	}
	println("Goodbye!")
}

func requestCompanyData(stock string) (result string) {
	companyName, validName := requestCompanyName(stock)
	price, validPrice := requestPrice(stock)

	if validName && validPrice {
		result = "The current price per share for " + companyName + " is $" + price
	} else {
		result = "No data available for " + stock
	}

	return
}

func requestCompanyName(stock string) (result string, valid bool) {
	valid = false
	url := stockEndpoint + stock + "/company"
	response, err := http.Get(url)

	if err != nil {
		return
	}

	defer response.Body.Close()
	jsonData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	var f interface{}
	err = json.Unmarshal(jsonData, &f)

	if err != nil {
		return
	}

	temp := f.(map[string]interface{})

	for key, value := range temp {
		if key == "companyName" {
			result = value.(string)
			valid = true
			break
		}
	}

	return
}

func requestPrice(stock string) (result string, valid bool) {
	valid = false
	url := stockEndpoint + stock + "/price"
	response, err := http.Get(url)

	if err != nil {
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil || string(body) == "Unknown symbol" {
		return
	}

	result += string(body)
	valid = true

	return
}
