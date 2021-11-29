FROM golang:1-alpine as build

WORKDIR /app
RUN go build -o /live-puppet

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/live-puppet /app/live-puppet

EXPOSE 8180
ENTRYPOINT ["./live-puppet"]