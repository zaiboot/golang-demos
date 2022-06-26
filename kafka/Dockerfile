FROM golang:latest as build

ENV CGO_ENABLED=0:
ENV GOOS=linux

RUN mkdir -p /code/out/
WORKDIR /code/
COPY . /code/
RUN go mod tidy
RUN go build -ldflags='-extldflags "-static"' -o out/consome

FROM alpine AS package

WORKDIR /app/

COPY --from=build /code/out/consome /app/consome
CMD ["/app/consome"]

