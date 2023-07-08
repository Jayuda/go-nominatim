# go-telegram-message

Simple golang send message using apikey and telegram Id. 

```
package main

import(
    "github.com/Jayuda/go-telegram-message"
    "fmt"
)

func main() {
    ReverseData := telegram_message.Send("YourAPIKey", "Destination TelegramId Ex:-1981902810", "Your Message, Hello Boss")
    Address := ReverseData.Address
    fmt.Println(Address)
}
```
# go-telegram-message
