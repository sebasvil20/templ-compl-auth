go_run:
	@echo "\033[96m╭―――――――――――――╮\n│ Running API │\n╰―――――――――――――╯\033[0m"
	@go run cmd/api/main.go

templ:
	@echo "\033[96m╭――――――――――――――――――――――╮\n│ Generating templates │\n╰――――――――――――――――――――――╯"
	@templ generate
	@echo "\n\033[32m╭――――――――――――――――――――――――――╮\n│ (✓) Generating templates │\n╰――――――――――――――――――――――――――╯\033[0m"

tailwind:
	@echo "\033[96m╭―――――――――――――――――――――――╮\n│ Compiling tailwindcss │\n╰―――――――――――――――――――――――╯"
	@npx tailwindcss -i ./static/input.css -o ./static/output.css
	@echo "\n\033[32m╭―――――――――――――――――――――――――――╮\n│ (✓) Compiling tailwindcss │\n╰―――――――――――――――――――――――――――╯\033[0m"

run: tailwind templ go_run
