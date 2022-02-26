FROM golang:alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN apk add build-base
RUN go mod download
COPY . .
RUN go build -o gateway .

FROM alpine
WORKDIR /app
COPY --from=builder /app/gateway .
CMD [ "./gateway" ]