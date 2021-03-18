module discorddungeons.me/imageserver/main

go 1.16

replace discorddungeons.me/imageserver/iql => ./internal/iql

replace discorddungeons.me/imageserver/iql/iqlTypes => ./internal/iql/internal/iqlTypes

replace discorddungeons.me/imageserver/iql/modifyProperty => ./internal/iql/internal/modifyProperty

require (
	discorddungeons.me/imageserver/iql v0.0.0
	discorddungeons.me/imageserver/iql/iqlTypes v0.0.0
	discorddungeons.me/imageserver/iql/modifyProperty v0.0.0
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/joho/godotenv v1.3.0
)
