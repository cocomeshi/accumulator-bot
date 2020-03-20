package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	data "github.com/cocomeshi/accumulator-bot/data"
	_ "github.com/go-sql-driver/mysql"
)

func Insert(db *sql.DB, datas []data.Restaurant) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	ins, err := db.Prepare("insert into cocomeshi.restaurant(id, name, address, open_time, close_time, longitude, latitude, area_kind) values (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range datas {
		fmt.Println(d)
		r, e := ins.Exec(d.Id, d.Name, "", "", "", d.Geometry.Location.Longitude, d.Geometry.Location.Latitude, "kinki")
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(r)
	}
	tx.Commit()
}
