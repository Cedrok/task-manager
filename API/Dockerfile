FROM golang:1.22 AS builder
USER root
WORKDIR /task-manager

COPY go.mod go.sum /task-manager/
RUN go mod download

COPY API /task-manager/API
COPY common /task-manager/common

WORKDIR /task-manager/API
RUN make

FROM golang:latest
WORKDIR /task-manager

COPY --from=builder /task-manager/API/task-api /task-manager/
COPY API/configs/.env /task-manager/configs/

EXPOSE 9000

ENTRYPOINT [ "/task-manager/task-api" ]