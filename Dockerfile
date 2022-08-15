FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o user_service

RUN chmod +x /app/user_service

#build a tiny docker image
FROM alpine:3

RUN mkdir /app

#SANGAT PENGARUH!
#AGAR PATH SEKARANG MENJADI SAMA SEPERTI DI BINARY, KARENA GOLANG MULAI MEMBACA PATH SESUAI DENGAN PATH SEKARANG (JIKA ADA PROGRAM GOLANG YANG BUTUH BACA FILE INI SANGAT PERLU, CONTOH: godotenv())
#CONTOH COBA LIHAT DI TERMINAL PASTI POSISI NYA DI /d/test Microservices/Application 1/user-service
#COBA di cd ..
#LALU run main.go nya seperti ini -> user-service/main.go
#PASTI ERROR karena golangnya tidak ketemu file go.mod di path yang sekarang (/d/test Microservices/Application 1) (kalo go run sepertinya masih butuh go mod)
WORKDIR /app

COPY --from=builder /app/user_service /app
COPY --from=builder /app/.env /app

ENTRYPOINT [ "./user_service" ]