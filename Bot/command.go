package bot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	MY_COMMAND_ID = "!"
)

func OpenBot(Token string) {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if len(m.Content) == 0 || m.Content[0:1] != MY_COMMAND_ID {
		return
	}

	contents := strings.Split(m.Content, " ")

	switch contents[0][1:] {
	case "edit_template":
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintln(
			"```edit_task\nhogehoge:testing\n```",
		))

	case "test":
		tasksinsert(contents)
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "pong":
		s.ChannelMessageSend(m.ChannelID, "ping!")

	//ヘルプ
	case "help":
		s.ChannelMessageSendComplex(m.ChannelID, helpcommand(contents))
	case "h":
		s.ChannelMessageSendComplex(m.ChannelID, helpcommand(contents))

		//タスク操作
	case "tlget":
		s.ChannelMessageSendComplex(m.ChannelID, tasklistget())
	case "tget":
		s.ChannelMessageSendComplex(m.ChannelID, taskget(contents))
	case "tdget":
		s.ChannelMessageSendComplex(m.ChannelID, taskdetailget(contents))
	case "tlpost":
		s.ChannelMessageSendComplex(m.ChannelID, tasklistinsert(contents))
	case "tpost":
		s.ChannelMessageSendComplex(m.ChannelID, tasksinsert(contents))
	case "tput":
		s.ChannelMessageSendComplex(m.ChannelID, tasksupdate(contents))
	case "tlput":
		s.ChannelMessageSendComplex(m.ChannelID, tasklistupdate(contents))
	case "tcomp":
		s.ChannelMessageSendComplex(m.ChannelID, taskcomplete(contents))
	case "tldel":
		s.ChannelMessageSendComplex(m.ChannelID, tasklistdelete(contents))
	case "tdel":
		s.ChannelMessageSendComplex(m.ChannelID, taskdelete(contents))
	default:
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "そんなコマンドないで:)",
				Description: "!help option[<command>]でコマンドを確認する",
			},
		})
	}

}
