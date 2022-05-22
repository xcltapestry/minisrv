# minsrv

MiniSrv is a thin wrapper around Go's HTTP server

## http server:
```go
func main() {
    minisrv.NewHTTPServer().
    AddRoute(route).
    AddMiddleware(middleware).
    ListenAndServe() // or ListenAndServe(":8082")
}
```
### Route
```go
// mux route
func route(m *mux.Router) {
    m.HandleFunc("/", indexHandler)
    m.HandleFunc("/health", healthHandler)
    m.HandleFunc("/api/v1/actid/%d", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Welcome to the actid page!")
    })
}
```

### Middeleware

```go
func middleware(n *negroni.Negroni) {
    n.Use(negroni.HandlerFunc(Authorizer))
    n.Use(negroni.HandlerFunc(APIMiddleware))
}
```

## Thanks

| package                           | type    |
|-----------------------------------|---------|
| https://github.com/gorilla/mux    | Route     |
| https://github.com/urfave/negroni | Middeleware |
