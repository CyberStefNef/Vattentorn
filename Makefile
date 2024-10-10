build-css:
	@echo "Building Tailwind CSS..."
	tailwindcss -i ./web/assets/css/input.css -o ./web/assets/css/styles.css --watch

start-server:
	@echo "Starting Air for Go live-reload..."
	air

start:
	@echo "Starting both Air (Go) and Tailwind CSS build..."
	tailwindcss -i ./web/assets/css/input.css -o ./web/assets/css/styles.css --watch & air

