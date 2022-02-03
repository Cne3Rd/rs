# rs
rs recursively traverse a directory

# Install 
```
go get -v github.com/Cne3Rd/rs
```

# Usage

```
f := rs.Walk("C:\\users\\views")
for _, v := range f {
    fmt.Println(v)
}

```

