package response_handler

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

func Response(session *discordgo.Session, interaction *discordgo.InteractionCreate, msgFormat string) {
	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msgFormat,
		},
	})
}

func FileResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, filename string, file *os.File) {
	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Files: []*discordgo.File{
				&discordgo.File{
					Name:   filename,
					Reader: file,
				},
			},
		},
	})

}
