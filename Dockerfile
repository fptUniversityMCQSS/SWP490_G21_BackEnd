FROM golang:1.15
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/dist .
EXPOSE 80
EXPOSE 443
CMD ./out/dist