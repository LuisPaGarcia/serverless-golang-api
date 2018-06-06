# serverless-golang-api
Golang api example

How to create two functions to POST and GET info with golang, aws and serverless framework.

### Prereq

1. You'll need NPM installed. [How to install NPM](https://www.npmjs.com/get-npm).
1. You'll need serverless framework installed. 

```sh
$ npm i serverless -g
```

3. You'll need to have your serverless credentials added, you can see this video [Set Serverless cretenials 4min](https://www.youtube.com/watch?v=HSd9uYj2LJA)

### Ready

MacOS & Linux
```sh
$ git clone https://github.com/LuisPaGarcia/serverless-golang-api.git
$ cd serverless-golang-api
$ make 
$ sls deploy
```

Windows
```cmd
> git clone https://github.com/LuisPaGarcia/serverless-golang-api.git
> cd serverless-golang-api
> dep ensure
> go build -s -w -o bin/hello hello/main.go
> go build -s -w -o bin/get get/main.go
> sls deploy
```

Restult
```sh
endpoints:
  POST - https://ajshdjashdjaskh.execute-api.us-east-3.amazonaws.com/dev/hello
  GET - https://ajshdjashdjaskh.execute-api.us-east-3.amazonaws.com/dev/get
functions:
  hello: myservice-dev-hello
  get: myservice-dev-get
```

You can test yout endpoints with [Postman](https://www.getpostman.com/)
