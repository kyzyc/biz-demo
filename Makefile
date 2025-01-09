.PHONY: gen-demo-proto gen-demo-thrift run-demo-proto run-demo-thrift

gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/kyzyc/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --module github.com/kyzyc/biz-demo/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

run-demo-proto:
	@cd demo/demo_proto && go run .

run-demo-thrift:
	@cd demo/demo_thrift && go run .


