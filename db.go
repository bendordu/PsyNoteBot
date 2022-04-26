package main

import (
	"database/sql"
	"strconv"

	"log"

	"fmt"

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

func InsertUser(userbot UserBot, db *sql.DB) {
	db.Exec(fmt.Sprintf(`INSERT INTO %s (chat_id) VALUES (?);`, "user"), userbot.ChatID)
}

func UpdateData(userbot UserBot, db *sql.DB) {
	db.Exec(fmt.Sprintf(`UPDATE %s SET score=?, number=?, level=?, test_id=? WHERE chat_id=%d;`,
		"user", userbot.ChatID), userbot.Score, userbot.Number, userbot.Level, userbot.TestID)
}

func InsertResult(tresult Tresult, db *sql.DB) {
	db.Exec(fmt.Sprintf(`INSERT INTO %s (result, date, test_id, user_id) VALUES (?, ?, ?, ?);`,
		"t_user"), tresult.Result, tresult.Date, tresult.TestID, tresult.UserID)
}

func InsertTestID(userbot UserBot, nameEng string, db *sql.DB) {
	db.Exec(fmt.Sprintf(`UPDATE %s SET test_id=(SELECT test_id FROM %s WHERE name="%s") WHERE chat_id=%d;`,
		"user", "test", nameEng, userbot.ChatID))
}

func UpdateLevel(userbot UserBot, db *sql.DB) (sql.Result, error) {
	return db.Exec(fmt.Sprintf(`UPDATE %s SET level=? WHERE chat_id=%d;`,
		"user", userbot.ChatID), userbot.Level)
}

func UpdateNumber(userbot UserBot, db *sql.DB) (sql.Result, error) {
	return db.Exec(fmt.Sprintf(`UPDATE %s SET number=? WHERE chat_id=%d;`,
		"user", userbot.ChatID), userbot.Number)
}

func UpdateScore(userbot UserBot, db *sql.DB) (sql.Result, error) {
	return db.Exec(fmt.Sprintf(`UPDATE %s SET score=? WHERE chat_id=%d;`,
		"user", userbot.ChatID), userbot.Score)
}

func Select(smth string, userbot UserBot, db *sql.DB) (unknown int) {
	rows, err := db.Query(fmt.Sprintf(`SELECT %s FROM %s WHERE chat_id=%d;`,
		smth, "user", userbot.ChatID))
	if err != nil {
		log.Fatal(err)
	}
	var un string
	for rows.Next() {
		if err := rows.Scan(&un); err != nil {
			log.Fatal(err)
		}
	}
	defer rows.Close()
	unknown, _ = strconv.Atoi(un)
	return unknown
}
