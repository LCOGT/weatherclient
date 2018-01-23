FROM golang:1.9-alpine
MAINTAINER Austin Riba <ariba@lco.global>

WORKDIR /go/src/weather
EXPOSE 80

COPY . .
RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
