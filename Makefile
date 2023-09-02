tw:
	npx tailwindcss -i tailwind.css -o ./assets/tailwind.css

gen:
	(rm -rf db) && (sqlc generate)

o:
	(rm -rf bin) && (go build -ldflags "-w -s" -o bin/main -gcflags=all="-l -B";) && (cp -r assets bin/)

qtc:
	find ./views -name '*.go' -exec rm {} + && find ./templates -name '*.html' -exec cp {} views \; && qtc -dir views -ext html && find ./views -name '*.html' -exec rm {} + && find ./templates -name '*.go' -exec cp {} views \;

nodemon:
	nodemon --watch templates --watch data -e md,go --signal SIGTERM --exec make all --on-change-only

all:
	(make gmt) && (make tw) && (make qtc) && (go run main.go)

rename:
	find . -type f -exec sed -i -e 's/github.com\/Jiang-Gianni\/gthc/github.com\/Jiang-Gianni\/$(name)/g' {} \;

gmt:
	gmt -dir templates/md

doc:
	qtc -dir docgen -ext md && go run docgen/docgen.go