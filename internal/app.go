package internal

import (
	"fmt"

	"github.com/cocomeshi/accumulator-bot/data"
	repo "github.com/cocomeshi/accumulator-bot/infrastructure"
	api "github.com/cocomeshi/accumulator-bot/interface"
	_ "github.com/go-sql-driver/mysql"
)

func Exec(key string) {

	// モックデータ
	point := data.Coordinates{
		Latitude:  34.726799,
		Longitude: 135.401687,
	}
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
