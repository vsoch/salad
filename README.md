# salad
>> If you fork this repo... you'll have a salad fork.

**@GodloveD**

[![asciicast](https://asciinema.org/a/159960.png)](https://asciinema.org/a/159960?speed=2)


## Local Development
You will need to install Go, and then have this repository in your src.

`$GOPATH/src/github.com/vsoch/salad`

```
go run fork.go
```

## Docker


You can clone the repository first and build locally:

```
docker build -t vanessa/salad .
docker run vanessa/salad
```

or run directly from Docker Hub:

```
docker run vanessa/salad
```
