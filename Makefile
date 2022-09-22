.PHONY: build
# 编译生成可执行文件
build:
	mkdir -p bin/ && go build  -o ./bin/out 