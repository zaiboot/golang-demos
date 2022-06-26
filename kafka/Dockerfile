FROM golang:latest

ENV CGO_ENABLED=0:
ENV GOOS=linux

WORKDIR /code/
COPY . /code/
RUN go mod tidy
RUN go build -o consome
RUN ls /code/ & sleep 10
CMD ["/code/consome"]

