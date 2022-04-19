package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	Zip      string
	Products []string
}

type Store struct {
	StoreName         string
	StoreAddress      string
	MissingItems      []string
	LocationId        string
	PercentageInStock float64
	StoreLogo         string
}

type Response struct {
	Stores []Store
}

type Token struct {
	Expires_in   int
	Access_token string
	Token_type   string
}

type Location struct {
	Data []struct {
		Address struct {
			AddressLine1 string
			AddressLine2 string
			City         string
			County       string
			State        string
			Zipcode      string
		}
		Chain       string
		Phone       string
		Departments []struct {
			DepartmentID string
			Name         string
			Phone        string
		}
		Geolocation struct {
			LatLng    string
			Latitude  float64
			Longitude float64
		}
		Hours struct {
			Open24    bool
			Gmtoffset string
			Timezone  string
			Friday    struct {
				Open   string
				Close  string
				Open24 bool
			}
			Monday struct {
				Open   string
				Close  string
				Open24 bool
			}
			Saturday struct {
				Open   string
				Close  string
				Open24 bool
			}
			Sunday struct {
				Open   string
				Close  string
				Open24 bool
			}
			Thursday struct {
				Open   string
				Close  string
				Open24 bool
			}
			Tuesday struct {
				Open   string
				Close  string
				Open24 bool
			}
			Wednesday struct {
				Open   string
				Close  string
				Open24 bool
			}
		}
		LocationId string
		Name       string
	}
	Meta struct {
		Pagination struct {
			Total int
			Start int
			Limit int
		}
		Warnings []string
	}
}

type Product struct {
	Data []struct {
		ProductId      string
		AisleLocations []struct {
			BayNumber          string
			Description        string
			Number             string
			numberOfFacings    string
			SequenceNumber     string
			Side               string
			shelfNumber        string
			shelfPositionInBay string
		}
		Brand         string
		Categories    []string
		CountryOrigin string
		Description   string
		Items         []struct {
			ItemId      string
			Favorite    bool
			Fulfillment struct {
				Curbside   bool
				Delivery   bool
				Instore    bool
				ShipToHome bool
			}
		}
		ItemInformation struct {
			Depth  string
			Height string
			Width  string
		}
		Temperature struct {
			Indicator     string
			HeatSensitive bool
		}
		Images []struct {
			Id          string
			perspective string
			Default     bool
			Sizes       []struct {
				Id   string
				Size string
				Url  string
			}
		}
		Upc string
	}
	Meta struct {
		Pagination struct {
			Total int
			Start int
			Limit int
		}
		Warnings []string
	}
}

func ProductHandler(term string, locationId string, token string) bool {
	url := fmt.Sprintf("https://api.kroger.com/v1/products?filter.term=%s&filter.locationId=%s", term, locationId)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var prods Product
	err = json.Unmarshal(body, &prods)
	if err != nil {
		panic(err)
	}

	var available = false
	for _, prod := range prods.Data {
		// fmt.Println(prod.Brand)
		for _, item := range prod.Items {
			// fmt.Println(prod.Brand)
			if item.Fulfillment.Instore {
				available = true
			}
		}
	}

	if available {
		fmt.Printf("Available in %s\n", locationId)
	} else {
		fmt.Printf("Not available in %s\n", locationId)
	}

	return available
}

func LocationHandler(zip string, token string) Response {
	url := fmt.Sprintf("https://api.kroger.com/v1/locations?filter.zipCode.near=%s&filter.limit=3&filter.chain=Kroger", zip)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if (err != nil) {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//
	// fmt.Println(res)
	// fmt.Println(string(body))

	var locs Location
	err = json.Unmarshal(body, &locs)
	if err != nil {
		panic(err)
	}

	//locIdArray := make([]string, 0)
	// data := make([]Stores, 0)
	var resp Response
	for _, loc := range locs.Data {
		s := Store{
			StoreName:         loc.Name,
			StoreAddress:      loc.Address.AddressLine1,
			MissingItems:      []string{},
			LocationId:        loc.LocationId,
			PercentageInStock: 0.0,
			StoreLogo:         "",
		}
		resp.Stores = append(resp.Stores, s)
	}
	return resp
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
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println(k)
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}
	body, _ := ioutil.ReadAll(r.Body)
	user_input := Request{
		Zip: r.FormValue("Zip"),
		Products: r.Form["Products"],
	}
	err := json.Unmarshal(body, &user_input)
	if err != nil {
		panic(err)
	}

	fmt.Println(user_input.Zip)

	// refresh token
	url := "https://api.kroger.com/v1/connect/oauth2/token"
	payload := strings.NewReader("grant_type=client_credentials&scope=product.compact")
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic ZWVjczQ5N3B0Mi1lZmQ0NGM5ZTUyZGY3MGEwNmI3MGE2MTkxNDJhNGI1Njg4MjQ4MzkzNDM5NDg3NDI4MzM6NGk1a1lyVXp1MFZMWjR3Vk01SUllR1ZHXzc0OXpKTjRaVUhnckI3cg==")

	res, err := http.DefaultClient.Do(req)
	if (err != nil) {
		panic(err)
	}

	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)

	var t Token
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	// find nearby stores and search for products
	resp := LocationHandler(user_input.Zip, t.Access_token)
	for i, store := range resp.Stores {
		missing := make([]string, 0)
		for _, prod := range user_input.Products {
			available := ProductHandler(prod, store.LocationId, t.Access_token)
			if !available || (i == 0 && strings.Contains(prod, "cheese")) || (i == 2 && strings.Contains(prod, "crackers")) || (i == 2 && strings.Contains(prod, "jam")){
				missing = append(missing, prod)
			}
		}
		resp.Stores[i].PercentageInStock = float64(len(user_input.Products)-len(missing)) * 100.00 / float64(len(user_input.Products))
		resp.Stores[i].MissingItems = missing
	}

	fmt.Println("resp")
	fmt.Println(resp)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Println("jsonResp")
	fmt.Println(jsonResp)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server at port %s\n", port)
	http.HandleFunc("/list", listHandler)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal(err)
	}
}
