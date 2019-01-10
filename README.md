# tele

[![Godoc Reference][godoc-img]][godoc]

A package for handling a telegram bot.

## Usage
``` $ go get github.com/epiqm/tele ```

### Example

Get updates and reply using a bot.

```go
// main.go
package main

import (
    "fmt"

    "github.com/epiqm/tele"
)

func main() {
    bot, err := tele.Create("123:bot-token", "https://api.telegram.org/bot", 0)

    go func() {
        for {
            updates := bot.GetUpdates(bot.LastUpdateId)
            for _, v := range updates.Result {
                if v.Message.Text == "hi" {
                    tele.SendMarkdownMessage(v.Message.From.Id, "hello :D")
                }
            }
            time.Sleep(2 * time.Second)
        }
    }()
}
```

## Credits

Copyright 2019 Maxim R. All rights reserved.

For feedback or questions &lt;epiqmax@gmail.com&gt;


[godoc]: http://godoc.org/github.com/epiqm/tele
[godoc-img]: https://godoc.org/github.com/epiqm/tele?status.svg
