FROM golang:1.23.1 AS builder

WORKDIR /home_work_basic

COPY /hw15_go_sql ./

RUN go mod download && make install-deps && go build -o hw16_docker ./cmd/main.go


FROM alpine:latest

WORKDIR /home_work_basic

COPY --from=builder /home_work_basic/hw16_docker ./

EXPOSE 8088
EXPOSE 5033

CMD ["./hw16_docker"]