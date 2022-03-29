package main

import (
	"fmt"
	"log"
	"net/http"
  "io/ioutil"
	"encoding/json"
	"strings"
)

type Request struct {
    Zip string
    Products string
}

type Token struct {
	expires_in int
	access_token string
}

func ProductHandler(term string, locationId string) {
  url := fmt.Sprintf("https://api.kroger.com/v1/products?filter.term=%s&filter.locationId=%s", term, locationId)

  req, _ := http.NewRequest("GET", url, nil)

  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImprdSI6Imh0dHBzOi8vYXBpLmtyb2dlci5jb20vdjEvLndlbGwta25vd24vandrcy5qc29uIiwia2lkIjoiWjRGZDNtc2tJSDg4aXJ0N0xCNWM2Zz09IiwidHlwIjoiSldUIn0.eyJhdWQiOiJlZWNzNDk3cHQyLWVmZDQ0YzllNTJkZjcwYTA2YjcwYTYxOTE0MmE0YjU2ODgyNDgzOTM0Mzk0ODc0MjgzMyIsImV4cCI6MTY0ODQ4MDYyOCwiaWF0IjoxNjQ4NDc4ODIzLCJpc3MiOiJhcGkua3JvZ2VyLmNvbSIsInN1YiI6IjU0Y2I0MTc5LWExYTQtNWE2MC05ODUxLTI2YTRiNzczZTVlZSIsInNjb3BlIjoicHJvZHVjdC5jb21wYWN0IiwiYXV0aEF0IjoxNjQ4NDc4ODI4MDYzNjM4NTc3LCJhenAiOiJlZWNzNDk3cHQyLWVmZDQ0YzllNTJkZjcwYTA2YjcwYTYxOTE0MmE0YjU2ODgyNDgzOTM0Mzk0ODc0MjgzMyJ9.aAPRlMhWsvS7uc-_Ez56D9RheeobXCwqJKoR6olxy5GGmEInL8IvJcB15AZ2sHELzdPf51DAzXgW9CTuoSsTh0IJ7OtuCRWgkLpzHwpbhkRjlLT0V9hvszj4NhgwyA5y7J6at7FxLGfuZ1E_1c024vcVROB2pRzYXeeAdsNfLSQbqKnXd6uwToQ5qk_JSmJdLiTxCgGxY-5YIT7z-1VCiARqyjm7aY6OhT8jEbA5YjHCOLbepybH6GOxvUZlgn3uN8Uo8Yp-Vd_I2GKkgkunK0JxML7TEVDoXVFtNHaprJ9uARLUAF2B7LgnhrVRX3Y9d1dAoC8pi3L8JWggRSyOIg")

  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  fmt.Println(res)
  fmt.Println(string(body))
}

func LocationHandler(zip string) {
	url := fmt.Sprintf("https://api.kroger.com/v1/locations?filter.zipCode.near=%s", zip)

  req, _ := http.NewRequest("GET", url, nil)

  req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImprdSI6Imh0dHBzOi8vYXBpLmtyb2dlci5jb20vdjEvLndlbGwta25vd24vandrcy5qc29uIiwia2lkIjoiWjRGZDNtc2tJSDg4aXJ0N0xCNWM2Zz09IiwidHlwIjoiSldUIn0.eyJhdWQiOiJlZWNzNDk3cHQyLWVmZDQ0YzllNTJkZjcwYTA2YjcwYTYxOTE0MmE0YjU2ODgyNDgzOTM0Mzk0ODc0MjgzMyIsImV4cCI6MTY0ODQ4MDYyOCwiaWF0IjoxNjQ4NDc4ODIzLCJpc3MiOiJhcGkua3JvZ2VyLmNvbSIsInN1YiI6IjU0Y2I0MTc5LWExYTQtNWE2MC05ODUxLTI2YTRiNzczZTVlZSIsInNjb3BlIjoicHJvZHVjdC5jb21wYWN0IiwiYXV0aEF0IjoxNjQ4NDc4ODI4MDYzNjM4NTc3LCJhenAiOiJlZWNzNDk3cHQyLWVmZDQ0YzllNTJkZjcwYTA2YjcwYTYxOTE0MmE0YjU2ODgyNDgzOTM0Mzk0ODc0MjgzMyJ9.aAPRlMhWsvS7uc-_Ez56D9RheeobXCwqJKoR6olxy5GGmEInL8IvJcB15AZ2sHELzdPf51DAzXgW9CTuoSsTh0IJ7OtuCRWgkLpzHwpbhkRjlLT0V9hvszj4NhgwyA5y7J6at7FxLGfuZ1E_1c024vcVROB2pRzYXeeAdsNfLSQbqKnXd6uwToQ5qk_JSmJdLiTxCgGxY-5YIT7z-1VCiARqyjm7aY6OhT8jEbA5YjHCOLbepybH6GOxvUZlgn3uN8Uo8Yp-Vd_I2GKkgkunK0JxML7TEVDoXVFtNHaprJ9uARLUAF2B7LgnhrVRX3Y9d1dAoC8pi3L8JWggRSyOIg")

  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  fmt.Println(res)
  fmt.Println(string(body))
}

func basicGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
	}

	fmt.Fprintf(w, "Success!")
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/list" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
	}

	// body = r.Body
	req_body := Request{
		"48104",
		"tempeh tofu",
	}
	zip := req_body.Zip
	products := req_body.Products

	// refresh token
	url := "https://api.kroger.com/v1/connect/oauth2/token"

  payload := strings.NewReader("grant_type=client_credentials&scope=product.compact")

  req, _ := http.NewRequest("POST", url, payload)

  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Authorization", "Basic ZWVjczQ5N3B0Mi1lZmQ0NGM5ZTUyZGY3MGEwNmI3MGE2MTkxNDJhNGI1Njg4MjQ4MzkzNDM5NDg3NDI4MzM6NGk1a1lyVXp1MFZMWjR3Vk01SUllR1ZHXzc0OXpKTjRaVUhnckI3cg==")

  res, _ := http.DefaultClient.Do(req)


  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)
	var t Token
	err := json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}


	// find nearby stores
	url = fmt.Sprintf("https://api.kroger.com/v1/locations?filter.zipCode.near=%s&filter.limit=3", zip)

	req, _ = http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImprdSI6Imh0dHBzOi8vYXBpLmtyb2dlci5jb20vdjEvLndlbGwta25vd24vandrcy5qc29uIiwia2lkIjoiWjRGZDNtc2tJSDg4aXJ0N0xCNWM2Zz09IiwidHlwIjoiSldUIn0.eyJhdWQiOiJlZWNzNDk3cHQyLWVmZDQ0YzllNTJkZjcwYTA2YjcwYTYxOTE0MmE0YjU2ODgyNDgzOTM0Mzk0ODc0MjgzMyIsImV4cCI6MTY0ODQ4MDYyOCwiaWF0IjoxNjQ4NDc4ODIzLCJpc3MiOiJhcGkua3JvZ2VyLmNvbSIsInN1YiI6IjU0Y2I0MTc5LWExYTQtNWE2MC05ODUxLTI2YTRiNzczZTVlZSIsInNjb3BlIjoicHJvZHVjdC5jb21wYWN0IiwiYXV0aEF0IjoxNjQ4NDc4ODI4MDYzNjM4NTc3LCJhenAiOiJlZWNzNDk3cHQyLWVmZDQ0YzllNTJkZjcwYTA2YjcwYTYxOTE0MmE0YjU2ODgyNDgzOTM0Mzk0ODc0MjgzMyJ9.aAPRlMhWsvS7uc-_Ez56D9RheeobXCwqJKoR6olxy5GGmEInL8IvJcB15AZ2sHELzdPf51DAzXgW9CTuoSsTh0IJ7OtuCRWgkLpzHwpbhkRjlLT0V9hvszj4NhgwyA5y7J6at7FxLGfuZ1E_1c024vcVROB2pRzYXeeAdsNfLSQbqKnXd6uwToQ5qk_JSmJdLiTxCgGxY-5YIT7z-1VCiARqyjm7aY6OhT8jEbA5YjHCOLbepybH6GOxvUZlgn3uN8Uo8Yp-Vd_I2GKkgkunK0JxML7TEVDoXVFtNHaprJ9uARLUAF2B7LgnhrVRX3Y9d1dAoC8pi3L8JWggRSyOIg")

	res, _ = http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	// // store_ids = []
	//
	// search for products
	locationId := "01400943"
	url = fmt.Sprintf("https://api.kroger.com/v1/products?filter.term=%s&filter.locationId=%s", products, locationId)

  req, _ = http.NewRequest("GET", url, nil)

  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImprdSI6Imh0dHBzOi8vYXBpLmtyb2dlci5jb20vdjEvLndlbGwta25vd24vandrcy5qc29uIiwia2lkIjoiWjRGZDNtc2tJSDg4aXJ0N0xCNWM2Zz09IiwidHlwIjoiSldUIn0.eyJhdWQiOiJlZWNzNDk3cHQyLWVmZDQ0YzllNTJkZjcwYTA2YjcwYTYxOTE0MmE0YjU2ODgyNDgzOTM0Mzk0ODc0MjgzMyIsImV4cCI6MTY0ODQ4MDYyOCwiaWF0IjoxNjQ4NDc4ODIzLCJpc3MiOiJhcGkua3JvZ2VyLmNvbSIsInN1YiI6IjU0Y2I0MTc5LWExYTQtNWE2MC05ODUxLTI2YTRiNzczZTVlZSIsInNjb3BlIjoicHJvZHVjdC5jb21wYWN0IiwiYXV0aEF0IjoxNjQ4NDc4ODI4MDYzNjM4NTc3LCJhenAiOiJlZWNzNDk3cHQyLWVmZDQ0YzllNTJkZjcwYTA2YjcwYTYxOTE0MmE0YjU2ODgyNDgzOTM0Mzk0ODc0MjgzMyJ9.aAPRlMhWsvS7uc-_Ez56D9RheeobXCwqJKoR6olxy5GGmEInL8IvJcB15AZ2sHELzdPf51DAzXgW9CTuoSsTh0IJ7OtuCRWgkLpzHwpbhkRjlLT0V9hvszj4NhgwyA5y7J6at7FxLGfuZ1E_1c024vcVROB2pRzYXeeAdsNfLSQbqKnXd6uwToQ5qk_JSmJdLiTxCgGxY-5YIT7z-1VCiARqyjm7aY6OhT8jEbA5YjHCOLbepybH6GOxvUZlgn3uN8Uo8Yp-Vd_I2GKkgkunK0JxML7TEVDoXVFtNHaprJ9uARLUAF2B7LgnhrVRX3Y9d1dAoC8pi3L8JWggRSyOIg")

  res, _ = http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ = ioutil.ReadAll(res.Body)

  fmt.Println(res)
  fmt.Println(string(body))
}

func main() {
	fmt.Printf("Starting server at port 8080\n")
	http.HandleFunc("/hello", basicGetHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	// LocationHandler("48104")
	// ProductHandler("tempeh", "01400943")
}
