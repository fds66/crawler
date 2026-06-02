PATH1=https://learnwebscraping.dev/practice/ecommerce/products/ashenfang-longsword-fan-1001/ 
PATH2=https://learnwebscraping.dev/practice/ecommerce/
.DEFAULT_GOAL := list

list:
	ls
build:
	go build -o web_crawler ./...
run1: build
	./web_crawler ${PATH1}
run2:
	./web_crawler ${PATH2}
test:
	go test ./...