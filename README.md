# ImageServer
Small microservice for image generation.

# Environment Variables

| Property    | Type   | Description                  |
|-------------|--------|------------------------------|
| SERVER_PORT | number | The server port to listen on |

# Docker image

Run with

`docker run -p 8080:8080 --env-file .env imgserver`