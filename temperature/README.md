# Greenhouse temperature module

## build

`env GOOS=linux GOARCH=arm GOARM=6 go build`

## upload to greenhouse-test

`rsync -avz temperature pi@greenhouse-test:greenhouse/temperature`

## wire

![ds18b20-schema.png](ds18b20-schema.png)
