package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

func NewBeverageService(client HTTPClient) *BeverageService {
	return &BeverageService{
		client: client,
	}
}

type BeverageService struct {
	client HTTPClient
}

func (b *BeverageService) GetBoth() Beverages {
	var wg sync.WaitGroup
	beveragesChan := make(chan Beverage)

	wg.Add(2)
    
	// Fetch coffees
	go func() {
		defer wg.Done()
		coffees := b.GetCoffees()
		for _, coffee := range coffees {
			beveragesChan <- coffee
		}
	}()

	// Fetch wines
	go func() {
		defer wg.Done()
		wines := b.GetWines()
		for _, wine := range wines {
			beveragesChan <- wine
		}
	}()

	go func() {
		wg.Wait()
		close(beveragesChan)
	}()

	var beverages Beverages
	for beverage := range beveragesChan {
		beverages = append(beverages, beverage)
	}

	return beverages
}

func (b *BeverageService) GetCoffees() Beverages {
	coffees := new(Coffees)
	b.getJson("https://api.sampleapis.com/coffee/hot", coffees)
	var beverages Beverages
	for _, coffee := range *coffees {
		beverages = append(beverages, coffee)
	}
	//fmt.Println(beverages, "coffee nika")
	return beverages
}

func (b *BeverageService) GetWines() Beverages {
	wines := new(Wines)
	b.getJson("https://api.sampleapis.com/wines/reds", wines)
	var beverages Beverages
	for _, wine := range *wines {
		beverages = append(beverages, wine)
	}
   // fmt.Println(beverages, "wines nika")
	return beverages
}

func (b *BeverageService) getJson(url string, target any) error {
	r, err := b.client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type Beverage interface {
	GetTitle() string
	GetID() int
}

type Coffee struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
}

type Wine struct {
	Title string `json:"wine"`
	ID    int    `json:"id"`
}

type Beverages []Beverage

type Coffees []Coffee

type Wines []Wine

func (c Coffee) GetTitle() string {
	return c.Title
}

func (c Coffee) GetID() int {
	return c.ID
}

func (w Wine) GetTitle() string {
	return w.Title
}

func (w Wine) GetID() int {
	return w.ID
}
