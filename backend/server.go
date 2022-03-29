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
	Expires_in int
	Access_token string
	Token_type string
}

func ProductHandler(term string, locationId string, token string) {
  url := fmt.Sprintf("https://api.kroger.com/v1/products?filter.term=%s&filter.locationId=%s", term, locationId)

  req, _ := http.NewRequest("GET", url, nil)

  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  // fmt.Println(res)
  // fmt.Println(string(body))
}

func LocationHandler(zip string, token string) {
	url := fmt.Sprintf("https://api.kroger.com/v1/locations?filter.zipCode.near=%s&filter.limit=3", zip)
  req, _ := http.NewRequest("GET", url, nil)

  req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  // _, _ := ioutil.ReadAll(res.Body)
	//
  // fmt.Println(res)
  // fmt.Println(string(body))
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

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(r.Body)
	// var user_input Request
	user_input := Request{
		Zip: "48104",
		Products: "tempeh",
	}
	err := json.Unmarshal(body, &user_input)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(user_input.Zip)

	// refresh token
	url := "https://api.kroger.com/v1/connect/oauth2/token"

  payload := strings.NewReader("grant_type=client_credentials&scope=product.compact")

  req, _ := http.NewRequest("POST", url, payload)

  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Authorization", "Basic ZWVjczQ5N3B0Mi1lZmQ0NGM5ZTUyZGY3MGEwNmI3MGE2MTkxNDJhNGI1Njg4MjQ4MzkzNDM5NDg3NDI4MzM6NGk1a1lyVXp1MFZMWjR3Vk01SUllR1ZHXzc0OXpKTjRaVUhnckI3cg==")

  res, _ := http.DefaultClient.Do(req)


  defer res.Body.Close()
  body, _ = ioutil.ReadAll(res.Body)

  // fmt.Println(res)
  // fmt.Println(string(body))

	var t Token
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	fmt.Println("refreshed access token")

	// // find nearby stores
	LocationHandler(user_input.Zip, t.Access_token)

	// // search for products
	locationId := "01400943"
	ProductHandler(user_input.Products, locationId, t.Access_token)
}

func main() {
	fmt.Printf("Starting server at port 8080\n")
	// http.HandleFunc("/hello", basicGetHandler)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
	http.HandleFunc("/list", listHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
