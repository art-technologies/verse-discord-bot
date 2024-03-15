package discord

import (
	"fmt"
	"verse-discord-go/internal/verse"
)

type Listing struct {
	Title          string
	ArtistName     string
	CollectionName string
	ImageUrl       string
	NFTUrl         string
	Seller         string
	SellerUrl      string
	Amount         string
	Marketplace    string
}

func ToListing(a *verse.ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry) *Listing {
	listing := &Listing{
		Title:       *a.Artwork.Title,
		Amount:      a.Amount,
		Marketplace: verse.GetMarketplace(a.EntryType),
	}

	sellerUsername, SellerUrl := verse.GetSeller(a)

	listing.Seller = sellerUsername
	listing.SellerUrl = SellerUrl

	switch asset := a.Asset.Asset.(type) {
	case *verse.ActivityPageActivityPageActivityEntryConnectionNodesActivityEntryAssetAssetEdition:
		listing.ArtistName = *asset.Artwork.Artist.Name
		listing.CollectionName = *asset.Artwork.Collection.Name
		listing.ImageUrl = verse.GetImageUrlFromStaticAsset(asset.StaticAsset)
		listing.NFTUrl = verse.NftUrl(a.Artwork.Id, fmt.Sprintf("%d", asset.EditionNumber))
	default:
		// do not print Bookmarks and other assets
		return nil
	}

	return listing
}
