package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cocomeshi/accumulator-bot/data"
)

// 'NearbySearch' API からデータ取得
func NearbySearch(key string, p data.Coordinates) (data.RestaurantResponse, error) {

	baseUrl := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?"
	radius := "3000"
	placetype := "restaurant"
	q := data.NewNearbysearchQuery(key, baseUrl, radius, placetype)
	q.SetLocation(p)
	url := q.QueryGen()
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var datas data.RestaurantResponse
	if err := json.Unmarshal(body, &datas); err != nil {
		fmt.Println(err)
	}
	return datas, err
}

// 'DetailSearch' API からデータ取得
func DetailSearch(key string, placeId string) (data.Detail, error) {
	url := "https://maps.googleapis.com/maps/api/place/details/json?place_id=" + placeId + "&key=" + key
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var detail data.Detail
	if err := json.Unmarshal(body, &detail); err != nil {
		fmt.Println(err)
	}
	return detail, err

}
