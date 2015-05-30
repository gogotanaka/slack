# slack/notifier

```go
package main

import "github.com/gogotanaka/slack/notifier"

func main() {
  notifier := notifier.New("https://hooks.slack.com/services/XXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
  notifier.Username = "gogotanaka"

  notifier.Ping("Hi!")
  notifier.Ping("Hi!Hi!")

  notifier.IconEmoji = ":poop:"
  notifier.Username = "poop"
  notifier.Ping("I'm poop!")
}
```

![Screen Shot 2015-05-30 at 4.11.19 AM.png](https://qiita-image-store.s3.amazonaws.com/0/30440/9d870116-4459-2c79-5012-cf1de51a2659.png)


### params
```
Username  
IconEmoji
IconUrl
Channel
```
