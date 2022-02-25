set GOPATH="D:\work\go\Public";"D:\work\go\WorkSpace2"
set GOARCH=amd64
set GOOS=linux

go build -o %cd%\ReleaseWin\Admin %cd%\main.go

pause
