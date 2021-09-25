package models

import (
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/crypto/sha3" // https://pkg.go.dev/golang.org/x/crypto/sha3
)

var (
	DB  *sql.DB
	err error
)

/*
テーブル名の一蘭
*/
const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	// SQLドライバーとDBNameを渡してDBコネクションを取得
	DB, err = sql.Open(config.Config.SQLDriver, config.Config.DBName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	DB.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	DB.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	DB.Exec(cmdS)
}

/*
createUUID は UUID を作成します。
*/
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

/*
Encrypt は パスワードをハッシュ値で暗号化します。
*/
func Encrypt(plainText string) (crypText string) {
	crypText = fmt.Sprintf("%x", sha3.Sum256([]byte(plainText)))
	return crypText
}
