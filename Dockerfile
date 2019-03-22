FROM ubuntu:18.04

# Обвновление списка пакетов
RUN apt-get -y update
RUN apt install -y git wget gcc gnupg

# Установка golang
RUN wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz
RUN tar -xvf go1.11.linux-amd64.tar.gz
RUN mv go /usr/local

# Выставляем переменную окружения для сборки проекта
ENV GOROOT /usr/local/go
ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

# Копируем исходный код в Docker-контейнер
WORKDIR /server
COPY . .

# Объявлем порт сервера
EXPOSE 5051

RUN ls

RUN go build cmd/server/main.go
CMD ./main