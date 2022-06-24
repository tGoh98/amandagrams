FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./amandagrams .


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/amandagrams .
EXPOSE 8000
ENTRYPOINT ["./amandagrams"]