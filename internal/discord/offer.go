package discord

import (
	"fmt"
	"verse-discord-go/internal/verse"
)

type Offer struct {
	Title          string
	ArtistName     string
	CollectionName string
	ImageUrl       string
	NFTUrl         string
	Buyer          string
	BuyerUrl       string
	Amount         string
	Marketplace    string
}

func ToOffer(a *verse.ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry) *Offer {
	offer := &Offer{
		Title:       *a.Artwork.Title,
		Amount:      a.Amount,
		Marketplace: verse.GetMarketplace(a.EntryType),
	}

	buyerUsername, buyerUrl := verse.GetBuyer(a)

	offer.Buyer = buyerUsername
	offer.BuyerUrl = buyerUrl

	switch asset := a.Asset.Asset.(type) {
	case *verse.ActivityPageActivityPageActivityEntryConnectionNodesActivityEntryAssetAssetEdition:
		offer.ArtistName = *asset.Artwork.Artist.Name
		offer.CollectionName = *asset.Artwork.Collection.Name
		offer.ImageUrl = verse.GetImageUrlFromStaticAsset(asset.StaticAsset)
		offer.NFTUrl = verse.NftUrl(a.Artwork.Id, fmt.Sprintf("%d", asset.EditionNumber))
	default:
		// do not print Bookmarks and other assets
		return nil
	}

	return offer
}
