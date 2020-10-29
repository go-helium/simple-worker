# Helium simple worker implementation

## Example

```go
package main

import (
    "context"
    "time"

    worker "github.com/go-helium/simple-worker""time"
)

func runner(context.Context) {
    // Do something
}

func main() {
    worker.WrapJob("test", runner,
        worker.WithImmediately(), // run at start
        worker.WithTimer(time.Second * 10), // repeat every 10 seconds
    )
}
```