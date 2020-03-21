package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/cocomeshi/accumulator-bot/data"
	repo "github.com/cocomeshi/accumulator-bot/infrastructure"
	api "github.com/cocomeshi/accumulator-bot/interface"
	_ "github.com/go-sql-driver/mysql"
)

const breakCounterLimit int = 70

func Scanning(key string) {
	// 走査対象データファイルからJSONデータを読み込む
	area := getRegion()
	kinkiArea := area.Kinki
	amagasakiRange := kinkiArea[1].PointRange
	latRange := amagasakiRange.Lat
	lonRange := amagasakiRange.Lon
	point := data.Coordinates{
		Latitude:  latRange[0],
		Longitude: lonRange[0],
	}
	ps := data.PointerScope{
		Current: point,
		PointRange: data.Range{
			Lat: latRange,
			Lon: lonRange,
		},
	}
	fmt.Println(ps)
	// 指定の座標にて検索
	regionSearch(key, ps)

}

func regionSearch(key string, pointScope data.PointerScope) {

	breakCounter := 0
	go func() {
		for {
			fmt.Println("counter : " + strconv.Itoa(breakCounter))
			if breakCounter == 0 {
				Search(key, pointScope.Current)
			} else {
				err := pointScope.Next()
				if err != nil {
					break
				}
				Search(key, pointScope.Current)
			}
			if breakCounter >= breakCounterLimit {
				// 連続して送信するリクエスト数の上限を超えた場合、
				// 一秒間のインターバルを挟む
				// Places APIのリクエスト上限は、100/sec のため
				breakCounter = 0
				time.Sleep(60 * time.Second)
			} else {
				breakCounter++
			}

		}
	}()

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

// 走査地域データ（JSON）を取得する
func getRegion() data.Area {
	raw, err := ioutil.ReadFile("./searchRegion.json")
	if err != nil {
		fmt.Println(err)
	}
	var data data.Area
	json.Unmarshal(raw, &data)
	fmt.Println(data)
	return data
}
