FROM golang:1.17.8-alpine

RUN mkdir /restAss

WORKDIR /restAss

# ADD ./ /restAss

COPY . .

RUN go mod download

RUN go build -o main

EXPOSE 3000

CMD ["/restAss/main"]                             