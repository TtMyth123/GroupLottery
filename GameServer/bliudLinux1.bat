set GOPATH="D:\work\Go\GoPath\Public";"D:\work\Go\GoPath\Game"
set GOARCH=amd64
set GOOS=linux

go build -o %cd%\ReleaseWin\main1 %cd%\main1.go

pause
