# contextor
Contextor is a way to consistently store a value on a context.

# Usage
Start with a new `Contextor` value:

```go
var RequestPath contextor.New[string]("requestPath")
```

Every time you need to write to your context, you can use the "baked" `Contextor` to do the work:

```go
func AnnotateContextWithRequestPath(ctx contetxt.Context, request *http.Request) (context.Context, error){
	// apply value to the context:
	return RequestPath.Set(ctx, request.Path)
}
```

Retrieval is just as easy:

```go
func HandlePath(ctx context.Context, apiReq *APIRequest) (apiResp *APIResponse, error) {
	path, err := RequestPath.Get(ctx)
	if err != nil {
		return nil, err
	}
	
	return internalHandler(path, APIRequest)
}
```
