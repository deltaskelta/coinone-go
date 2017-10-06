### Coinone API client in Go

Basic Public and Private Coinone API endpoints library.

```go
import coinone "github.com/deltaskelta/coineone-go"

func main() {
    coinone, err := coinone.NewAPI(YOUR_API_KEY, YOUR_SECRET_KEY)
    if err != nil {
        panic(err)
    }

    // use the endpoints as needed...
}

```

### Contributing

In order to run the tests you will need to define an `apiKey` and `secretKey` somewhere in
the package
