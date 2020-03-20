package internal

import (
	"fmt"

	repo "github.com/cocomeshi/accumulator-bot/infrastructure"
	api "github.com/cocomeshi/accumulator-bot/interface"
	_ "github.com/go-sql-driver/mysql"
)

func Exec(key string) {

	datas, err := api.NearbySearch(key)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("%+v¥n", datas)
	db, err := repo.NewInstance()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	repo.Insert(db, datas)

}
