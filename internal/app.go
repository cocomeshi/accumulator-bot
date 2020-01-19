package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cocomeshi/accumulator-bot/data"
)

func Fetch(key string) {

	datas, err := get(key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(datas)

}

func get(key string) (data.RestaurantResponse, error) {

	keyword := url.QueryEscape("マクドナルド")
	url := "https://api.gnavi.co.jp/RestSearchAPI/v3/?keyid=" + key + "&name=" + keyword
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
