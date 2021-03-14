package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"discorddungeons.me/imageserver/iql"
	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")

	if !strings.HasPrefix(serverPort, ":") {
		serverPort = fmt.Sprintf(":%s", serverPort)
	}

	v, err := iql.Parse([]byte("4085551212"))

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(v)
	}

	http.HandleFunc("/", handler)

	go func() {
		for {
			time.Sleep(time.Second)

			log.Println("[ImageServer] Checking if server's started")

			resp, err := http.Get(fmt.Sprintf("http://localhost%s", serverPort))

			if err != nil {
				log.Println("Failed:", err)
				continue
			}
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				log.Println("Not OK:", resp.StatusCode)
				continue
			}

			// Reached this point: server is up and running.
			break
		}

		log.Printf("[ImageServer] Listening on port %s", serverPort)
	}()

	log.Println("[ImageServer] Starting server...")
	log.Fatal(http.ListenAndServe(serverPort, nil))
}
