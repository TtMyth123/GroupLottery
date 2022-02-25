set GOPATH=%GOPATH%;"D:\work\Go\GoPath\Public";"D:\work\Go\GoPath\Game"
set GOARCH=amd64
set GOOS=linux

go build -o %cd%\ReleaseWin\Staff %cd%\main.go

pause
