apps=aliddns-mac aliddns-win64.exe aliddns-win32.exe aliddns-linux64 aliddns-arm5 aliddns-arm6 aliddns-arm7

VERSION=`cat version`
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -s"
WINLDFLAGS=-ldflags "-H windowsgui -X main.Version=${VERSION} -X main.Build=${BUILD}"

all: ${apps}

aliddns-mac:
	env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS}  -o $@

aliddns-win64.exe:
	env GOOS=windows GOARCH=amd64 go build ${WINLDFLAGS}  -o $@

aliddns-win32.exe:
	env GOOS=windows GOARCH=386 go build ${WINLDFLAGS}  -o $@

aliddns-linux64:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS}  -o $@

aliddns-arm5:
	env GOOS=linux GOARCH=arm GOARM=5 go build ${LDFLAGS}  -o $@

aliddns-arm6:
	env GOOS=linux GOARCH=arm GOARM=6 go build ${LDFLAGS}  -o $@

aliddns-arm7:
	env GOOS=linux GOARCH=arm GOARM=7 go build ${LDFLAGS}  -o $@

.PHONY:clean
clean:
	@rm -rf ${apps}
