package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cocomeshi/accumulator-bot/data"
)

func Fetch(key string) {

	datas, err := get(key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v¥n", datas)

	// mongoに接続
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:28001"))

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// collections := client.Database("cocomeshi").Collection("meshiya")

	// for _, rest := range datas.Restaurants {
	// 	insertResult, err := collection.InsertOne(ctx, rest)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(insertResult)
	// }

}

func get(key string) (data.RestaurantResponse, error) {

	latitude := 34.726799
	longitude := 135.401687
	strLatitude := strconv.FormatFloat(latitude, 'f', -1, 64)
	strLongitude := strconv.FormatFloat(longitude, 'f', -1, 64)
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?" + "location=" + strLatitude + "," + strLongitude + "&radius=3000&type=restaurant&key=" + key
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
