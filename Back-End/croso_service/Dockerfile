FROM golang:latest as builder
WORKDIR /app
COPY ./croso_service/ .
COPY ./authorization/ /authorization/
COPY ./model/ /model/
COPY ./nats/ /nats/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]