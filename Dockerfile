FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM scratch

WORKDIR /

COPY --from=build /app .

EXPOSE 8080

CMD ["/main"]