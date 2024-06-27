run: build
	@ echo "Running server..."
	@./bin/fastcache

build:
	@go build -o bin/fastcache .

