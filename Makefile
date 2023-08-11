tw:
	npx tailwindcss -i tailwind.css -o ./assets/tailwind.css

gen:
	(rm -rf db) && (sqlc generate)

o:
	(rm -rf bin) && (go build -ldflags "-w -s" -o bin/main -gcflags=all="-l -B";) && (cp -r assets bin/)

qtc:
	find ./views -name '*.go' -exec rm {} + && find ./templates -name '*.html' -exec cp {} views \; && qtc -dir views -ext html && find ./views -name '*.html' -exec rm {} +

nodemon:
	nodemon --watch templates -e html,go --signal SIGTERM --exec make all --on-change-only

all:
	(make tw) && (make qtc) && (go run main.go)