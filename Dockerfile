FROM golang:1.20-alpine AS builder
ENV PROJECT_PATH=/app
WORKDIR $PROJECT_PATH
COPY . .
RUN go build -o main .
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:latest
ENV PROJECT_PATH=/app
WORKDIR $PROJECT_PATH
COPY --from=builder $PROJECT_PATH/main .
COPY --from=builder $PROJECT_PATH/app.env .
COPY --from=builder $PROJECT_PATH/migrate ./migrate
COPY --from=builder $PROJECT_PATH/run.sh .
COPY --from=builder $PROJECT_PATH/wait-for.sh .
COPY db/migrations ./db/migrations
RUN chmod +x ./migrate

EXPOSE 8080
CMD ["./main"]
ENTRYPOINT ["./run.sh"]