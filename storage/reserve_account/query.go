package reserve_account

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type ReverseAcc struct {
	Id      int
	Service int
	OrderId int
	Cost    int
}

func ReverseAccInsert(db *sql.DB, ra ReverseAcc) error {
	_, err := db.Exec(
		"INSERT INTO reserve_account (id, service, order_id, cost) VALUES ($1, $2, $3, $4 )",
		ra.Id, ra.Service, ra.OrderId, ra.Cost)
	if err != nil {
		log.Println("insert problem", err)
		return err
	}
	return nil
}

func ReverseAccSelect(db *sql.DB) error {
	rows, err := db.Query("SELECT * from reserve_account")
	if err != nil {
		log.Println("zapros err", err)
	}
	var items []ReverseAcc
	for rows.Next() {
		ra := ReverseAcc{}
		err = rows.Scan(&ra.Id, &ra.Service, &ra.OrderId, &ra.Cost)
		if err != nil {
			log.Println("scan problem", err)
			return err
		}
		items = append(items, ra)
	}
	err = rows.Close()
	if err != nil {
		log.Println("close problem", err)
		return err
	}
	for _, val := range items {
		fmt.Println(val)
	}
	return nil
}