module discorddungeons.me/imageserver/main

go 1.16

replace discorddungeons.me/imageserver/iql => ./internal/iql

replace discorddungeons.me/imageserver/iql/syntax => ./internal/iql/internal/syntax

replace discorddungeons.me/imageserver/iql/schema => ./internal/iql/internal/schema

require (
	discorddungeons.me/imageserver/iql v0.0.0
	github.com/joho/godotenv v1.3.0
)
