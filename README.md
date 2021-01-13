Library to show reproducible memory leak in rust regex crate when using it in a ffi context with threads.

You will need musl target 

Build and run:

```
cargo build --lib --target x86_64-unknown-linux-musl
go run --ldflags "-linkmode external -extldflags '-static'" main.go 
```