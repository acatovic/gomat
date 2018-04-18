# Gomat
Package __gomat__ is a simple matrix implemenation in Go. At its core is the `Matrix` struct, which simply wraps a slice of floats, i.e. `[]float64`. The package exposes a set of functions that create, operate on, and return `Matrix` structs. We make a light attempt at optimising cache locality; however it is by no means the most optimal implementation.

## Examples ##
### Adding two matrices ###
```go
import (
  "fmt"
  "github.com/acatovic/gomat"
)

ma := gomat.Randn(16, 4)
mb := gomat.Randn(16, 4)
mc := gomat.Add(ma, mb)
fmt.Println(mc)
```

### Dot product ###
```go
import (
  "fmt"
  "github.com/acatovic/gomat"
)

ma := gomat.New([][]float64{{1,2},
                            {3,4},
                            {5,6}})
mb := gomat.New([][]float64{{1,2},
                            {3,4}})
mc := gomat.Dot(ma, mb)
fmt.Println(mc)
```
