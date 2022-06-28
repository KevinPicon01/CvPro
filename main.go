package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"kevinPicon/go/src/CvPro/handlers"
	Middleware "kevinPicon/go/src/CvPro/middleware"
	"kevinPicon/go/src/CvPro/server"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	serv, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})
	if err != nil {
		log.Fatal(err)
	}
	serv.Start(BindRouters)
}
func BindRouters(s server.Server, r *mux.Router) {
	//Middleware
	r.Use(Middleware.CheckAuthMiddleware(s))
	//Routes
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods("GET")
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods("POST")
	//r.HandleFunc("/signup/verifyUser", handlers.VerifyUserHandler(s)).Methods("POST")
}
