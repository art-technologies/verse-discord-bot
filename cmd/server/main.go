package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"verse-discord-go/internal/discord"
	"verse-discord-go/internal/verse"

	"github.com/joho/godotenv"
)

//go:embed guilds.json
var guildsData []byte

func main() {
	err := godotenv.Load()

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		panic("BOT_TOKEN is required")
	}

	apiUrl := os.Getenv("API")
	if botToken == "" {
		panic("API is required")
	}

	var guilds verse.Guilds
	err = json.Unmarshal(guildsData, &guilds)
	if err != nil {
		panic(err)
	}

	verseService := verse.NewService(&verse.Config{
		API: apiUrl,
	})

	discordService, err := discord.NewDiscordService(&discord.Config{
		Token: botToken,
	})
	if err != nil {
		panic(err)
	}

	activities, err := verseService.GetActivity()
	if err != nil {
		panic(err)
	}

	// New activity comes in
	for _, activity := range activities {
		for _, guild := range guilds {
			// satisfy events filter
			if guild.Filters != nil && !verse.ContainsEvents(&activity, guild.Filters.Events) {
				continue
			}

			// satisfy artists filter
			if guild.Filters != nil && !verse.ContainsArtist(&activity, guild.Filters.Artists) {
				continue
			}

			// satisfy collaborators filter
			if guild.Filters != nil && !verse.ContainsCollaborators(&activity, guild.Filters.Collaborators) {
				continue
			}

			// satisfy collections filter
			if guild.Filters != nil && !verse.ContainsCollections(&activity, guild.Filters.Collections) {
				continue
			}

			processActivity(discordService, guild.ChannelID, activity)
		}
	}
}

func processActivity(discordService *discord.Service, channelId string, activity verse.ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry) error {
	switch activity.EntryType {
	case verse.ActivityEntryTypePmSale, verse.ActivityEntryTypeSmSale, verse.ActivityEntryTypeOsSale:
		sale := discord.ToSale(&activity)
		if sale != nil {
			return discordService.SaleEvent(context.Background(), channelId, sale)
		}
	case verse.ActivityEntryTypeSmListed, verse.ActivityEntryTypeOsListed:
		listing := discord.ToListing(&activity)
		if listing != nil {
			return discordService.ListingEvent(context.Background(), channelId, listing)
		}

	case verse.ActivityEntryTypeSmOffer, verse.ActivityEntryTypeSmGlobalOffer, verse.ActivityEntryTypeOsOffer:
		listing := discord.ToListing(&activity)
		if listing != nil {
			return discordService.ListingEvent(context.Background(), channelId, listing)
		}
	default:
		// ignore other types of activities for now
		fmt.Println(activity.EntryType)
	}

	return nil
}
