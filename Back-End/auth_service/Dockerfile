FROM golang:latest as builder
WORKDIR /app
COPY ./auth_service/ .
COPY ./authorization/ /authorization/
RUN cd /authorization && go mod download
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY /rbac/auth_model.conf/ .
COPY /auth_service/policy.csv .
EXPOSE 8000
CMD ["./main"]