if [ "$1" == "node" ]; then
    GOOS=js GOARCH=wasm go run -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .
elif [ "$1" == "node-test" ]; then
    GOOS=js GOARCH=wasm go test -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec"
else
    GOOS=js GOARCH=wasm go build -o main.wasm
fi