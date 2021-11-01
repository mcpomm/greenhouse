FROM arm32v7/golang:latest

WORKDIR /usr/src/app

COPY soil-temperature ./

RUN chmod +x ./soil-temperature
CMD ["./soil-temperature"]