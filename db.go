package main

import (
	"database/sql"
)

func main() {
	db, err := sql.Open("mysql", "root:Test@/deneme")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//query := "CREATE TABLE List (id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,firstname VARCHAR(30) NOT NULL,lastname VARCHAR(30) NOT NULL)"
	//query := "INSERT INTO List(id,firstname,lastname) VALUES(1,'Celal','ORMANLI')"
	query := "DELETE FROM List WHERE id=2"
	db.Exec(query)

}
