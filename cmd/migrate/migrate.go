package main

import (
	"borrow/internal/db"
	"borrow/internal/env"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	env.Init()

	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[1]

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "fix":
		version := -1
		dirty := false
		row := db.QueryRow("SELECT * FROM public.schema_migrations ORDER BY version DESC LIMIT 1")
		row.Scan(&version, &dirty)

		if version != -1 {
			if err := m.Force(version); err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		}
	case "rollback":
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "version":
		version := -1
		dirty := false
		row := db.QueryRow("SELECT * FROM public.schema_migrations ORDER BY version DESC LIMIT 1")
		row.Scan(&version, &dirty)

		if dirty {
			fmt.Printf("%d dirty", version)
		} else {
			fmt.Printf("%d clean", version)
		}
	}
}
