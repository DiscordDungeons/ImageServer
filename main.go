package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

	scanner := bufio.NewScanner(os.Stdin)
	code := ""

	for scanner.Scan() {
		code += scanner.Text()
	}

	iqlScanner := iql.NewScanner(code)

	tokens := iqlScanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}

	parser := iql.NewParser(tokens)

	statements := parser.Parse()

	//astPrinter := iql.NewASTPrinter()

	for _, statement := range statements {
		fmt.Printf("Statement: %s\n", statement)
	}

	// if err := scanner.Err(); err != nil {
	// 	log.Println(err)
	// }

	// lexer := iql.

	// v := lexer.Lex(iql.yySymType)

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println(v)
	// }

	// lexer := iql.NewLexer(os.Stdin)

	// iql.yyParse(lexer)

	//yyParse(NewLexer(os.Stdin))

	// v, err := iql.Parse([]byte("LOAD IMAGE FROM URL https://res.discorddungeons.me/images/achievements/killimanjaro/1.png AS ach_image;"))

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println(v)
	// }

	// http.HandleFunc("/", handler)

	// go func() {
	// 	for {
	// 		time.Sleep(time.Second)

	// 		log.Println("[ImageServer] Checking if server's started")

	// 		resp, err := http.Get(fmt.Sprintf("http://localhost%s", serverPort))

	// 		if err != nil {
	// 			log.Println("Failed:", err)
	// 			continue
	// 		}
	// 		resp.Body.Close()
	// 		if resp.StatusCode != http.StatusOK {
	// 			log.Println("Not OK:", resp.StatusCode)
	// 			continue
	// 		}

	// 		// Reached this point: server is up and running.
	// 		break
	// 	}

	// 	log.Printf("[ImageServer] Listening on port %s", serverPort)
	// }()

	// log.Println("[ImageServer] Starting server...")
	// log.Fatal(http.ListenAndServe(serverPort, nil))
}
