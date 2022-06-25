FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY go.mod go.sum main.go preprocessing.go .
COPY ./data/scrabbleWords.txt ./data/
RUN go mod download
RUN go build -o ./amandagrams .
RUN ./amandagrams -g


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/amandagrams .
COPY --from=builder /app/data/lettersToWords* ./data/
EXPOSE 8000
ENTRYPOINT ["./amandagrams"]