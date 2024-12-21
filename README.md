# Learning GO
Repository for my dummy projects, files and resources that I created during learning Go.

# Purpose of this repo
Only has one purpose: track my progress with GO. This repo will not have any other purpose than just contain dummy/demo files, projects or other resources.
This might contain something useful, might not. Who knows...

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