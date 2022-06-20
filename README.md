# TODO cli app written in Go

## Development

```javascript
go build
./cliApp [option] [args]
OR
go run cliApp.go [option] [args]
```

Help

```css
go run main.go --help
```

Add a todo item with default priority 3

```css
go run main.go add "Some todo"
```

Add a todo item with a priority 5

```css
go run main.go add "format list" -p5
```

List to do items

```css
go run main.go list
```

Mark 3rd todo item as done

```css
go run main.go done 3
```
