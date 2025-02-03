# ttodo

TODO management CLI application.

## Usage

### Prerequisites

- [Go](https://go.dev/)
- [SQLite](https://www.sqlite.org/)

### Installation

```bash
$ git clone <repo-url>
$ cd ttodo

# With go itself
$ go build -o bin/ttodo

# With Makefile
$ make build

# Running the application
$ ./bin/ttodo
```

### Help menu

```bash
Usage:
  todo <command> [options]

Commands:
  add       -t "Title"           Add a new todo
  toggle    -i ID                Toggle completion status of a todo
  update    -i ID -t "New Title" Update a todo title
  remove    -i ID                Remove a todo
  clean                          Remove completed todos
  list                           Show all todos
  help, --help, -h               Show help menu
```

For convenience, you can set up an alias. It is my preferred way ;)

```bash
alias ttodo='./Code/ttodo/bin/ttodo'
```
