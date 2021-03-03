package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	conn, err := sql.Open("mysql", "test:dlatldyd1!@tcp(192.168.41.125:3306)/test")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := conn.Query("SELECT b_id, title, content FROM board")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var b_id int
	var title string
	var content string

	for rows.Next() {

		rows.Scan(&b_id, &title, &content)
		fmt.Printf("id : %d,  title : %s, content : %s\n", b_id, title, content)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
		fmt.Printf("id : %d,  title : %s, content : %s\n", b_id, title, content)
	})

	http.ListenAndServe(":8080", nil)

	//DB 접속 해제
	conn.Close()

}
