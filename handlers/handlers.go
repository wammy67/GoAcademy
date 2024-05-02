package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type HTTPClient interface {
    Get(url string) (*http.Response, error)
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(client HTTPClient, url string, target any) error {
    r, err := client.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func GetCoffeess(client HTTPClient) Coffees {
coffees := new(Coffees) 
getJson(client, "https://api.sampleapis.com/coffee/hot", coffees)

return (*coffees)
}

func GetWines(client HTTPClient) Wines {
	wines := new(Wines) 
	getJson(client, "https://api.sampleapis.com/wines/reds", wines)
	
	return (*wines)
}


type Coffee struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Image       string   `json:"image"`
	ID          int      `json:"id"`
}

type Coffees []Coffee

type Rating struct {
    Average  string `json:"average"`
    Reviews  string `json:"reviews"`
}

type Wine struct {
    Winery   string `json:"winery"`
    Wine     string `json:"wine"`
    Rating   Rating `json:"rating"`
    Location string `json:"location"`
    Image    string `json:"image"`
    ID       int    `json:"id"`
}

type Wines []Wine
