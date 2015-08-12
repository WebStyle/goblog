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
        switch message.Text {
	        case "/start":
	          bot.SendMessage(message.Chat, "Hi, i'm Telegram bot. Let's start sign up please\n /setup - send my this text", nil)
          case "/help":
            bot.SendMessage(message.Chat, " Welcome to telegram chat!\n You can use this futures:\n  /online - see online users\n /share - share\n /search - search friends\n /rating - the must populer users\n /votemap  - votemap\n /settings - my profile settings\n /help - help\n have fun !", nil)
          case "/setup":
            bot.SendMessage(message.Chat, "Setup form", nil)
          case "/online":
            bot.SendMessage(message.Chat, "Online list", nil)
          case ":userName":
            bot.SendMessage(message.Chat, "Send message bor :userName", nil)
        }
    }
}
