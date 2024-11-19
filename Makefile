build:
	go build -o betis-oprec main.go

# Run command that depends on build
run: build
	./betis-oprec

# Watch command to re-run on file changes
watch:
	reflex -s -r '\.go$$' -- make run