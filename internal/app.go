package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/cocomeshi/accumulator-bot/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Fetch(key string) {

	datas, err := get(key)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("%+v¥n", datas)

	// mongoに接続
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:28001"))

	if err != nil {
		fmt.Println(err)
	}

	collection := client.Database("cocomeshi").Collection("meshiya")

	for _, rest := range datas.Restaurants {
		insertResult, err := collection.InsertOne(ctx, rest)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(insertResult)
	}

}

func get(key string) (data.RestaurantResponse, error) {

	latitude := 34.726799
	longitude := 135.401687
	url := "https://api.gnavi.co.jp/RestSearchAPI/v3/?keyid=" + key + "&latitude=" + strconv.FormatFloat(latitude, 'f', -1, 64) + "&longitude=" + strconv.FormatFloat(longitude, 'f', -1, 64) + "&range=5&hit_per_page=100"
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
