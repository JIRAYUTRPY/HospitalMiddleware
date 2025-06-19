FROM golang:1.24.4-alpine3.22 as dev
WORKDIR /app
ENV GIN_MODE=debug
ENV TZ Asia/Bangkok
RUN apk update && \
apk add --no-cache tzdata
RUN go install github.com/cosmtrek/air@v1.27.3
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download && go mod verify
CMD ["air", "-c", ".air.toml"]
EXPOSE 8080

FROM nginx:stable-alpine as nginx_revproxy
COPY ./nginx/default.conf /etc/nginx/conf.d/
EXPOSE 80/tcp
CMD ["/bin/sh", "-c", "exec nginx -g 'daemon off;';"]
WORKDIR /usr/share/nginx/html


FROM golang:1.24.4 as build
WORKDIR /app
ENV TZ Asia/Bangkok
ENV GIN_MODE=release
RUN apk update && \
apk add --no-cache tzdata
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download && go mod verify
COPY ./src/. .
RUN go build -o . ./src/main.go


FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=build --link /app/main ./

ENTRYPOINT ["/app/main"]