package main

import (
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

type CacheConfig struct {
	ENABLE_CACHE    bool
	CACHE_DIRECTORY string
}

var cacheConfig CacheConfig

var cacheInstance *cache.Cache

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
		http.Error(w, err.Error(), 400)
	}

	i := 0

	for _, img := range res {
		if i > 0 {
			continue
		}

		w.Header().Set("Content-Type", "image/png")

		err := png.Encode(w, img)

		if err != nil {
			http.Error(w, "can't return image", http.StatusBadRequest)
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

func main() {
	godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	enableCache := true

	c, err := strconv.ParseBool(os.Getenv("ENABLE_CACHE"))

	if err == nil {
		enableCache = c
	}

	cacheConfig = CacheConfig{
		ENABLE_CACHE:    enableCache,
		CACHE_DIRECTORY: os.Getenv("CACHE_DIRECTORY"),
	}

	cacheInstance = cache.NewCache(cacheConfig.CACHE_DIRECTORY)

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
