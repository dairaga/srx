.PHONY: clean test srx js std enum form

GOPATH=$(shell go env GOPATH)
WASMEXEC=${GOPATH}/bin/wasmbrowsertest
WASM_HEADLESS=on
PKG=github.com/dairaga/srx

test: ${WASMEXEC}
	${MAKE} js
	${MAKE} enum
	$(MAKE) srx
	
srx: ${WASMEXEC}
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test ${PKG} -exec=${WASMEXEC} -test.v

js: ${WASMEXEC}
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test ${PKG}/js -exec=${WASMEXEC} -test.v

enum: ${WASMEXEC}
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test ${PKG}/enum -exec=${WASMEXEC} -test.v
	
${WASMEXEC}:
	go get -u github.com/dairaga/wasmbrowsertest