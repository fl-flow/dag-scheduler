FROM golang:1.18.1-alpine as builder

# RUN apk add --no-cache \
#     wget \
#     make \
#     git \
#     gcc \
#     binutils-gold \
#     musl-dev
WORKDIR /build

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update --no-cache && apk add --no-cache gcc musl-dev

# Cache dependencies
COPY go.mod .
COPY go.sum .
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/httpserver ./main.go


# Executable image
FROM alpine

ENV TZ Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/httpserver ./httpserver

EXPOSE 8000

CMD [ "./httpserver" ]