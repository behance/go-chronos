#go-chronos

Go wrapper for the chronos API.


##Usage
See the examples directory for some more examples.

```
import (
  chronos "github.com/behance/go-chronos/chronos"
)

// chronos.NewDefaultConfig() has the url set to http://127.0.0.1:4400
config := chronos.Config{
  URL: "http://some-url:4400"
}

client, err := chronos.NewClient(config)

if err != nil {
  //handle however you want
}

jobs, err := client.Jobs()   // To get all jobs chronos knows about
...
```
