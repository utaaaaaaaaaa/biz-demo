export ROOT_MOD=github.com/utaaaaaaaaaa/biz-demo/gomall
.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module ${ROOT_MOD}/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --module ${ROOT_MOD}/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/cart_page.proto --service frontend --module ${ROOT_MOD}/app/frontend -I ../../idl

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type=RPC  --service=cart --module=${ROOT_MOD}/rpc_gen --idl=../idl/cart.proto --I=../idl
	@cd app/cart && cwgo server --type=RPC  --service=cart --module=${ROOT_MOD}/app/cart --pass="-use ${ROOT_MOD}/rpc_gen/kitex_gen" --idl=../../idl/cart.proto --I=../../idl
