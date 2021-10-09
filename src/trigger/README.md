## Docker

`$ docker buildx build -t greenhouse-trigger . --platform linux/arm/v7`
`$ docker tag greenhouse-trigger:latest mcpomm/greenhouse-trigger:<tag>`
`$ docker push mcpomm/greenhouse-trigger:<tag>`
