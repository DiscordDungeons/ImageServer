package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"discorddungeons.me/imageserver/iql"
	"github.com/joho/godotenv"
)

func statusHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK")
}

func handler(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	runner := iql.NewIQLRunner()

	res, err := runner.RunIQL(string(body))

	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	i := 0

	for _, img := range res {
		if i > 0 {
			continue
		}

		buf := new(bytes.Buffer)

		err := png.Encode(buf, img)

		if err != nil {
			http.Error(w, "can't return image", http.StatusBadRequest)
			return
		}

		data := buf.Bytes()

		w.Header().Set("Content-Type", "image/png")

		w.Write(data)

		i++
	}
}

func main() {
	godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	serverPort := os.Getenv("SERVER_PORT")

	if !strings.HasPrefix(serverPort, ":") {
		serverPort = fmt.Sprintf(":%s", serverPort)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/status", statusHandler)

	go func() {
		for {
			time.Sleep(time.Second)

			log.Println("[ImageServer] Checking if server's started")

			resp, err := http.Get(fmt.Sprintf("http://localhost%s/status", serverPort))

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
