# Transformerインタフェース

https://pkg.go.dev/golang.org/x/text/transform#Transformer

```go
type Transformer interface {
	Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error)
	Reset()
}
```

transform.Chainを用いて複数のTransformerを連結することで、複数の変換を一度に行うことができる。

```go
t := transform.Chain(transform.Nop, transform.Discard)
```

encodingパッケージのEncoder, Decoderは、Transformerインタフェースを実装している。