

FROM golang:alpine AS builder

#  Build 
ENV TZ=Asia/Ho_Chi_Minh
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
RUN mkdir -p /api
WORKDIR /api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./build/app
COPY .env.production ./build/.env
COPY serviceAccountKey.json ./build

# Run
FROM alpine:latest
RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/build .
EXPOSE 3000
ENTRYPOINT ["./app"]