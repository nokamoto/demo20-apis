PROTOTOOL=docker run -v $(PWD):/work --rm uber/prototool:1.10.0 prototool

all:
	${PROTOTOOL} format -d . || ${PROTOTOOL} format -w .
	docker run --rm -v $(PWD):/protobuf -w /protobuf nokamotohub/demo20-protoc:latest ./scripts/build.sh
