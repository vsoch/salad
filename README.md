# salad
>> If you fork this repo... you'll have a salad fork.

**@GodloveD**

[![asciicast](https://asciinema.org/a/160642.svg)](https://asciinema.org/a/160642?speed=2)

## Local Development
You will need to install Go, and then have this repository in your src.

`$GOPATH/src/github.com/vsoch/salad`

```
go run salad.go
```

and to format your code (beware, TABS)

```
go fmt salad.go
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

If you don't give arguments, the entry point will by default run the "fork" executable (the Dockerfile CMD) to print a colored fork:

```
docker run vanessa/salad

 Fork off, already.  

                   ________  .====
                  [________>< :===
                             '==== 

```

And if you ask for help, you can see that you can ask for a spoon too!

```
NAME:
   Salad Fork - Stick this in your salad, with caution, grasshopper!

USAGE:
   salad [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     fork     Fork
     spoon    Spoon
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version   
```

So in fact, the command that we were running is `./salad fork`. Let's ask for a spoon...

```
$ docker run vanessa/salad spoon

 They say I'd rather spoontaneous!  


        __________        .-"""-.
       /          ''''---' .'    \
       \__________....---. '.    /
                          '-...-'
         
```

For either command, you can ask for help.

```
$ docker run vanessa/salad spoon --help
NAME:
   salad spoon - Spoon

USAGE:
   salad spoon [command options] [arguments...]

DESCRIPTION:
   Create a spoon

OPTIONS:
   --message value, -m value  Give wisdom for your spoon.
   --color value, -c value    color your spoon
   
```

And from the above we learn that we can add a custom color or message.

```
$ docker run vanessa/salad spoon --message "I wish I had a cupcake"


I wish I had a cupcake  

          ___           .-""-.
         /   '''---...-'.'  '\\
         \___...---"""-._-.__//
                         '---'

       
```

## Github Actions

We provide an example Github Actions to deploy the metadata to Github pages,
see the [.github/main.workflow](.github/main.workflow). For more about
schema.org extractors, see the [openschemas/extractors](https://github.com/openschemas/extractors) repository.
