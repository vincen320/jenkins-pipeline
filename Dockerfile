FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o user_service

RUN chmod +x /app/user_service

#build a tiny docker image
FROM alpine:3

RUN mkdir /app

COPY --from=builder /app/user_service /app

COPY --from=builder /app/.env /app

CMD [ "/app/user_service" ]