# Poker

一个类Docker的容器实现

A container technology as like Docker.

看似神奇的Docker，实际上只是执行了Linux的System Calls，这力量本不属于Docker，Linux赋予其魔力

The seemingly magical Docker actually only executes Linux's system calls. This power does not belong to Docker, Linux gives it magic.

## The functions to be realized(将要实现的功能)

### Poker commands:

> build	Build an image from a Dockerfile
>
> exec	Run a command in a running container
>
> images	List images
>
> logs	Fetch the logs of a container✅
>
> ps	List containers✅
>
> rename	Rename a container✅
>
> restart	Restart one or more containers✅
>
> rm	Remove one or more containers✅
>
> rmi	Remove one or more images
>
> run	Run a command in a new container✅
>
> start Start one or more stopped containers✅
>
> stop	Stop one or more running containers✅
>
> top	Display the running processes of a container
>

### Poker-daemon:

Future: use sysctl hosting

### Use:

```shell
$ make
$ poker -h
Poker is a container technology as like docker

Usage:
  poker [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  logs        Fetch the logs of a container
  ps          List containers
  rename      Rename a container
  restart     Restart one or more containers
  rm          Remove one or more containers
  run         Run a command in a new container
  start       Start one or more exited containers
  stop        Stop one or more running containers

Flags:
  -h, --help      help for poker
  -v, --version   version for poker

Use "poker [command] --help" for more information about a command.

```

