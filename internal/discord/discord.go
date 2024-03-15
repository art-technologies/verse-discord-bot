package discord

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"slices"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Token string
}

type Service struct {
	config *Config
	client *discordgo.Session
	logger *slog.Logger
}

func NewDiscordService(config *Config) (*Service, error) {
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	dg.Identify.Intents = discordgo.IntentGuildMessages

	return &Service{
		config: config,
		client: dg,
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}, nil
}

func (s *Service) SaleEvent(ctx context.Context, channelID string, sale *Sale) error {
	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Buyer",
			Value:  fmt.Sprintf("[%s](%s)", sale.Buyer, sale.BuyerUrl),
			Inline: false,
		},
		{
			Name:   "Price",
			Value:  "$" + sale.Amount,
			Inline: true,
		},
		{
			Name:   "Type",
			Value:  sale.Type,
			Inline: true,
		},
	}

	// insert seller if exists
	if sale.Seller != "" {
		fields = slices.Insert(fields, 1, &discordgo.MessageEmbedField{
			Name:   "Seller",
			Value:  fmt.Sprintf("[%s](%s)", sale.Seller, sale.SellerUrl),
			Inline: false,
		})
	}

	data := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title: fmt.Sprintf("%s by %s", sale.Title, sale.ArtistName),
			URL:   sale.NFTUrl,
			Image: &discordgo.MessageEmbedImage{
				URL: sale.ImageUrl,
			},
			Fields: fields,
		},
		Components: []discordgo.MessageComponent{
			&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.Button{
						Style:    discordgo.LinkButton,
						Label:    "View Artwork",
						URL:      sale.NFTUrl,
						Disabled: false,
						Emoji:    discordgo.ComponentEmoji{Name: "🖼️"},
					},
					&discordgo.Button{
						Style:    discordgo.LinkButton,
						Label:    sale.Buyer,
						URL:      sale.BuyerUrl,
						Disabled: false,
						Emoji:    discordgo.ComponentEmoji{Name: "🖼️"},
					},
				},
			},
		},
	}

	return s.sendMessage(ctx, channelID, data)
}

func (s *Service) ListingEvent(ctx context.Context, channelID string, listing *Listing) error {
	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Offerer",
			Value:  fmt.Sprintf("[%s](%s)", listing.Seller, listing.SellerUrl),
			Inline: false,
		},
		{
			Name:   "Price",
			Value:  "$" + listing.Amount,
			Inline: true,
		},
		{
			Name:   "Marketplace",
			Value:  listing.Marketplace,
			Inline: true,
		},
	}

	data := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title: fmt.Sprintf("%s by %s", listing.Title, listing.ArtistName),
			URL:   listing.NFTUrl,
			Image: &discordgo.MessageEmbedImage{
				URL: listing.ImageUrl,
			},
			Fields: fields,
		},
		Components: []discordgo.MessageComponent{
			&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.Button{
						Style:    discordgo.LinkButton,
						Label:    "View Artwork",
						URL:      listing.NFTUrl,
						Disabled: false,
						Emoji:    discordgo.ComponentEmoji{Name: "🖼️"},
					},
					&discordgo.Button{
						Style:    discordgo.LinkButton,
						Label:    listing.Seller,
						URL:      listing.SellerUrl,
						Disabled: false,
						Emoji:    discordgo.ComponentEmoji{Name: "🖼️"},
					},
				},
			},
		},
	}

	return s.sendMessage(ctx, channelID, data)
}

func (s *Service) OfferEvent(ctx context.Context, channelID string, listing *Offer) error {
	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Buyer",
			Value:  fmt.Sprintf("[%s](%s)", listing.Buyer, listing.BuyerUrl),
			Inline: false,
		},
		{
			Name:   "Price",
			Value:  "$" + listing.Amount,
			Inline: true,
		},
		{
			Name:   "Marketplace",
			Value:  listing.Marketplace,
			Inline: true,
		},
	}

	data := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title: fmt.Sprintf("%s by %s", listing.Title, listing.ArtistName),
			URL:   listing.NFTUrl,
			Image: &discordgo.MessageEmbedImage{
				URL: listing.ImageUrl,
			},
			Fields: fields,
		},
		Components: []discordgo.MessageComponent{
			&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.Button{
						Style:    discordgo.LinkButton,
						Label:    "View Artwork",
						URL:      listing.NFTUrl,
						Disabled: false,
						Emoji:    discordgo.ComponentEmoji{Name: "🖼️"},
					},
					&discordgo.Button{
						Style:    discordgo.LinkButton,
						Label:    listing.Buyer,
						URL:      listing.BuyerUrl,
						Disabled: false,
						Emoji:    discordgo.ComponentEmoji{Name: "🖼️"},
					},
				},
			},
		},
	}

	return s.sendMessage(ctx, channelID, data)
}

func (s *Service) sendMessage(ctx context.Context, channelID string, data *discordgo.MessageSend) error {
	_, err := s.client.ChannelMessageSendComplex(channelID, data)
	if err != nil {
		s.logger.Error("sendMessage", err)
		return err
	}
	return nil
}
