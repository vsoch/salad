# salad
>> If you fork this repo... you'll have a salad fork.

**@GodloveD**

[![asciicast](https://asciinema.org/a/159960.png)](https://asciinema.org/a/159960?speed=2)



## Learning Go
This repository may be about forks and salads, but the underlying motivation is for learning some Go! I'm going to write notes here about some design patterns that I noticed across different repositories.

### Package Names
There can be any number of "go" scripts in the top level folder, and the line at the top that declares a package, if it's the one called "main," tells us that script is the executable entrypoint. The rest of the files must be named according to the repository name. For example, 

```
 /github.com/vsoch/salad
                   salad.go
                   fork.go
```

Let's say that the above [salad.go](salad.go) has the top line say `package main` - this tells us it's the entrypoint to the application. Since the repository folder is `salad` we would put all other scripts in the `salad` namespace (`package salad`).

### Imports
There are some standard imports that I see a lot (e.g., `os`,`string` and `fmt`) and the rest are assumed to be folders again installed at `$GOPATH`. To install a new package you type `go get` and I would bet doing a `go install` will download and install packages that you don't have. I like how the mapping from the repository to the package name and repository is pretty seamless.

The other cool thing about this is that we can thus have **subfolders** in our repository that we add the includes that are, each in themself, akin to submodules. For example:

```
import (
	"os"
	"github.com/urfave/cli"
	"github.com/vsoch/salad/cmd"
	"github.com/vsoch/salad/config"
)
```

The folders "cmd" and "config" are nice organized folders that are for each of package `cmd` and `config`, and just happen to live with the package `salad`. In the same way that the top level folder is called salad and the package and main file are called salad, this is also the case with cmd and config.


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

## Usage

The entry point will by default run the "fork" executable to print a colored fork:

```
docker run vanessa/salad

 Fork off, already.  

                   ________  .====
                  [________>< :===
                             '==== 

```

And if you ask for help, you can see that you can add a message!

```
 docker run vanessa/salad --help
NAME:
   salad fork - Fork

USAGE:
   salad fork [command options] [arguments...]

DESCRIPTION:
   Create a fork

OPTIONS:
   --message value, -m value  Give wisdom for your fork.
   
```
```
 docker run vanessa/salad --message "I wish I had a cupcake"

I wish I had a cupcake  

          _________________  .========
         [_________________>< :======
                             '======== 

```

The base entrypoint for the executable isn't exposed yet (e.g., running just `./salad` instead of `./salad fork`) because we haven't (yet) added other commands or arguments. TBA! 
