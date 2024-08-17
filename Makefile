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

deploy:
	rsync -avz --delete --exclude='.history/' --exclude='logs/' ./ $(SERVER):~/gtools/
	# ssh $(SERVER) $(COMMAND)