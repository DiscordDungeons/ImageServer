# ImageServer
Small microservice for image generation.

# License

ImageServer is licensed under the EUPL-1.2-or-later.

# Environment Variables

| Property        | Type    | Description                                     | Default value |
|-----------------|---------|-------------------------------------------------|---------------|
| SERVER_PORT     | number  | The server port to listen on                    | 8080          |
| ENABLE_CACHE    | boolean | If the server should cache generated images     | true          |
| CACHE_DIRECTORY | string  | The directory cached images should be stored in | cache         |

# Docker image

ImageServer also comes as a Docker image, [`discorddungeons/imageserver`](https://hub.docker.com/repository/docker/discorddungeons/imageserver).

It is recommended that you attach a permanent volume mounted to your cache folder when running the docker image.

Run with

`docker run -p 8080:8080 --env-file .env discorddungeons/imageserver`
