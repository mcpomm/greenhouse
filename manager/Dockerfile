FROM arm32v7/golang:latest AS builder
WORKDIR /usr/src/build
RUN go get -d -v golang.org/x/net/html  
COPY . .
RUN GOOS=linux GOARCH=arm GOARM=6 go build  .


FROM arm32v7/golang:latest
WORKDIR /usr/src/app
COPY config config
COPY --from=builder /usr/src/build/manager .
CMD ["./manager"]