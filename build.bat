SET GOOS=linux
SET GOARCH=arm
SET CGO_ENABLED=0
go build main.go
adb push ./main /data/local/tmp/main
adb shell chmod 755 /data/local/tmp/main
adb shell /data/local/tmp/main server -d "$@"
