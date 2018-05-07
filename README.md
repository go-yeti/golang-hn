# Hacker News - Golang App
A little **Golang HackerNews App**. I've used [DEP] package to **Dependency Management** and [GINKGO] for **Unitary Tests**.

**SOLID**, **KISS** and **DRY** principles was applied as well as some useful and relevant (to this application) **Design Patterns**, such as **DAO**, **Dependency Injection** and more.

[DEP]: <https://github.com/golang/dep>
[GINKGO]: <https://onsi.github.io/ginkgo/>

#### Get the project:
```sh
$ git clone https://github.com/jamesmallon/golang-hn.git golang-hn
```

#### Updating the dependencies:
I would recommend to add the dep binary file to your /usr/local/bin
```sh 
$ dep ensure
```
or if you don't want to add:
```sh
$ $GOPATH/bin/dep ensure
```

#### Build:
```sh
$ go build -v -o main
```

#### Running Tests:
If you have [GINKGO] in your /usr/local/bin, run:
```sh 
$ ginkgo -v -r 
```
if not:
```sh 
$ $GOPATH/bin/ginkgo -v -r 
```

#### Running:
```sh
$ ./main 
```

#### Running to save in a CSV files:
```sh 
$ ./main -csv=./csv/ 
