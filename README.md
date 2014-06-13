### spark

##### Sparklines for Go, inspired by [holman/spark](https://github.com/holman/spark).

A quick example:

```go
import (
  "fmt"
  "github.com/joliv/spark"
)

boring_data := []float64{1, 2, 3, 4, 5, 6, 7, 8}

sparkline := spark.Spark(boring_data)

fmt.Println(sparkline)

> "▁▂▃▄▅▆▇█"

```

Grab it with ```go get github.com/joliv/spark```.

Now some more interesting examples.

Nats season batting averages at a certain point in their 2014 season:

```go
avgs := []float64{.270, .272, .293, .310, .274, .239, .237, .238, .111}
spark.Spark(avgs)

> "▇▇██▇▆▆▆▁"
```

Not too cool, but it is easy to see where the problem is in this lineup. You'll have to blame the [National League's rules](http://en.wikipedia.org/wiki/Designated_hitter) though, not Treinen. Anyway, have a look at average monthly highs in Phoenix:

```go
temps := []float64{67, 71, 77, 85, 95, 104, 106, 105, 100, 89, 76, 66}
spark.Spark(temps)

> "▁▂▃▄▆███▇▅▃▁"
```

Doesn't say much without knowing the min and max there (about 65° and 105°...why do people live there, again?) but you can clearly see the seasonal trend.

Zach Holman does a great sell too, and this is really just a port of [his neat tool](https://github.com/holman/spark). There are some cool examples there.
