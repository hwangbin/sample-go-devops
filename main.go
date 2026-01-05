package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL 드라이버
)

func main() {
	// host=localhost가 아니라 host=my-postgres여야 합니다!
	connStr := "host=my-postgres port=5432 user=postgres password=pass123 dbname=postgres sslmode=disable"

	// 1. DB 연결
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. 테이블 생성 (PostgreSQL 문법)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS visitors (id SERIAL PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "손님"
		}

		// 3. DB 저장
		_, err := db.Exec("INSERT INTO visitors (name) VALUES ($1)", name)
		if err != nil {
			log.Println("저장 실패:", err)
			http.Error(w, "DB 저장 실패", 500)
			return
		}

		fmt.Fprintf(w, "안녕하세요 %s님! 도커 DB에 저장되었습니다.", name)
	})

	fmt.Println("서버 시작: http://localhost:1234/greet")
	http.ListenAndServe(":1234", nil)
}
