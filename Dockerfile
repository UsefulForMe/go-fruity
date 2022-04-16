

FROM golang:alpine AS builder

#  Build 
ENV TZ=Asia/Ho_Chi_Minh
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
RUN mkdir -p /go-ecommerce
WORKDIR /go-ecommerce
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./build/app
COPY .env.production ./build/.env
COPY serviceAccountKey.json ./build

# Run
FROM alpine:latest
RUN mkdir -p /go-ecommerce
WORKDIR /go-ecommerce
COPY --from=builder /go-ecommerce/build .
EXPOSE 3000
ENTRYPOINT ["./app"]