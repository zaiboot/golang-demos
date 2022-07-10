FROM golang:latest as build

ENV GOOS=linux

RUN mkdir -p /code/out/
WORKDIR /code/
COPY . /code/
RUN go mod tidy
RUN go build -tags musl -o out/consome

FROM alpine AS package

WORKDIR /app/
COPY --from=build /code/out/consome /app/consome
CMD ["/app/consome"]

