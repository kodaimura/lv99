FROM golang:1.24

WORKDIR /app

RUN apt-get update \
  && apt-get install -y logrotate \
  && rm -rf /var/lib/apt/lists/*

COPY ../../api .

RUN go mod tidy