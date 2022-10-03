package alert

import (
	"database/sql"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

type Price struct {
	Time  time.Time `json:"time"`
	Price float64   `json:"price"`
	Coin  string    `json:"coin"`
}

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Print("connected")
	DB = db
}

func DBinsert(price float64) {
	sqlStatement := `INSERT INTO price VALUES (CURRENT_TIME(), ?, "btc")`
	_, err := DB.Exec(sqlStatement, price)
	if err != nil {
		log.Fatal("Insertion error: ", err)
	}
	log.Print("Inserted ")

}

func GetPrice(date string, limit int, offset int) []*Price {
	bDate := date
	aDate := strings.Replace(date, "00:00:00", "23:59:59", 1)
	// Execute the query
	results, err := DB.Query("SELECT * FROM price WHERE time BETWEEN ? AND ? LIMIT ?, ?;", bDate, aDate, offset, limit)
	if err != nil {
		log.Fatal("Error in DB while do a select: ", err)
	}
	var prices []*Price
	for results.Next() {
		var p Price
		// for each row, scan the result into our tag composite object
		err = results.Scan(&p.Time, &p.Price, &p.Coin)
		if err != nil {
			log.Fatal("Error inserting into variable: ", err)
		}
		prices = append(prices, &p)
	}
	return prices
}
