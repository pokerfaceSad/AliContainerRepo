FROM golang:1.14.10-alpine
WORKDIR /platform_backend
COPY go.mod go.sum ./
RUN GOPROXY=https://goproxy.cn,direct GO111MODULE=on go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' .
CMD ["sh"]
