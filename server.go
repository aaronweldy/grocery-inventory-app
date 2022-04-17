package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Request struct {
	Zip      string
	Products []string
}

type Response struct {
	Store []struct {
		Name string
		Address string
		LocationId string
		PercentAvailable float
		Missing []string
	}
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
			Hours        string
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

func ProductHandler(term string, locationId string, token string) {
	url := fmt.Sprintf("https://api.kroger.com/v1/products?filter.term=%s&filter.locationId=%s", term, locationId)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var prods Product
	err := json.Unmarshal(body, &prods)
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
		fmt.Println(fmt.Sprintf("Available in %s", locationId))
	}	else {
		fmt.Println(fmt.Sprintf("Not available in %s", locationId))
	}


}

func LocationHandler(zip string, token string) []string {
	url := fmt.Sprintf("https://api.kroger.com/v1/locations?filter.zipCode.near=%s&filter.limit=3", zip)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//
	// fmt.Println(res)
	// fmt.Println(string(body))

	var locs Location
	err := json.Unmarshal(body, &locs)
	if err != nil {
		panic(err)
	}

	locIdArray := make([]string, 0)

	for _, prod := range locs.Data {
		locIdArray = append(locIdArray, prod.LocationId)
	}
	return locIdArray
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
		Zip:      "48104",
		Products: []string{"goat cheese", "tempeh"},
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

	// find nearby stores and search for products
	var response Response

	locationIds := LocationHandler(user_input.Zip, t.Access_token)
	for _, id := range locationIds {
		for _, prod := range user_input.Products {
			ProductHandler(prod, id, t.Access_token)
		}


	}

}

func main() {
	fmt.Printf("Starting server at port 8080\n")
	http.HandleFunc("/list", listHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
