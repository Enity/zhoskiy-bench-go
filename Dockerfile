FROM golang:1.17.1-alpine as builder
WORKDIR /app
COPY . .
# for ci cache
ENV GOPATH=/app/.gopath-cache
# golang-alpine dont includes git
RUN apk add --no-cache git
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

############################

FROM scratch
WORKDIR /app
COPY --from=builder /app/main ./main
CMD ["./main"]
