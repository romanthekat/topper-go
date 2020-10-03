# topper-go
This little tool parses your `.bash_history` file, and calculates the most frequent commands.

## How to build
* with makefile: 

`make build` will produce topper-go executable file

`make build.mac` or `make build.linux`

* with go command

`go build -o topper-go main.go`

 or

`go build -o tg main.go`

that will produce `topper-go` or `tg` executable file.

## How to use

`tg` - to show top 10 commands

`tg 5` - to show top 5 commands

with output format:
```
  155: ping ya.ru (x26)
   21: ./launch.sh  (x26)
   39: ls (x16)
   94: git pull origin (x13)
  228: git status (x10)
  309: psql (x7)
...
```

Where for `155: ping ya.ru (x26)`:

- `155` - command number in terms of bash, e.g. `!155` allows to repeat command execution
- `ping ya.ru` - command itself
- `(x26)` - current execution count
