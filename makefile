run:
	go run *.go

wasm: export GOOS=js
wasm: export GOARCH=wasm
wasm:
	go build -o assets/main.wasm

css:
	npx tailwindcss-cli@latest build assets/config.css -o assets/styles.css

deploy: export NODE_ENV=production
deploy:
	npx tailwindcss-cli@latest build assets/config.css -o assets/styles.css
	go build -o main
	sam deploy
	make wasm
	aws s3 sync ./assets s3://timer.pyros2097.dev/assets --delete --exclude main.wasm
	brotli -Z -j assets/main.wasm
	mv assets/main.wasm.br assets/main.wasm
	aws s3 cp assets/main.wasm s3://timer.pyros2097.dev/assets/main.wasm --content-encoding br --content-type application/wasm
	aws cloudfront create-invalidation --distribution-id E53G56K101AX2 --paths "/*"

local:
	sam local start-api
