PATH1=https://learnwebscraping.dev/practice/ecommerce/products/ashenfang-longsword-fan-1001/ 
PATH2=https://learnwebscraping.dev/practice/ecommerce/
MAXPAGES=5
MAXCONCURRENCY=3
.DEFAULT_GOAL := list


list:
	ls
build:
	go build -o web_crawler ./...
run1: build
	./web_crawler ${PATH1} ${MAXCONCURRENCY} ${MAXPAGES}
run2:
	./web_crawler ${PATH2} ${MAXCONCURRENCY} ${MAXPAGES}
test:
	go test ./...