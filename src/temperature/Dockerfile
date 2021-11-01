FROM arm32v7/golang:latest

WORKDIR /usr/src/app

COPY temperature ./

RUN chmod +x ./temperature
CMD ["./temperature"]