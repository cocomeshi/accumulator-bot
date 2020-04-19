package internal

import (
	"fmt"
	"log"

	"github.com/cocomeshi/accumulator-bot/infrastructure"
	repo "github.com/cocomeshi/accumulator-bot/infrastructure"
	api "github.com/cocomeshi/accumulator-bot/interface"
)

func AdditionalUpdate(key string) {

	db := repo.GetInstance()

	placeIdArr, err := repo.SelectIds(db)
	if err != nil {
		fmt.Println(err)
	}

	for _, placeId := range placeIdArr {

		detailData, err := api.DetailSearch(key, placeId)
		if err != nil {
			fmt.Println(err)
		}
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
			continue
		}
		err = infrastructure.Update(placeId, detailData.ResultData.Address, db)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			continue
		}
		tx.Commit()
	}

}
