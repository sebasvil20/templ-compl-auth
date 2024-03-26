go_run:
	@go run cmd/api/main.go

templ:
	@templ generate

run: templ go_run
