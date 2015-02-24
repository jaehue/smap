# shared-map
Built-in map doesn't support concurrent. This is concurrent map using channel, without mutex.

## Usage
Go get:
```
$ go get github.com/jaehue/smap
```

Import the package:
```
import (
    "github.com/jaehue/smap"
)
```

## Example
```
// Create new map
m := smap.New()

// Set item
m.Set("foo", "bar")

// Get item
t, ok := m.Get("foo")

// Check if item exists
if ok {
  bar := t.(string)
  fmt.Println("bar: ", bar)
}

// Remove item
m.Remove("foo")

// Count
fmt.Println("Count: ", m.Count())
```

## Test
Running tests:
```
$ go test github.com/jaehue/smap
```

## License
MIT (see [LICENSE](LICENSE) file)
