PATH1=https://learnwebscraping.dev/practice/ecommerce/products/ashenfang-longsword-fan-1001/ 
.DEFAULT_GOAL := list

list:
	ls
build:
	go build -o web_crawler ./...
run: build
	./web_crawler ${PATH1}
test:
	go test ./...