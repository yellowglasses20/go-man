go-man
===
go-man is simple API test tool.

## help

```sh
$ ./go-man help
go-man is simple API test tool.

Usage:
  go-man [command]

Available Commands:
  get         http get request
  head        http head request
  help        Help about any command
  version     Print the version number of go-man

Flags:
  -H, --Headers stringArray   Any HTTP headers(-H "Authorization:Bearer token")
  -h, --help     help for go-man
  -p, --pretty   Pretty print

Use "go-man [command] --help" for more information about a command.
```

## usage
You can use pretty print with the ```-p``` option

```sh
$ ./go-man get -p http://localhost:8080
{
  "Id": 1234,
  "Text": "well"
}
```

## TODO
### In Progress
- [ ] [IN PROGRESS] delete/post/put
- [ ] [IN PROGRESS] Automatic examination by scenario
