clear:
	@rm -rf output
	@mkdir output

build: clear
	@go build -o ./output/ias-simulator ./cmd/ias-simulator/main.go

run: build
	@cd output && ./ias-simulator
