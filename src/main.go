package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"database/sql"
	_"github.com/lib/pq"

)

const(
	HOST = "localhost"
	DATABASE = "postgres"
	USER = "postgres"
	PASSWORD = "postgres"
)

type User struct {
	Id int
	Name string
}

func main(){
	http.HandleFunc("/",selectHandler)
	http.ListenAndServe(":8080",nil)

}

func selectHandler(w http.ResponseWriter,_ *http.Request){
	db, err := sql.Open("postgres", "postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	/*
	//テーブル削除
	_, err = db.Exec("DROP TABLE IF EXISTS users;")
	if err != nil {
		panic(err)
	}
	
	// テーブル作成
	_, err = db.Exec("CREATE TABLE users (user_id serial PRIMARY KEY, user_name VARCHAR(50));")
	if err != nil {
		panic(err)
	}

	// データ挿入
	_, err = db.Exec(`
		INSERT INTO
			users (user_name)
		VALUES
			('太郎'),
			('二郎'),
			('三郎')
	`)
	if err != nil {
		panic(err)
	}
	*/
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "ID:%v,NAME:%v\n",strconv.Itoa(user.Id), user.Name)
	}
}