run:
	go run *.go

wasm: export GOOS=js
wasm: export GOARCH=wasm
wasm:
	go build -o assets/main.wasm

css:
	npx tailwindcss-cli@latest build assets/config.css -o assets/styles.css

build: export NODE_ENV=production
build:
	npx tailwindcss-cli@latest build assets/config.css -o assets/styles.css
	go build
	go build -o assets/main.wasm

local:
	sam local start-api

sync:
	aws s3 sync ./assets s3://timer.pyros2097.dev/assets --delete

deploy:
	sam deploy