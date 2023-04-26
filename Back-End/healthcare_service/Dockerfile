FROM golang:latest as builder
WORKDIR /app
COPY ./healthcare_service/go.mod ./healthcare_service/go.sum ./
RUN go mod download
COPY ./healthcare_service/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY /rbac/healthcare_model.conf/ .
COPY /healthcare_service/policy.csv .
EXPOSE 8000
CMD ["./main"]