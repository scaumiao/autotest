PKGMAP=Mproto/auth/policy.proto=auth/policy
GOOGLEAPI=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis


AUTH_PROTO_DIR=$(PWD)/auth
AUTH_PROTO=$(AUTH_PROTO_DIR)/policy.proto \
$(AUTH_PROTO_DIR)/group.proto

REWRITE_PROTO_DIR=$(PWD)/rewrite
REWRITE_PROTO=$(REWRITE_PROTO_DIR)/rewrite.proto 


PKG=${GOPATH}/src/github.com/yourhe/grpc-gateway
GO_PLUGIN=bin/protoc-gen-go
GATEWAY_PLUGIN=$(GOPATH)/bin/protoc-gen-grpc-gateway
SWAGGER_PLUGIN=${GOPATH}/bin/protoc-gen-swagger

define BUILDPROTO
	protoc -I/usr/local/include \
			-I. \
			-I$(1)/$(2) \
			-I${GOPATH}/src \
			-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			-I${GOPATH}/src/github.com/protocolbuffers/protobuf/src \
			-I${GOPATH}/src/github.com/scaumiao/autotest \
			--go_out=plugins=grpc:${GOPATH}/src/. \
			--plugin=$(GATEWAY_PLUGIN) \
			--grpc-gateway_out=logtostderr=true:${GOPATH}/src/. \
			$(1)/$(2)/$(2).proto
endef

define BUILDPROTO2
	protoc -I/usr/local/include \
			-I. \
			-I$(1)/$(2) \
			-I${GOPATH}/src \
			-I${GOPATH}/src/github.com/yourhe/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:$(1)/$(2)/. \
			--plugin=$(GATEWAY_PLUGIN) \
			--grpc-gateway_out=logtostderr=true:$(1)/$(2)/. \
			$(1)/$(2)/$(2).proto
endef

define BUILDSWAGGER
	protoc -I/usr/local/include \
		-I. \
		-I$(1)/$(2) \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/yourhe/grpc-gateway/third_party/googleapis \
		--plugin=$(SWAGGER_PLUGIN) \
		--swagger_out=logtostderr=true:$(1)/$(2)/. \
		$(1)/$(2)/$(2).proto
endef


define BUILDAPI
	docker run --rm -v "$(1):/work" swaggerapi/swagger-codegen-cli generate \
		-i /work/$(2)/$(2).swagger.json \
		-t /work/swagger-codegen/swagger_template_js \
		-l javascript --additional-properties packageName=$(2) \
		-o /work//$(2)/clients/$(2)-js
endef

# role:
# 	$(call BUILDPROTO,$(PWD),$@)
# 	$(call BUILDSWAGGER,$(PWD),$@)
# 	$(call BUILDAPI,$(PWD),$@)
# 	cp -rf $(PWD)/$@/clients/$@-js/src/api/*.js ${DR2AMUI}/proto/

stages:
	$(call BUILDPROTO,$(PWD),$@)

task:
	$(call BUILDPROTO,$(PWD),$@)

.PHONY: stages task


# java -jar scripts/swagger-codegen-cli-2.2.3.jar generate -i ./docs/api.swagger.json -l go -o ./sdk/go/hydra/swagger
