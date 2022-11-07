# Constructors

In Go, we do not have constructors, just functions.

Most importantly: **Make the zero value useful**.

Examples of ready to use types:

* slices
* sync.WaitGroup
* bytes.Buffer
* strings.Builder
* sync.Mutex

You can use those right away.

Also, one might argue that struct initialization is preferred over a separate method, for example:

```go
var route = Route{
    Start: "Leipzig",
    Stop: "Berlin",
}

var route = NewRoute("Leipzig", "Berlin")
```

## Handling Options

> How do you handle options in Go?

Create a separate options type. This is a good idea, because it makes it easier
to add new options later.

Examples for various projects (search for "opts")

* https://pkg.go.dev/github.com/ph/moby/api/types#ImageSearchOptions
* https://github.com/hetznercloud/hcloud-go/blob/1fa5f34f79e8ecc5ec52d9b915e1ed47a911a1e2/hcloud/volume.go#L90-L95

## Functional Options Pattern

Another approach is the "functional options pattern", for example used here:

* https://github.com/hetznercloud/hcloud-go/blob/1fa5f34f79e8ecc5ec52d9b915e1ed47a911a1e2/hcloud/client.go#L137-L169

How it works:

* define a function type, that is your options (e.g. "ClientOptions")
* define functions that return the function type

Example:

```go
// WithToken configures a Client to use the specified token for authentication.
func WithToken(token string) ClientOption {
	return func(client *Client) {
		client.token = token
	}
}
```

* accept a variable number of these options in your "constructor"

Example usage:

```go
client := hcloud.NewClient(hcloud.WithToken("token"))
```

