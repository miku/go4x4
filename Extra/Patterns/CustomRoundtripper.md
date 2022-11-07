id: CustomRoundtripper
refs:
- https://medium.com/@_jesus_rafael/making-http-client-more-resilient-in-go-d24c66a64bd1
---

# Custom Roundtripper

Smaller interfaces can help writing alternative implementations quickly. The
[http.Client](https://golang.org/pkg/net/http/#Client) delegates the HTTP
transaction to a
[http.Roundtripper](https://golang.org/pkg/net/http/#RoundTripper), a single
method interface, which

> executes a single HTTP transaction, returning a Response for the provided Request.

A few alternative roundtrippers:

* [https://github.com/gregjones/httpcache](https://github.com/gregjones/httpcache);
  allows to cache HTTP in a mostly
[RFC7234](https://datatracker.ietf.org/doc/html/rfc7234) compliant way.  HTTP
caching can be implemented in a variety of ways, e.g. also outside the http
package. However, this approach uses an extension point in the library (the
http.RoundTripper interface) and implements a solution close to the problem.
Client code gets simpler, too: one only needs to change the client transport.




