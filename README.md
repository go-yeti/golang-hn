# golang hacker news api
A little Golang/HackerNews. I've used dep package to take care about the dependencies and GINKGO FOR TESTS.

Building the project
> git clone https://github.com/jamesmallon/golang-hn.git golang-hn

> dep ensure

Running Tests
> ginkgo -v -r 

Runnning
> go build -v -o main

> ./main 

> ./main -csv="./csv"
