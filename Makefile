Name = fr

.PHONY:  run test build release linux win cross proxy pprof clean


run:
	go run main.go

test:
	go test -v ./...

build:
	go build -ldflags "-s -w"

withoutgui:
	go build -ldflags "-s -w -H=windowsgui"

linux:
	set  GOOS=linux
	go build -ldflags "-s -w" -o $(Name)

win:
	set GOOS=win
	go build -ldflags "-s -w" -o $(Name).exe
	upx -9 $(Name).exe
	
cross:linux win

#学会获取参数才是真的
proxy:
	go env -w  GOPROXY=https://goproxy.io,direct

pprof:run
	go tool pprof -http=:8080 *.pprof
clean:
	rm -f $(Name)
	rm -f *.exe
	rm -f *.pprof

