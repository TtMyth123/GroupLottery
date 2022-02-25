set GOPATH="D:\work\Go\GoPath\Game"
set GOARCH=amd64
set GOOS=linux

D:\Go\bin\go build -o %cd%\ReleaseWin\LotteryResultServer %cd%\main.go

pause
