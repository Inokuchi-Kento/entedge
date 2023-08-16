package infrastructure

import (
	"entedge/ent"
	"log"

	_ "github.com/lib/pq"
)

func OpenDB() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5544 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return client
}
