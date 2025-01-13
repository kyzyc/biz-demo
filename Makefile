.PHONY: gen-demo-proto gen-demo-thrift gen-frontend gen-client gen-server run-demo-proto run-demo-thrift

gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/kyzyc/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --module github.com/kyzyc/biz-demo/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

run-demo-proto:
	@cd demo/demo_proto && go run .

run-demo-thrift:
	@cd demo/demo_thrift && go run .

gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module github.com/kyzyc/biz-demo/gomall/app/frontend -I ../../idl

gen-server:
	@cd app/user && go mod init github.com/kyzyc/biz-demo/gomall/app/user && cwgo server --type RPC --idl ../../idl/user.proto --service user --module github.com/kyzyc/biz-demo/gomall/app/user --pass "-use github.com/kyzyc/gomall/rpc_gen/kitex_gen" -I ../../idl

gen-client:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/user.proto --service user --module github.com/kyzyc/biz-demo/gomall/rpc_gen -I ../idl



