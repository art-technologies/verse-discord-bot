package discord

import (
	"fmt"
	"verse-discord-go/internal/verse"
)

type Sale struct {
	Title          string
	ArtistName     string
	CollectionName string
	ImageUrl       string
	NFTUrl         string
	Buyer          string
	BuyerUrl       string
	Seller         string
	SellerUrl      string
	Amount         string
	Type           string
}

func ToSale(a *verse.ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry) *Sale {
	sale := &Sale{
		Title:  *a.Artwork.Title,
		Amount: a.Amount,
		Type:   verse.SaleType(a.EntryType),
	}

	buyerUsername, buyerUrl := verse.GetBuyer(a)
	sellerUsername, SellerUrl := verse.GetSeller(a)

	sale.Buyer = buyerUsername
	sale.BuyerUrl = buyerUrl
	sale.Seller = sellerUsername
	sale.SellerUrl = SellerUrl

	switch asset := a.Asset.Asset.(type) {
	case *verse.ActivityPageActivityPageActivityEntryConnectionNodesActivityEntryAssetAssetEdition:
		sale.ArtistName = *asset.Artwork.Artist.Name
		sale.CollectionName = *asset.Artwork.Collection.Name
		sale.ImageUrl = verse.GetImageUrlFromStaticAsset(asset.StaticAsset)
		sale.NFTUrl = verse.NftUrl(a.Artwork.Id, fmt.Sprintf("%d", asset.EditionNumber))
	default:
		// do not print Bookmarks and other assets
		return nil
	}

	return sale
}
