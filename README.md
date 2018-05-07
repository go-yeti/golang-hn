# Hacker News - Golang App
A little Golang/HackerNews. I've used [golang/dep] package to take care about the dependencies and [GINKGO] for tests.

[golang/dep]: <https://github.com/golang/dep>
[GINKGO]: <https://onsi.github.io/ginkgo/>

#### Get the project:
```sh
$ git clone https://github.com/jamesmallon/golang-hn.git golang-hn
```

#### Updating the dependencies:
```sh 
$ dep ensure
```

#### Running Tests:
```sh 
$ ginkgo -v -r 
```

#### Build:
```sh
$ go build -v -o main
```

#### Running:
```sh
$ ./main 
```

#### Running to save in a CSV files:
```sh 
$ ./main -csv="./csv" 
```
