package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kaungmyathan22/golang-graphql-social-network/config"
	"github.com/kaungmyathan22/golang-graphql-social-network/database"
)

func main() {
	ctx := context.Background()
	config.LoadEnv(".env")
	conf := config.New()
	db := database.New(ctx, conf)
	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running....")
}
