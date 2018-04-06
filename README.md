# go-zenefits
A Go client library for accessing the [Zenefits API](https://developers.zenefits.com/docs/getting-started)

Go version 1.7 or greater

## Usage

```go
package main

import (
	"context"
	"flag"

	"github.com/rucas/go-zenefits/zenefits"
	"golang.org/x/oauth2"
)

func main() {
	id := flag.Int("id", 0, "Zenefits Company Id")
	k := flag.String("key", "", "Zenefits Api Key")
	flag.Parse()
	
	ctx := context.Background()
  	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *k})
  	tc := oauth2.NewClient(ctx, ts)
  	client := zenefits.NewClient(tc)
}
```

See [examples/](example) for more code samples.

## Contributing

:wave: :point_right: Check out the [Contributing](CONTRIBUTING.md) doc to get you started.

## License

This library is distributed under the MIT license found in the [LICENSE](LICENSE) file
