FROM registry.cn-shanghai.aliyuncs.com/joy2fun/go:master

WORKDIR /app

COPY . .

RUN go build -o /usr/bin/wessage -v main.go

ENTRYPOINT ["/bin/bash"]

