package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("Не получилось подключиться:\n" + err.Error())
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Println("Не получилось пропинговать:\n" + err.Error())
		return nil, err
	}
	return db, nil
}

func existDB(uid string, db *sql.DB) (bool, error) {
	var count int
	//row := db.QueryRow("SELECT EXISTS (SELECT * FROM json_table WHERE uid = $1);", uid)
	row := db.QueryRow("SELECT COUNT (DISTINCT uid) FROM json_table WHERE uid = $1;", uid)
	err := row.Scan(&count)
	if err != nil {
		log.Println("exist check error", err)
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func selectJSON(uid string, db *sql.DB) (error, string) {
	var res string
	rows, err := db.Query("select information from json_table where uid = $1", uid)
	if err != nil {
		log.Println("select error", err)
		return err, "1"
	}
	defer rows.Close()
	result := make([]interface{}, 0)
	for rows.Next() {
		err = rows.Scan(&res)
		if err != nil {
			log.Println("не получилось просканить поле с json", err)
		}
		result = append(result, res)
	}
	//resByte := []byte(res)
	return err, res
}

func insertDB(str string, uid string, db *sql.DB) error {
	if check, _ := existDB(uid, db); !check {
		stmt := `INSERT INTO json_table (uid, information) VALUES ($1, $2);`
		if _, err := db.Exec(stmt, uid, str); err != nil {
			log.Println("insert info error", err)
			return err
		}
		log.Println("Add json with uid: ", str)
	} else {
		log.Println("elem exist(insert is impossible)")
	}
	return nil
}

func mapDB(db *sql.DB) map[string]string {
	var chtot1, chtot2 string
	m := make(map[string]string)
	rows, err := db.Query("SELECT * FROM json_table;")
	//row := db.QueryRow(`SELECT * FROM json_table;`)
	//err := row.Scan(&chtot1, &chtot2)
	for rows.Next() {
		if err = rows.Scan(&chtot1, &chtot2); err != nil {
			log.Println("select map error", err)
			return nil
		}
		m[chtot1] = chtot2
	}

	//fmt.Println(chtot1, chtot2)
	//m := make(map[string]string)
	//m[chtot1] = chtot2
	return m
}
