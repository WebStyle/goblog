package main

import (
    "time"
    "github.com/tucnak/telebot"
    "fmt"
)

func main() {
    bot, err := telebot.NewBot("119744006:AAF5OAyhSl5pqs0E8cGNPM5rMTP7WJqPrQs")
    if err != nil {
        return
    }

    messages := make(chan telebot.Message)
    bot.Listen(messages, 1*time.Second)
    for message := range messages {
        fmt.Println(message.Sender)
        fmt.Println(message.Text)
        if message.Text == "/hi" {
            bot.SendMessage(message.Chat,
                "Hello, "+message.Sender.FirstName+"!", nil)
        }
    }
}
