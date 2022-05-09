package main

import (
	"database/sql"
	"strconv"

	"fmt"

	"log"

	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func InsertUser(chatID int64, db *sql.DB) {
	db.Exec(fmt.Sprintf(`INSERT INTO %s (chat_id) VALUES (?);`, "user"), chatID)
}

func InsertResult(tresult Tresult, db *sql.DB) {
	db.Exec(fmt.Sprintf(`INSERT INTO %s (result, date, test_id, chat_id) VALUES (?, ?, ?, ?);`,
		"test_result"), tresult.Result, tresult.Date, tresult.TestID, tresult.ChatID)
}

func SelectTestID(nameEng string, db *sql.DB) (testID int) {
	rows, err := db.Query(fmt.Sprintf(`SELECT test_id FROM %s WHERE name="%s";`, "test", nameEng))
	if err != nil {
		log.Fatal(err)
	}
	var tID string
	for rows.Next() {
		if err := rows.Scan(&tID); err != nil {
			log.Fatal(err)
		}
	}
	defer rows.Close()
	testID, _ = strconv.Atoi(tID)
	return testID
}

func SelectOldResult(chatID int64, typesTest TypesTest, db *sql.DB) (str string) {
	rows, err := db.Query(fmt.Sprintf(`SELECT name, date, result FROM %s
										JOIN test USING (test_id)
									WHERE chat_id=%d
									ORDER BY name, date;`, "test_result", chatID))
	if err != nil {
		log.Fatal(err)
	}

	var (
		name, result string
		date         time.Time
	)

	str = "📅<b>Ваши результаты:</b>\n\n"

	for rows.Next() {
		if err := rows.Scan(&name, &date, &result); err != nil {
			log.Fatal(err)
		}
		for _, typeTest := range typesTest.TestTypes {
			for _, test := range typeTest.Tests {
				if test.TestName == name {
					name = test.TestNameRus
				}
			}
		}
		str += fmt.Sprintf("Шкала %s [%s] - %s\n", name, date.Format("02.01.2006 15:04"), result)
	}
	defer rows.Close()
	return str
}

func SelectTests(typesTest TypesTest, db *sql.DB) (namesRus []string) {
	rows, err := db.Query(fmt.Sprintf(`SELECT name FROM test ORDER BY name;`))
	if err != nil {
		log.Fatal(err)
	}

	var name string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		for _, typeTest := range typesTest.TestTypes {
			for _, test := range typeTest.Tests {
				if test.TestName == name {
					namesRus = append(namesRus, test.TestNameRus)
				}
			}
		}
	}
	defer rows.Close()
	return namesRus
}

/*
CREATE TABLE test_result(
    t_result_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    date DATE,
	result INT,
    test_id INT,
    chat_id INT64,
    FOREIGN KEY (chat_id) REFERENCES user (chat_id)
)
CREATE TABLE user(
    user_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    chat_id INT64 NOT NULL UNIQUE
)
*/
