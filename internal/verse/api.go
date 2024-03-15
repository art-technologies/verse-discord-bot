package verse

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"verse-discord-go/internal/helpers"

	"github.com/Khan/genqlient/graphql"
)

const maxPages = 1

type Config struct {
	API string
}

type Service struct {
	config *Config
	logger *slog.Logger
	client *graphql.Client
}

func NewService(config *Config) *Service {
	client := graphql.NewClient(config.API, http.DefaultClient)

	return &Service{
		config: config,
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		client: &client,
	}
}

func (s *Service) GetRecentActivityPage(cursor *string) (*ActivityPageResponse, error) {
	filter := &ActivityPageFilter{
		AssetType: helpers.Ptr(AssetTypeEdition),
		ActivityType: []ActivityEntryType{
			ActivityEntryTypePmSale,
			ActivityEntryTypeSmSale,
			ActivityEntryTypeOsSale,

			ActivityEntryTypeSmOffer,
			ActivityEntryTypeSmGlobalOffer,
			ActivityEntryTypeOsOffer,

			ActivityEntryTypeSmListed,
			ActivityEntryTypeOsListed,
		},
	}
	return ActivityPage(context.Background(), *s.client, cursor, filter)
}

func (s *Service) GetActivity() ([]ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry, error) {
	var cursor *string = nil

	pagesCount := 0

	for {
		println("Getting page", pagesCount)
		activityPage, err := s.GetRecentActivityPage(cursor)
		if err != nil {
			return nil, err
		}

		if !activityPage.ActivityPage.PageInfo.HasNextPage {
			return activityPage.ActivityPage.Nodes, nil
		}

		cursor = activityPage.ActivityPage.PageInfo.EndCursor

		if pagesCount >= maxPages {
			return activityPage.ActivityPage.Nodes, nil
		}
		pagesCount += 1
	}
}
