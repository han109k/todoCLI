# TODO cli app written in Go

## Add environment variables

```console
export CLIAPP_DATAFIE=$HOME/todos.json
```

OR put it in $HOME/.cliApp.yaml

```console
datafile: $HOME/todos.json
```

---

## Start Using

```console
go build
./cliApp ...
OR
go run main.go ...
```

- Help

```console
go run main.go --help
```

- Add a todo item with default priority 3

```console
go run main.go add "Some todo"
```

- Add a todo item with a priority 5

```console
go run main.go add "format list" -p5
```

- List to do items

```console
go run main.go list
go run main.go list --all
go run main.go list --done
```

- Mark 3rd todo item as done

```console
go run main.go done 3
```
