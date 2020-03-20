package infrastructure

import (
	_ "github.com/go-sql-driver/mysql"
	data "github.com/yuta4j1/accumulator-bot/data"
)

func Insert(db *sql.DB, datas []data.Restaurant) {
	tx, err := cnn.Begin()
	if err != nil {
		log.Fatal(err)
	}
	ins, err := cnn.Prepare("insert into cocomeshi.restaurant(id, name, address, open_time, close_time, longitude, latitude, area_kind) values (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range datas.Restaurants {
		fmt.Println(d)
		r, e := ins.Exec(d.Id, d.Name, "", "", "", d.Geometry.Location.Longitude, d.Geometry.Location.Latitude, "kinki")
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(r)
	}
	tx.Commit()
}