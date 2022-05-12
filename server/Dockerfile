FROM golang:1.17.8-alpine3.15

RUN mkdir server

WORKDIR /server

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN export GO111MODULE=on
RUN cd /server 
RUN go build -o main.exe

EXPOSE 9000

CMD [ "/server/main.exe" ]