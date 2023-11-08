FROM golang:1.21-alpine as builder

RUN apk update && apk add git

WORKDIR /usr/src/app

COPY . .
RUN go get -v ./...

RUN go build -o /app

FROM alpine:latest
WORKDIR /home
COPY --from=builder /app ./app

ENV GOOGLE_OAUTH_CLIENT_ID = ""
ENV GOOGLE_OAUTH_CLIENT_SECRET = ""
ENV REDIRECT_URL = ""
ENV MONGO_URL = ""
ENV JWT_SECRET = ""
ENV TOKEN_EXPIRED_IN=60m
ENV TOKEN_MAXAGE=60

EXPOSE 8080

CMD ["./app"]

