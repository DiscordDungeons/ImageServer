# ImageServer
Small microservice for image generation.

# License

ImageServer is licensed under the EUPL-1.2-or-later.

# Environment Variables

| Property    | Type   | Description                  |
|-------------|--------|------------------------------|
| SERVER_PORT | number | The server port to listen on |

# Docker image

Run with

`docker run -p 8080:8080 --env-file .env imgserver`
