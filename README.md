# Hacker News - Golang App
A little **Golang HackerNews App**. I've used [DEP] package to **Dependency Management** and [GINKGO] for **Unitary Tests**.

 - **SOLID**, **KISS** and **DRY** principles was applied as well as some useful and relevant (to this application) **Design Patterns**, such as **DAO**, **Dependency Injection** and more.

 - By default the app return 20 top stories top Hacker News, but you can specify the quantity by using the flag **-tpqtt**.

 - It is possible to save the result to a .csv file by using the flag **-csv**, which expects the path to save the file. The file name is automatically generated, accordingly with the quantity of topstories and the timestamp, to avoid name conflict and overwritings.

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

#### Running (default quantity of top stories - 20):
```sh
$ ./main 
```

#### Running and specifying a quantity of 45 topstories:
```sh
$ ./main -tpqtt=45
```

#### Running to save in a CSV files (default quantity of top stories - 20):
```sh 
$ ./main -csv=./csv/ 
```

#### Running to save in a CSV files (default quantity of top stories - 20):
```sh 
$ ./main -csv=./csv/ -tpqtt=45
```
