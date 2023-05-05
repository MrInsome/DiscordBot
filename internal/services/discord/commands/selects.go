package commands

import (
	"github.com/bwmarrin/discordgo"
)

func (c *Commands) Selects(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse
	switch i.ApplicationCommandData().Options[0].Name {
	case "single":
		response = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Now let's take a look on selects. This is single item select menu.",
				Flags:   discordgo.MessageFlagsEphemeral,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.SelectMenu{
								// Select menu, as other components, must have a customID, so we set it to this value.
								CustomID:    "select",
								Placeholder: "Choose your favorite programming language 👇",
								Options: []discordgo.SelectMenuOption{
									{
										Label: "Go",
										// As with components, this things must have their own unique "id" to identify which is which.
										// In this case such id is Value field.
										Value: "go",
										Emoji: discordgo.ComponentEmoji{
											Name: "🦦",
										},
										// You can also make it a default option, but in this case we won't.
										Default:     false,
										Description: "Go programming language",
									},
									{
										Label: "JS",
										Value: "js",
										Emoji: discordgo.ComponentEmoji{
											Name: "🟨",
										},
										Description: "JavaScript programming language",
									},
									{
										Label: "Python",
										Value: "py",
										Emoji: discordgo.ComponentEmoji{
											Name: "🐍",
										},
										Description: "Python programming language",
									},
								},
							},
						},
					},
				},
			},
		}
	case "multi":
		minValues := 1
		response = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Now let's see how the multi-item select menu works: " +
					"try generating your own stackoverflow search link",
				Flags: discordgo.MessageFlagsEphemeral,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.SelectMenu{
								CustomID:    "stackoverflow_tags",
								Placeholder: "Select tags to search on StackOverflow",
								// This is where confusion comes from. If you don't specify these things you will get single item select.
								// These fields control the minimum and maximum amount of selected items.
								MinValues: &minValues,
								MaxValues: 3,
								Options: []discordgo.SelectMenuOption{
									{
										Label:       "Go",
										Description: "Simple yet powerful programming language",
										Value:       "go",
										// Default works the same for multi-select menus.
										Default: false,
										Emoji: discordgo.ComponentEmoji{
											Name: "🦦",
										},
									},
									{
										Label:       "JS",
										Description: "Multiparadigm OOP language",
										Value:       "javascript",
										Emoji: discordgo.ComponentEmoji{
											Name: "🟨",
										},
									},
									{
										Label:       "Python",
										Description: "OOP prototyping programming language",
										Value:       "python",
										Emoji: discordgo.ComponentEmoji{
											Name: "🐍",
										},
									},
									{
										Label:       "Web",
										Description: "Web related technologies",
										Value:       "web",
										Emoji: discordgo.ComponentEmoji{
											Name: "🌐",
										},
									},
									{
										Label:       "Desktop",
										Description: "Desktop applications",
										Value:       "desktop",
										Emoji: discordgo.ComponentEmoji{
											Name: "💻",
										},
									},
								},
							},
						},
					},
				},
			},
		}
	case "auto-populated":
		response = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "The tastiest things are left for the end. Meet auto populated select menus.\n" +
					"By setting `MenuType` on the select menu you can tell Discord to automatically populate the menu with entities of your choice: roles, members, channels. Try one below.",
				Flags: discordgo.MessageFlagsEphemeral,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.SelectMenu{
								MenuType:     discordgo.UserSelectMenu,
								CustomID:     "user_select",
								Placeholder:  "Pick a user from the voice channels!",
								MaxValues:    1,
								ChannelTypes: []discordgo.ChannelType{discordgo.ChannelTypeGuildVoice},
								//Options: ,
							},
						},
					},
				},
			},
		}
	}
	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		panic(err)
	}
}

//func getUsersFromSameVoice(s *discordgo.Session, i *discordgo.InteractionCreate) []discordgo.SelectMenuOption{
//
//	// Find the voice channel that the user is in.
//	voiceState, err := s.State.VoiceState(i.GuildID, i.Member.User.ID)
//
//	if err != nil {
//		log.Println(err)
//		return nil
//	}
//
//	// Find all the users in the same voice channel.
//	usersInVoice := make(map[string][]string)
//
//	for _, vs := range guild.VoiceStates {
//		if vs.ChannelID != "" {
//			if _, ok := usersInVoice[vs.ChannelID]; !ok {
//				usersInVoice[vs.ChannelID] = []string{}
//			}
//			usersInVoice[vs.ChannelID] = append(usersInVoice[vs.ChannelID], vs.UserID)
//		}
//	}
//
//	// Create the select menu with the list of users.
//	selectMenu := discordgo.SelectMenu{
//		CustomID:    "user_select",
//		Placeholder: "From the voice channel!",
//		MaxValues:   1,
//		Options:     []discordgo.SelectMenuOption{},
//	}
//	for _, user := range users {
//		option := discordgo.SelectMenuOption{
//			Label: user.Username,
//			Value: user.ID,
//		}
//		selectMenu.Options = append(selectMenu.Options, option)
//	}
//	return selectMenu.Options
//}
