KBot - A Simple Telegram Bot

you can interact with  the bot by clicking the link below
https://t.me/kbotprometheus_bot

KBot is a simple Telegram bot created using the Golang programming language and the telebot package. It can perform a few basic tasks such as greeting the user, providing the current time, and providing limited help and support.
Features

    Greets the user with a personalized message.
    Provides the current time in the user's timezone.
    Provides limited help and support for the available commands.

Getting Started

To use the KBot Telegram bot, you will need to follow these steps:
Install go, go get github.com/gopkg.in/telebot.v3 or find the binary kbot in the repository and start ./kbot start it also depends on TG_BOT_TOKEN do not foget to setup it yourself

1.Clone this repository to your local machine.
cd kbot 

2.Set the TG_BOT_TOKEN environment variable to the token of your Telegram bot. You can obtain a token by creating a new bot using the BotFather bot in Telegram. 
read -s TG_BOT_TOKEN(ctr shift v, ctr v, command v), export TG_BOT_TOKEN

3.Run the following command to build the binary:

in the root directory of your bartaadalbert repository you need your
go build -ldflags "-X="github.com/bartaadalbert/kbot/cmd.appVersion=v1.0.3
appVersion=v1.0.3 in my example you can set any strin there
./kbot start

Usage

To interact with the KBot Telegram bot, you can send it messages using the Telegram app. The following commands are available:

    /start - Starts the bot and greets the user with a personalized message.
    time, now, час, data, teper, koli, Яка година, година, Дата - Provides the current time in the user's timezone.
    hello, hi, hey, szia, mizu, привіт, здравствуйте, здрастуйте, privet, zdrastvuy, zdrastvuyte - Greets the user with a personalized message.
    help, support, допомога, підтримка, suport, dopomoha, pomozi meni, sco robiti, what to do, how, i cant - Provides limited help and support for the available commands.
    thx, thanks, thank you, by, poka, dyakuyu, Дякую, Пока - Ends the conversation with the bot.

License

This project is licensed under the MIT License - see the LICENSE file for details.
