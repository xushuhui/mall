APP_API_PROTO_FILES=$(shell find api/app -name *.proto)
COUPON_API_PROTO_FILES=$(shell find api/coupon -name *.proto)
Interface_API_PROTO_FILES=$(shell find api/mall/interface -name *.proto)
Admin_API_PROTO_FILES=$(shell find api/mall/admin -name *.proto)
Order_API_PROTO_FILES=$(shell find api/order -name *.proto)
User_API_PROTO_FILES=$(shell find api/user -name *.proto)
Payment_API_PROTO_FILES=$(shell find api/payment -name *.proto)
Spu_API_PROTO_FILES=$(shell find api/spu -name *.proto)
.PHONY: api.app
api.app:
		protoc --proto_path=./api/app/service \
    	       --proto_path=./third_party \
     	       --go_out=paths=source_relative:./app/app/service/internal/service \
     	       --go-http_out=paths=source_relative:./app/app/service/internal/service \
     	       --go-grpc_out=paths=source_relative:./app/app/service/internal/service \
    	       --openapi_out=fq_schema_naming=true,default_response=false:. \
    	       $(APP_API_PROTO_FILES)
.PHONY: api.user
api.user:
		protoc --proto_path=./api/user/service \
    	       --proto_path=./third_party \
     	       --go_out=paths=source_relative:./app/user/service/internal/service \
     	       --go-http_out=paths=source_relative:./app/user/service/internal/service \
     	       --go-grpc_out=paths=source_relative:./app/user/service/internal/service \
    	       --openapi_out=fq_schema_naming=true,default_response=false:. \
    	       $(User_API_PROTO_FILES)
.PHONY: api.order
api.order:
		protoc --proto_path=./api/order/service \
    	       --proto_path=./third_party \
     	       --go_out=paths=source_relative:./app/order/service/internal/service \
     	       --go-http_out=paths=source_relative:./app/order/service/internal/service \
     	       --go-grpc_out=paths=source_relative:./app/order/service/internal/service \
    	       --openapi_out=fq_schema_naming=true,default_response=false:. \
    	       $(Order_API_PROTO_FILES)
.PHONY: api.spu
api.spu:
		protoc --proto_path=./api/spu/service \
    	       --proto_path=./third_party \
     	       --go_out=paths=source_relative:./app/spu/service/internal/service \
     	       --go-http_out=paths=source_relative:./app/spu/service/internal/service \
     	       --go-grpc_out=paths=source_relative:./app/spu/service/internal/service \
    	       --openapi_out=fq_schema_naming=true,default_response=false:. \
    	       $(Spu_API_PROTO_FILES)
.PHONY: api.payment
api.payment:
		protoc --proto_path=./api/payment/service \
    	       --proto_path=./third_party \
     	       --go_out=paths=source_relative:./app/payment/service/internal/service \
     	       --go-http_out=paths=source_relative:./app/payment/service/internal/service \
     	       --go-grpc_out=paths=source_relative:./app/payment/service/internal/service \
    	       --openapi_out=fq_schema_naming=true,default_response=false:. \
    	       $(Payment_API_PROTO_FILES)
.PHONY: api.coupon
api.coupon:
		protoc --proto_path=./api/coupon/service \
    	       --proto_path=./third_party \
     	       --go_out=paths=source_relative:./app/coupon/service/internal/service \
     	       --go-http_out=paths=source_relative:./app/coupon/service/internal/service \
     	       --go-grpc_out=paths=source_relative:./app/coupon/service/internal/service \
    	       --openapi_out=fq_schema_naming=true,default_response=false:. \
    	       $(COUPON_API_PROTO_FILES)
.PHONY: api.mall.interface
api.mall.interface:
		protoc --proto_path=./api/mall/interface \
    	       --proto_path=./third_party \
     	       --go_out=paths=source_relative:./app/mall/interface/internal/service \
     	       --go-http_out=paths=source_relative:./app/mall/interface/internal/service \
     	       --go-grpc_out=paths=source_relative:./app/mall/interface/internal/service \
    	       --openapi_out=fq_schema_naming=true,default_response=false:. \
    	       $(Interface_API_PROTO_FILES)
.PHONY: api.mall.admin
api.mall.admin:
		protoc --proto_path=./api/mall/admin \
    	       --proto_path=./third_party \
     	       --go_out=paths=source_relative:./app/mall/admin/internal/service \
     	       --go-http_out=paths=source_relative:./app/mall/admin/internal/service \
     	       --go-grpc_out=paths=source_relative:./app/mall/admin/internal/service \
    	       --openapi_out=fq_schema_naming=true,default_response=false:. \
    	       $(Admin_API_PROTO_FILES)