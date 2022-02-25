set GOPATH="D:\work\go\Public";"D:\work\go\WorkSpace2"
SET CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux

go build -o %cd%\ReleaseWin\UserInfoRpc %cd%\main.go

pause
