#==================================================================
# Reference:
#	https://shields.io/
#	https://makefiletutorial.com/
#==================================================================

all: help

## build@编译
.PHONY: build
build:
	@go vet
	@go build
	@echo "\033[31m 🚀  编译完毕\033[0m";


## linter@代码检查
.PHONY: linter
linter:
	@golangci-lint run -c .golangci.yaml || exit 1


## clean@清理编译、日志和缓存等
.PHONY: clean
clean:
	@rm -rf ./bin;
	@rm -rf ./logs;
	@rm -rf ./log;
	@rm -rf ./cache;
	@rm -rf ./pid;
	@rm -rf ./release;
	@rm -rf ./debug;
	@rm -rf ./tmp;
	@rm -rf ./temp;
	@rm -rf ./vendor/*;
	@echo "\033[31m ✅  清理完毕\033[0m";


## commit <msg>@Git提交，如:make push [msg=<message>]
.PHONY: commit
message:=$(if $(msg),$(msg),"rebuilded at $$(date '+%Y/%m/%d %H:%M:%S')")
commit:
	@echo "\033[0;34mPush to remote...\033[0m"
	@git add .
	@git commit -m $(message)
	@echo "\033[0;31m 💿 Commit完毕\033[0m"


## push <msg>@Git提交并推送，如:make push [msg=<message>]
.PHONY: push
push:commit
	@git push #origin master
	@echo "\033[0;31m ⬆️ Push完毕\033[0m"


## update@更新Git和Submodule
.PHONY: update
update:
	@git submodule update --init --recursive;


## test@单元测试
.PHONY: test
test:
	@go test;
	@go test -v -failfast -race -count=1 ./... >/dev/null || exit 1;


## help@查看make帮助
.PHONY: help
help:Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n "s/^##//p" $< | column -t -s '@' |grep --color=auto "^[[:space:]][a-z]\+[[:space:]]"
	@echo
	@echo "更多内容,请参考 https://github.com/hollson\n"
