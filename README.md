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
    "time"

    "tele"
)

var mBot *tele.Bot // a bot instance

func main() {
    mBot, err := tele.Create("12345:ABCD-EFGH", "https://api.telegram.org/bot")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    for {
        updates := mBot.GetUpdates()
        for _, v := range updates.Result {
            if v.Message.Text == "hi" {
                // reply using markdown
                mBot.SendMarkdownMessage(v.Message.From.Id, "**hello** :D")
            } else {
                // reply using regular message
                mBot.SendMessage(v.Message.From.Id, "Telegram Bot API is awesome.")
            }
        }

        // a pause before next getUpdates request
        time.Sleep(2 * time.Second)
    }
}
```

## Credits

Copyright 2019 Maxim R. All rights reserved.

For feedback or questions &lt;epiqmax@gmail.com&gt;


[godoc]: http://godoc.org/github.com/epiqm/tele
[godoc-img]: https://godoc.org/github.com/epiqm/tele?status.svg
