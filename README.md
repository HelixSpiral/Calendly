# Calendly
A Golang API wrapper for the Calendly platform

## Usage
Create a new client with `calendly.New()`

```go
cw := calendly.New("api-key-here")
```

Use any of the build-in functions to query against the Calendly API

```go
currentUser, err := cw.GetCurrentUser()
if err != nil {
    panic(err)
}

log.Printf("%+v\r\n", currentUser)
```