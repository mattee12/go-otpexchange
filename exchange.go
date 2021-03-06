package otpexchange

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CURRENCY string

const (
	EUR CURRENCY = "EUR"
	HUF CURRENCY = "HUF"
	USD CURRENCY = "USD"
	GBP CURRENCY = "GBP"
)

func Convert(from, to CURRENCY, amount float64) (buy float64, sell float64) {
	data, err := json.Marshal(map[string]string{
		"exchangeType":   "FOREIGN_EXCHANGE",
		"baseAmount":     fmt.Sprintf("%f", amount),
		"baseCurrency":   string(from),
		"resultCurrency": string(to),
	})
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.otpbank.hu/apps/exchangerate/api/exchangerate/exchange", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Host", "www.otpbank.hu")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Cache-Control", "no-store, no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Expires", "0")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Content-Length", "86")
	req.Header.Add("Origin", "https://www.otpbank.hu")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://www.otpbank.hu/portal/hu/Arfolyamok/OTP")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var unm map[string]interface{}
	json.Unmarshal(body, &unm)
	buy = unm["resultAmount"].(float64)
	if err != nil {
		panic(err)
	}
	data, err = json.Marshal(map[string]string{
		"exchangeType":   "FOREIGN_EXCHANGE",
		"baseAmount":     fmt.Sprintf("%f", 1.0),
		"baseCurrency":   string(to),
		"resultCurrency": string(from),
	})
	if err != nil {
		panic(err)
	}
	req, err = http.NewRequest("POST", "https://www.otpbank.hu/apps/exchangerate/api/exchangerate/exchange", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Host", "www.otpbank.hu")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Cache-Control", "no-store, no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Expires", "0")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Content-Length", "86")
	req.Header.Add("Origin", "https://www.otpbank.hu")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://www.otpbank.hu/portal/hu/Arfolyamok/OTP")
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &unm)
	sell = amount * (1 / unm["resultAmount"].(float64))
	return
}
