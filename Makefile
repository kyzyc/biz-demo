export ROOT_MOD=github.com/kyzyc/biz-demo
.PHONY: gen-demo-proto gen-demo-thrift gen-frontend gen-client gen-server run-demo-proto run-demo-thrift

gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module ${ROOT_MOD}/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --module ${ROOT_MOD}/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

run-demo-proto:
	@cd demo/demo_proto && go run .

run-demo-thrift:
	@cd demo/demo_thrift && go run .

gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/product_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl

gen-server:
	@cd app/user && go mod init ${ROOT_MOD}/app/user && cwgo server --type RPC --idl ../../idl/user.proto --service user --module ${ROOT_MOD}/app/user --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl

gen-client:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/user.proto --service user --module ${ROOT_MOD}/rpc_gen -I ../idl

.PHONY: gen
gen:
	@script/gen.sh

#.PHONY: gen-user
#gen-user:
#	@cd app/user && cwgo server --type RPC  --service user --module  ${ROOT_MOD}/app/user  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/user.proto
#	@cd rpc_gen && cwgo client --type RPC  --service user --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/user.proto

#.PHONY: gen-product
#gen-product:
#	@cd app/product && cwgo server --type RPC  --service product --module  ${ROOT_MOD}/app/product  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/product.proto
#	@cd rpc_gen && cwgo client --type RPC  --service product --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/product.proto
#
#.PHONY: gen-cart
#gen-cart:
#	@cd app/cart && cwgo server --type RPC  --service cart --module  ${ROOT_MOD}/app/cart  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/cart.proto
#	@cd rpc_gen && cwgo client --type RPC  --service cart --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/cart.proto




