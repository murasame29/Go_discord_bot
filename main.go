package main

import (
	bot "discord_bot/Bot"
	"flag"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	bot.OpenBot(Token)
}

func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.MessageComponentData().CustomID == "hellouser" {
		resp := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: fmt.Sprintf("みんなのアイドルGopherだよ"),
			},
		}
		if err := s.InteractionRespond(i.Interaction, resp); err != nil {
			log.Fatalln(err)
		}

	}

}
