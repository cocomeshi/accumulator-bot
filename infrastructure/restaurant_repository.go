package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	data "github.com/cocomeshi/accumulator-bot/data"
	_ "github.com/go-sql-driver/mysql"
)

// 現在登録されている全place IDの配列を返す
func SelectIds(db *sql.DB) ([]string, error) {

	rows, err := db.Query("select id from restaurant")
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	var res data.Restaurant
	var idArray []string
	for rows.Next() {
		err := rows.Scan(&res.Id)
		if err != nil {
			fmt.Println(err)
		}
		idArray = append(idArray, res.Id)
	}

	return idArray, nil

}

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
		r, e := ins.Exec(d.Id, d.Name, "", "", "", d.Geometry.Location.Longitude, d.Geometry.Location.Latitude, "kinki")
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(r)
	}
	tx.Commit()
}

func Update(id string, address string, db *sql.DB) error {

	upd, err := db.Prepare("update cocomeshi.restaurant set address = ? where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	r, e := upd.Exec(address, id)
	if e != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
