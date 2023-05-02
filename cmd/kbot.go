/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	TG_BOT_TOKEN = os.Getenv("TG_BOT_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started", appVersion)
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TG_BOT_TOKEN,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatal("Please check your TG_BOT_TOKEN env variable %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Printf("Message payload and text %s: %s", m.Message().Payload, m.Text())

			// payload := m.Message().Payload
			// log.Printf("Payload: %s", payload)
			// log.Printf("Text: %s", text)
			// I need JSON analize the payload
			jsonData, err := json.Marshal(m.Message())
			if err != nil {
				log.Printf("Failed to marshal message to JSON: %s", err)
				return err
			}

			// Write the JSON data to file
			err = ioutil.WriteFile("message.json", jsonData, 0644)
			if err != nil {
				log.Printf("Failed to write JSON data to file: %s", err)
				return err
			}

			// Print the JSON data
			log.Printf("Message data: %s", string(jsonData))

			//commands for bot usage are handeled here
			validCommands := map[string]string{
				"hello": "Hello %s!, I am kbot!",
				"time":  "The current time in your timezone is %s",
				"help":  "I am here to assist you with limited commands %s",
				"thx":   "Have a nice day, I was glad to help you.",
			}

			//Make arrays of handle text from user
			helloArray := strings.Split("hi,hello,hey,szia,mizu,привіт,здравствуйте,здрастуйте,privet,zdrastvuy,zdrastvuyte", ",")
			timeArray := strings.Split("time,now,час,data,teper,koli,Яка година,година,Дата", ",")
			helpArray := strings.Split("help,support,допомога,підтримка,suport,dopomoha,pomozi meni,sco robiti,what to do,how,i cant", ",")
			greatingsArray := strings.Split("thanx,thx,ok,thank you,by,poka,dyakuyu,thanks,thank,Дякую,Пока", ",")

			//get the text from user
			text := strings.ToLower(m.Text())

			switch text {
			case "/start":
				return m.Send(fmt.Sprintf("Hello, %s! I am kbot! What can i help you", m.Sender().FirstName))

			default:
				// Check matches any of the elements in the time array
				for _, t := range timeArray {
					if strings.Contains(strings.ToLower(m.Text()), t) {
						// Get the use time
						loc, err := time.LoadLocation(m.Sender().LanguageCode)
						if err != nil {
							loc = time.UTC
						}
						// Current time
						now := time.Now().In(loc).Format("2002-02-02 22:22:22")
						// Send the current time
						return m.Send(fmt.Sprintf(validCommands["time"], now))
					}
				}

				// Check matches any of the elements in the hello array
				for _, v := range helloArray {
					if strings.Contains(strings.ToLower(m.Text()), v) {
						return m.Send(fmt.Sprintf(validCommands["hello"], m.Sender().FirstName))
					}
				}
				// Check matches any of the elements in the help array
				for _, h := range helpArray {
					if strings.Contains(strings.ToLower(m.Text()), h) {
						keys := make([]string, 0, len(validCommands))
						for k := range validCommands {
							keys = append(keys, k)
						}
						return m.Send(fmt.Sprintf(validCommands["help"], keys))
					}
				}
				// Check matches any of the elements in the help array
				for _, g := range greatingsArray {
					if strings.Contains(strings.ToLower(m.Text()), g) {
						return m.Send(fmt.Sprintf(validCommands["thx"]))
					}
				}

				return m.Send("I apologize for being boring. Is there anything specific you would like me to help you with")
			}

			return err
		})
		kbot.Start()
	},
}

// function to check if my string array contains this  string slic
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
