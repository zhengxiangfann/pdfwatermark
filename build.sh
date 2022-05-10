# mac
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s"  -buildmode=c-shared -o watermark.so watermark.go
chmod 0777 watermark.so

# linux
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"  -buildmode=c-shared -o watermark.so watermark.go
chmod 0777 watermark.so

# win
CGO_ENABLED=1 GOOS=win64 GOARCH=amd64 go build -ldflags="-w -s"  -buildmode=c-shared -o watermark.so watermark.go
chmod 0777 watermark.so


# 编译为静态库
# mac darwin
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s"  -buildmode=c-archive -o libwatermark.a watermark.gon

# linux
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"  -buildmode=c-archive -o libwatermark.a watermark.go
