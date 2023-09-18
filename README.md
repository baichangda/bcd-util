# 跨平台编译
- macos跨平台编译
    - CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/bcd-util-macos
    - CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/bcd-util-linux
    - CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/bcd-util-windows.exe

- windows跨平台编译
    - set CGO_ENABLED=0&&set GOOS=darwin&&set GOARCH=amd64&&go build -o build/bcd-util-macos
    - set CGO_ENABLED=0&&set GOOS=linux&&set GOARCH=amd64&&go build -o build/bcd-util-linux
    - set CGO_ENABLED=0&&set GOOS=windows&&set GOARCH=amd64&&go build -o build/bcd-util-windows.exe

# 错误处理
使用第三方包errors

需要注意的是、为了获得更好的堆栈信息、需要遵循如下规则
- 所有调用第三方包返回的error、必须使用 ***errors.WithStack()*** 包装后再返回上层
- 所有调用自己方法返回的error、直接返回到上层即可
- 所有根据errors方法产生的error、直接返回到上层即可

这样可以保证、所有返回的错误都有堆栈信息、且只包含一份堆栈信息

因为使用WithStack处理errors返回或者处理过的错误、会导致附加多次堆栈信息