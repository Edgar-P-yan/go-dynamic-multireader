# Dynamic-MultiReader

Works as standard `io.MultiReader`, but instead of accepting an array of readers
`DynamicMultiReader` accepts `NewReaderFactory`, which should return new instances of `io.Reader`, or `nil`,
if no readers are remaining. This kind of `DynamicMultiReader` is needed when you get your source readers
not at once, but sequentially.

## Install

```sh
go get github.com/Edgar-P-yan/dynamic-multireader
```

## Example

For a better example look at the [./\_example/concat-chunked-files.go](./_example/concat-chunked-files.go).

```go
concatenatedReader := DynamicMultiReader(func(n int) io.Reader {
  if n + 1 <= chunksCount {
    return getReaderForChunk(n)
  } else {
    return nil
  }
})

// ...do something with you concatenated reader
```

## License

Released under [MIT License](./LICENSE).
