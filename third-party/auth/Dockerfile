FROM golang:alpine AS build

# Копируем исходный код в Docker-контейнер
WORKDIR /server
COPY . .

# Билдим бинарь
RUN go build -mod=vendor cmd/auth/main.go

# Копируем на чистый образ
FROM alpine

COPY --from=build /server/main /main

CMD ./main