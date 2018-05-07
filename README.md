
# Hacker News - Golang App
A little **Golang HackerNews App**. I've used [golang/dep] package to **Dependency Management** and [GINKGO] for **Unitary Tests**.

**SOLID**, **KISS** and **DRY** principles was applied as well as some useful and relevant (to this application) **Design Patterns**, such as **DAO**, **Dependency Injection** and more.

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

#### Build:
```sh
$ go build -v -o main
```

#### Running Tests:
```sh 
$ ginkgo -v -r 
```

#### Running:
```sh
$ ./main 
```

#### Running to save in a CSV files:
```sh 
$ ./main -csv="./csv" 
```
