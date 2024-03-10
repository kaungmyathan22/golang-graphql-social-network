package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kaungmyathan22/golang-graphql-social-network/config"
	"github.com/kaungmyathan22/golang-graphql-social-network/database"
	"github.com/kaungmyathan22/golang-graphql-social-network/graph"
)

func main() {
	ctx := context.Background()
	config.LoadEnv(".env")
	conf := config.New()
	db := database.New(ctx, conf)
	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Timeout(time.Second * 60))

	//router.Use(graph.DataloaderMiddleware(
	//	&graph.Repos{
	//		UserRepo: userRepo,
	//	},
	//))

	//router.Use(authMiddleware(authTokenService))
	router.Handle("/", playground.Handler("Twitter clone", "/query"))
	router.Handle("/query", handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					//AuthService:  authService,
					//TweetService: tweetService,
					//UserService:  userService,
				},
			},
		),
	))

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))

}
