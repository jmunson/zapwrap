#ZapWrap

***ZapWrap*** is a very simple wrapper that allows you to use a [Zap](http://github.com/uber-go/zap) logger with the [Echo](https://github.com/labstack/echo) framework, or anything else expecting a similar logging interface.



Neither Zap nor Echov2 have stable APIs at this point so this may break and should probably not be used in production at this time. This code itself is also not very well tested either, but it is pretty trivial.



#Example


```go
package main

import (
        "github.com/jmunson/zapwrap"
        "github.com/labstack/echo"
        "github.com/uber-go/zap"
)q

func main() {
        e := echo.New()
        log := zap.NewJSON()
        e.SetLogger(zapwrap.Wrap(log))
        e.Logger().Info("Hello!")
        //continue to use both echo and your zap logger as normal
}

```
