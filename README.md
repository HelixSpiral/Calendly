# Calendly
A Golang API wrapper for the Calendly platform

## Usage
Create a new client with `calendly.New()`, advanced usage for custom API url and headers can be seen in the unit tests.

```go
cw := calendly.New(&calendly.CalendlyWrapperInput{
    ApiKey: "api-key-here",
})
```

Use any of the build-in functions to query against the Calendly API

```go
currentUser, err := cw.GetCurrentUser()
if err != nil {
    panic(err)
}

log.Printf("%+v\r\n", currentUser)
```