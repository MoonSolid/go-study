package main

import (
	// 	"database/sql"
	"fmt"
	// 	"os"
	// 	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Print("hello world")
}

// func main() {
// 	conn, err := sql.Open("mysql", "dba:dbrhksflwk99((@tcp(192.168.41.125:3306)/base")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	rows, err := conn.Query("SELECT city_code, city_name FROM city")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	for rows.Next() {
// 		var city_code string
// 		var city_name string
// 		rows.Scan(&city_code, &city_name)
// 		fmt.Printf("city_code : %s, city_name : %s\n", city_code, city_name)
// 	}

// 	//DB 접속 해제
// 	conn.Close()
// }
