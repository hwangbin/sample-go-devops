# 1단계: 빌드 환경 (요리사)
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

# 2단계: 실행 환경 (서빙하는 사람)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 1234
CMD ["./main"]