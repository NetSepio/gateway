FROM golang:alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN apk add build-base
RUN go mod download
COPY . .
RUN apk add --no-cache git && go build -o gateway . && apk del git

FROM alpine
WORKDIR /app
COPY --from=builder /app/gateway .
RUN apk add --no-cache chromium
CMD [ "./gateway" ]
