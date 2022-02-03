# rs
rs recursively traverse a directory

# Install 
```
go get -v GitHub.com/Cne3Rd/rs
```

# Usage

```
f := rs.Walk("c:\\users\\view")
for _, v := range f {
    fmt.Println(v)
}

```

