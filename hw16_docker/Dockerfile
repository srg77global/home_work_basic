FROM golang:1.23.1 AS builder

WORKDIR /home_work_basic

COPY /hw15_go_sql ./

RUN go mod download && make install-deps && go build -o hw16 ./cmd/main.go


FROM alpine:latest

ENV CGO_ENABLED=0

WORKDIR /home_work_basic

COPY --from=builder ./home_work_basic/hw16 ./
COPY --from=builder ./home_work_basic/.env ./

EXPOSE 8080

CMD ["./hw16"]