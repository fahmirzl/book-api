package db

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*.sql
var dbMigrations embed.FS


var (
	DB *sql.DB
	err error
)


func DBConnection() {
   err = godotenv.Load(".env")
    if err != nil {
       panic("Error loading .env file")
    }

    psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
       os.Getenv("PGHOST"),
       os.Getenv("PGPORT"),
       os.Getenv("PGUSER"),
       os.Getenv("PGPASSWORD"),
       os.Getenv("PGDATABASE"),
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