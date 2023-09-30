# Go api server

this is my api server of my personal project, [iDid](https://idid.com).

## Requirements

| Name | Version | Description                                            | Github Repo                                     |
| ---- | ------- | ------------------------------------------------------ | ----------------------------------------------- |
| Go   | 1.16    | The Go programming language                            | [golang/go](https://github.com/golang/go)       |
| Just | 1.14.0  | A handy way to save and run project-specific commands. | [casey/just](https://github.com/casey/just)     |
| Air  | 1.29.0  | Live reload for Go apps                                | [cosmtrek/air](https://github.com/cosmtrek/air) |

For information on the above tools, please refer to the respective Github Repositories.

## Getting start

After cloning this repository, run this commands:
``` bash
$ just init
```

To run the server, run this commands:
``` bash
$ just run
```
The server will restart as the file changes, thanks to Air.

## Directory structure

``` bash
. # root
├── Brewfile
├── Brewfile.lock.json
├── Justfile
├── README.md
├── cmd
│   └── server
│       └── main.go # entry point
├── ent # entgo.io
│   ├── generate.go
│   └── schema
│       ├── mixins.go
│       └── user.go
├── go.mod
├── go.sum
└── internal # internal packages
    ├── controllers
    │   └── user.go
    ├── interface
    │   └── user.go
    └── settings
        └── settings.go
```
