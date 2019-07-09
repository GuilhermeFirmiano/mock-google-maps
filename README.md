# mock-google-maps
mock service google maps

## Coding

VS Code is probably the best option for coding in Go.

Download: [VS Code](https://code.visualstudio.com/download)

* VS Code has a really good support for Go. Open the project and install all suggested dependencies, it's enough for start coding.

## Installing Go

Follow [these](https://github.com/minio/cookbook/blob/master/docs/how-to-install-golang.md) instructions.

## dep

**dep** is most popular package manager for Go projects.

### Installing

### macOS

```
brew install dep
brew upgrade dep
```

### On other platforms

```
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

More info can be found [here](https://github.com/golang/dep)

## Building

Go to project folder containing **main.go** file, and then:

```
go build
```

## Tests

Go to root project folder, and then:

### Running all tests.

You will need to have Docker Compose running with Redis and Postgres to run tests.
```
go test ./...
```

### Running a spefic test

```
go test {{TEST_FILE_FOLDER}} -run ^{{TEST_NAME}}$
```

## Running

Go to project folder containing **main.go** file, and then:

```
go run main.go
```

## Building Docker 

* Restore all dependencies
* Run unit tests
* Run integration tests
* Build

Go to the root folder, and then:

```
docker build -t mock-google-maps .
```

## Compose

Building

```
docker-compose build
```

Up

```
docker-compose up
```

**docker-compose up** will automatically setup the API.