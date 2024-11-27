package svr

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type Svr struct {
	Router *chi.Mux
}

func StartSvr() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("PORT string is empty. Using Default")
		PORT = "8080"
	}

	s := &Svr{
		Router: chi.NewRouter(),
	}

	s.MountHdlrs()

	log.Printf("Starting svr on :%s", PORT)
	if err := http.ListenAndServe(":"+PORT, s.Router); err != nil {
		log.Fatal(err)
	}
}
