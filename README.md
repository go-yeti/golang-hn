# Hacker News - Golang App
A little **Golang HackerNews App**.

 - By default the app return 20 top stories top Hacker News, and other values can be specified by using the flag **-qtt**.

 - To save the result to a .csv file use the flag **-csv**

[DEP]: <https://github.com/golang/dep>

#### Updating the dependencies:
I would recommend to add the dep binary file to your /usr/local/bin
```sh 
$ dep ensure
```

#### Running (default quantity of top stories = 20):
```sh
$ ./main 
```

#### Running and specifying a quantity of 45 topstories:
```sh
$ ./main -qtt 45
```

#### Running to save in a CSV files (default quantity of top stories = 20):
```sh 
$ ./main -csv
```

#### Running to save in a CSV files and selecting 45 top stories:
```sh 
$ ./main -csv -qtt 45
```

**by [James Mallon]**

 [James Mallon]: <https://www.linkedin.com/in/thiago-mallon/>
