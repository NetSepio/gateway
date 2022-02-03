FROM golang as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o gateway .


FROM alpine
WORKDIR /app
COPY --from=builder /app/gateway .
RUN mkdir logs
CMD [ "./gateway" ]