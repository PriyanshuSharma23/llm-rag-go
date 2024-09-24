FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o llm-rag-go ./cmd/api/

FROM alpine:latest

COPY --from=builder /app/llm-rag-go /bin/llm-rag-go

EXPOSE 8080

CMD ["/bin/llm-rag-go"] 
