# Build Stage
FROM docker.io/library/golang:0.21.5 as builder
COPY . /app
WORKDIR /app
ENV CGO_ENABLED=0
RUN go build -o dbserver cmd/dbserver.go

# Run Stage
FROM gcr.io/distroless/static-debian12:nonroot
EXPOSE 8080
COPY --from=builder /app/dbserver /app/dbserver
ENTRYPOINT ["/app/dbserver"]
