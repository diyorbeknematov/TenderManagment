FROM golang:1.23.2 AS builder

WORKDIR /user-app

COPY . ./
RUN go mod download

COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../myapp

FROM alpine:latest

WORKDIR /user-app

COPY --from=builder /user-app/myapp .
COPY --from=builder /user-app/pkg/logs/app.log ./pkg/logs/
COPY --from=builder /user-app/.env .

EXPOSE 8080

CMD [ "./myapp" ]