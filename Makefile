run:
	export MODE_ENV=$(env) && go run *.go

init_server:
	hz new -module gtools -idl gtools.thrift
	go mod tidy
server:
	hz update -idl gtools.thrift
	go mod tidy

test:
	go test $(path) -v
test_func:
	go test $(path) -v -run $(func)

SERVER = guochenxu@guochenxu-server
# COMMAND = 'cd /home/guochenxu/gtools && bash start.sh'
RUN_NAME = gtools

deploy:
	bash -c 'rm -rf output'
	bash -c 'go env -w GOARCH="amd64" GOOS="linux" CGO_ENABLED="0"'
	bash -c 'go env | grep GOOS'
	bash -c 'echo build begin'
	bash -c 'mkdir -p output/conf'
	bash -c 'cp -r conf/* output/conf'
	bash -c 'cp start.sh output'
	bash -c 'go build -o output/$(RUN_NAME) -ldflags "-s -w"'
	bash -c 'echo build success'
	bash -c 'go env -w GOARCH="amd64" GOOS="windows" CGO_ENABLED="1"'
	bash -c 'go env | grep GOOS'

	rsync -avz --delete ./output/ $(SERVER):~/gtools/
	bash -c 'echo deploy success'
