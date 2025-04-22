# Learning GO
Repository for my dummy projects, files and resources that I created during learning Go.

# Purpose of this repo
Only has one purpose: track my progress with GO. This repo will not have any other purpose than just contain dummy/demo files, projects or other resources.
This might contain something useful, might not. Who knows...

# GO commands
Following table shows the commands, syntax, description, example of usage and notes about them.

| Command | Syntax | Description | Example | Notes |
|----------:|----------|----------|----------|----------|
| mod | go mod init <module path and name> | Create a module for building the entire application. Creates a go.mod file which contains the root namespace, go version and external dependencies. | go mod init dummy.com/myproject | Have to use it when the project uses external dependencies. |
| build | go build | Builds the application based on the .mod file. Pulls the dependencies and creates an .exe file. | go build | Not crucial for the project, only good for creating binary distribution. |
| run | go run <filename with .go extension> OR ./ <the .exe filename> | Runs the application. | go run my-project.go OR ./my-project |  |

# Modules, project structure, files and packages
A GO project starts with a folder and a go file like `app.go`. As the project grows, new folders and files appears. In a Go project a new package, which is not the `main` package, created in a new folder and in that folder different files contains the logic of that package. One application has to have only one main package and it is the entry point of the application. The packages and the main package collectivelly forms the application module.

## Example file structure
```
my-go-project/
├── go.mod                # Go module file (dependencies, name)
├── go.sum                # Module checksum file (automatically generated)
├── README.md             # Docs
├── main.go               # Entry point (main function)
├── internal/             # Internal packages (cannot import in other projects)
│   └── logic/
│       └── calculator.go # For example: bussiness logic
├── pkg/                  # External packages (can import for other projects)
│   └── utils/
│       └── strings.go    # Utility functions
├── api/                  # REST API endpoints
│   └── handlers.go
├── config/               # Config files
│   └── config.go
├── cmd/                  # CLI or other entry points
│   └── mytool/
│       └── main.go
└── test/                 # Tests
    └── logic_test.go
```

# Usage of fmt.Scan()
Scan is good for asking an input value from the user. When we ask for a string value we can't process 

# Control statements
## For loop
```
for initialisation_expression; evaluation_expression; iteration_expression{
   // one or more statement
}

Iteration by incrementing an iterator
for i := 0; i < 10; i++ {
    fmt.Print(i)
}

While loop in GO
i := 0
for i < 5 {
    fmt.Print(i)
    i++
}

Infite loop in GO
i := 0
for {
    fmt.Print(i)
    if i == 5 {
        break;
    }
}
```

## if-else
```
if condition {
    // one or more statement
} else {
    // one or more statement
}
```

## switch
```
switch <variable> {
    case <option>:
		// code that covers this option
        break
	case <other option>:
		// code that covers this option
        break
	default:
		// default branch
        return
}
```