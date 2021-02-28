## build

`env GOOS=linux GOARCH=arm GOARM=6 go build`

## Docker

`$ docker buildx build -t greenhouse-manager . --platform linux/arm/v7`
`$ docker tag greenhouse-manager:latest mcpomm/greenhouse-manager:<tag>`
`$ docker push mcpomm/greenhouse-manager:<tag>`

## thingspeak api secret

`$ k create secret generic thing-speak-secret --from-literal THING_SPEAK_API_KEY=xxxxx`
