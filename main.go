package main

import (
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"discorddungeons.me/imageserver/cache"
	"discorddungeons.me/imageserver/iql"
	"github.com/joho/godotenv"
)

const VERSION = "1.0.1"

type CacheConfig struct {
	ENABLE_CACHE    bool
	CACHE_DIRECTORY string
}

var cacheConfig CacheConfig

var cacheInstance *cache.Cache

// Sends the data as a JSON string to a http response.
func sendJSON(w http.ResponseWriter, data map[string]interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	if _, ok := data["statusCode"]; !ok {
		// No status code in the data
		data["statusCode"] = httpStatus
	}

	resp, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprint(w, string(resp))
}

// Sends an error to the responseWriter w, with the httpStatus, and an optional message
func sendError(w http.ResponseWriter, httpStatus int, message string) {
	data := make(map[string]interface{})

	data["httpError"] = http.StatusText(httpStatus)

	if len(message) != 0 {
		data["error"] = message
	}

	sendJSON(w, data, httpStatus)
}

// Gets an environment variable by key, or fallbacks to the fallback if it's not defined.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Handles requests to the /status endpoint
func statusHandler(w http.ResponseWriter, req *http.Request) {
	sendJSON(w, make(map[string]interface{}), 200)
}

// Handles requests to the / endpoint.
func handler(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" && req.Method != "POST" {
		sendError(w, http.StatusMethodNotAllowed, "")
		return
	}

	if req.Method == "GET" {
		data := make(map[string]interface{})

		data["version"] = VERSION

		sendJSON(w, data, http.StatusOK)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Can't read body")
		return
	}

	if cacheConfig.ENABLE_CACHE {
		hash := cacheInstance.ComputeHash(body)

		if cacheInstance.HasFile(hash + ".png") {
			http.ServeFile(w, req, cacheConfig.CACHE_DIRECTORY+"/"+hash+".png")

			return
		}
	}

	runner := iql.NewIQLRunner()

	res, err := runner.RunIQL(string(body))

	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	i := 0

	for _, img := range res {
		if i > 0 {
			continue
		}

		w.Header().Set("Content-Type", "image/png")

		err := png.Encode(w, img)

		if err != nil {
			sendError(w, http.StatusBadRequest, "Can't return image")

			return
		}

		if cacheConfig.ENABLE_CACHE {
			hash := cacheInstance.ComputeHash(body)

			err := cacheInstance.SavePngFile(hash+".png", img)

			if err != nil {
				fmt.Println("Can't save file to cache: " + err.Error())
			}
		}

		i++
	}

}

// Executes the program
func main() {
	godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	enableCache := true

	c, err := strconv.ParseBool(getEnv("ENABLE_CACHE", "true"))

	if err == nil {
		enableCache = c
	}

	cacheConfig = CacheConfig{
		ENABLE_CACHE:    enableCache,
		CACHE_DIRECTORY: getEnv("CACHE_DIRECTORY", "cache"),
	}

	cacheInstance = cache.NewCache(cacheConfig.CACHE_DIRECTORY)

	serverPort := getEnv("SERVER_PORT", "8080")

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
