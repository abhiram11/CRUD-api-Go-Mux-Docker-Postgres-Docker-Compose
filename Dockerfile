FROM golang:1.16.3-alpine3.13

WORKDIR /app

# copy source code: COPY <---> FROM <---> TO
COPY . .

# Download and install dependencies into the Docker Image
RUN go get -d -v ./...

# build, where "api" = module "name" in go.mod file
RUN go build -o api .

# as written in listenAndServe
EXPOSE 8000

CMD ["./api"]