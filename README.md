KBot - A Simple Telegram Bot

## Continuous Integration/Continuous Deployment (CI/CD)

We use GitHub Actions for our CI/CD pipeline. Whenever a commit is made to the `develop` branch, the pipeline is triggered. The steps include:

1. Checkout the repository.
2. Install Go.
3. Login to the GitHub Container Registry.
4. Extract the repository name.
5. Build and push the image.
6. Install `jq` and `yq`, and perform SSH.
7. Checkout the repository.
8. Increment the version.
9. Modify `values.yaml`.
10. Commit and push the changes.

You can find the GitHub Actions workflow configuration file in `.github/workflows/cicd.yml`.

![CI/CD Workflow](.demo/CICD_pipeline.png)


you can interact with  the bot by clicking the link below
https://t.me/kbotprometheus_bot

This Makefile provides an easy way to build and automate the creation of binaries for go app with both amd64 and arm64 architectures. The Makefile can run tests on the app and check the code for errors before deploying the binary. The Makefile has defaults to create amd64 binaries with make linux, make macos, and make windows. If the architecture is arm, you can create armlinux with make armlinux and armmacos with make armmacos.

Additionally, the Makefile can build Docker images with binaries like make imagelinux, make imagemacos, and make imagewindows. For arm architecture, you can build imagearmlinux with make imagearmlinux and imagearmmacos with make imagearmmacos.

The Makefile can also clean all binaries in the directory and created images. You can push to Docker Hub and save to images with a name path locally. The Makefile requires settings such as app name, version from git tag or randomly generated version, the registry, target architecture, and IMAGE_BUILDER.

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
