package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"embed"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*.sql
var dbMigrations embed.FS


var (
	DB *sql.DB
	err error
)


func DBConnection() {
	psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
       "localhost", 5432, "postgres", "root", "book-api",
    )

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

    err = DB.Ping()
    if err != nil {
       panic(err)
    }

    fmt.Println("Successfully connected!")
}

func DBMigrate() {
    migrations := &migrate.EmbedFileSystemMigrationSource{
       FileSystem: dbMigrations,
       Root:       "migrations",
    }

    n, errs := migrate.Exec(DB, "postgres", migrations, migrate.Up)
    if errs != nil {
       panic(errs)
    }

    fmt.Println("Migration success, applied", n, "migrations!")
}

func Init() {
	DBConnection()
	DBMigrate()
}