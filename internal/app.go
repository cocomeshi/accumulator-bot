package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cocomeshi/accumulator-bot/data"
	repo "github.com/cocomeshi/accumulator-bot/infrastructure"
	api "github.com/cocomeshi/accumulator-bot/interface"
	_ "github.com/go-sql-driver/mysql"
)

const (
	ByLatitude  float64 = 0.0555560
	ByLongitude float64 = 0.0666672
)

func Scanning(key string) {
	area := getRegion()
	kinkiArea := area.Kinki
	amagasakiRange := kinkiArea[1].PointRange
	latRange := amagasakiRange.Lat
	lonRange := amagasakiRange.Lon
	point := data.Coordinates{
		Latitude:  latRange[0],
		Longitude: lonRange[0],
	}
	// 指定の座標にて検索
	Search(key, point)

}

// 走査地域データ（JSON）を取得する
func getRegion() data.Area {
	raw, err := ioutil.ReadFile("./searchRegion.json")
	if err != nil {
		fmt.Println(err)
	}
	var data data.Area
	json.Unmarshal(raw, &data)
	fmt.Println("get region data!!")
	fmt.Println(data)
	return data
}

func Search(key string, point data.Coordinates) {

	datas, err := api.NearbySearch(key, point)
	if err != nil {
		fmt.Println(err)
	}
	db, err := repo.NewInstance()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	repo.Insert(db, datas.Restaurants)

}
