package sub

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

func WakeupCycle(userID string, currentVoice string, circle string, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	mutex.Lock()
	defer mutex.Unlock()
	circleINT, err := strconv.Atoi(circle)
	channels, err := s.GuildChannels(i.GuildID)
	if err != nil {
		return fmt.Errorf("failed to fetch guild state: %s", err)
	}
	var voiceChannels = []*discordgo.Channel{}
	for a := 0; a < circleINT*2; a++ {
		for _, channel := range channels {
			if channel.Type == discordgo.ChannelTypeGuildVoice && channel.ID != currentVoice {
				voiceChannels = append(voiceChannels, channel)
			}
		}
	}
	for _, vc := range voiceChannels {

		err = s.GuildMemberMove(i.GuildID, userID, &vc.ID)
		if err != nil {
			return fmt.Errorf("failed to move user: %s", err)
		}
	}

	err = s.GuildMemberMove(i.GuildID, userID, &currentVoice)
	if err != nil {
		return fmt.Errorf("failed to move user: %s", err)
	}

	return nil
}
