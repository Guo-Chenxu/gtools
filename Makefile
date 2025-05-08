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
	go test ./... -v -run $(func)

SERVER = guochenxu@guochenxu-server
# COMMAND = 'cd /home/guochenxu/gtools && bash start.sh'
RUN_NAME = gtools

deploy: build
	export ENV_MODE=prod
	cd output && bash start.sh
	@echo "deploy success"
build:
	rm -rf output
	mkdir -p output/conf
	cp -r conf/* output/conf/
	cp start.sh output/
	go mod tidy
	go build -o output/$(RUN_NAME) -ldflags "-s -w" -buildvcs=false -race
	@echo "build success"

deploy-win:
	bash -c 'rm -rf output'
	bash -c 'mkdir -p output/conf'
	bash -c 'cp -r conf/* output/conf'
	bash -c 'cp start.sh output'
	bash -c 'GOARCH="amd64" GOOS="linux" go build -o output/$(RUN_NAME) -ldflags "-s -w"'
	bash -c 'echo build success'

	rsync -avz --delete ./output/ $(SERVER):~/gtools/
	@bash -c 'echo deploy success'
