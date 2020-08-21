PROTOTOOL=docker run -v $(PWD):/work --rm uber/prototool:1.10.0

all:
	${PROTOTOOL} prototool format -d . || ${PROTOTOOL} prototool format -w .
