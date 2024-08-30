package mysqls

import (
	"database/sql"
	// "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

// defer 延迟执行
func G_mysql() {
	db, err := sql.Open("mysql", "root:199011081108@/go")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO song (name, singer, composer, lyricist) VALUES('青花瓷', '周杰伦', '周杰伦', '方文山')")
	if err != nil {
		panic(err.Error())
	}
	stmtIns.Exec()
	defer stmtIns.Close()

	stmtOut, err := db.Prepare("SELECT * FROM song ")
	if err != nil {
		panic(err.Error())
	}
	rows, err := stmtOut.Query()
	var id int
	var name string
	var singer string
	var lyricist string
	var composer string
	for rows.Next() {
		if err := rows.Scan(&id, &name, &singer, &lyricist, &composer); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id:%d, name:%s, singer:%s, lysicist:%s, composer:%s \n", id, name, singer, lyricist, composer)
	}
	defer stmtOut.Close()
}
