package verse

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

const rootVerseUrl = "https://verse.works"
const openSeaUrl = "https://opensea.io"

func ImageUrlFromS3Key(key string) string {
	parts := strings.Split(key, "/")
	for i, part := range parts {
		parts[i] = url.PathEscape(part)
	}
	encodedKey := strings.Join(parts, "/")
	encodedURL := fmt.Sprintf("%s/image/w640/%s@jpeg", rootVerseUrl, encodedKey)
	return encodedURL
}

func NftUrl(artworkId, editionNumber string) string {
	return fmt.Sprintf("%s/artworks/%s/%s", rootVerseUrl, artworkId, editionNumber)
}

func GetBuyer(a *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry) (string, string) {
	if a.ToUserId != nil {
		return *a.ToUserId.Username, fmt.Sprintf("%s/@%s", rootVerseUrl, *a.ToUserId.Username)
	}

	if a.ToEthereumAddress != nil {
		return a.ToEthereumAddress.Address, fmt.Sprintf("%s/%s", openSeaUrl, a.ToEthereumAddress.Address)
	}

	return "", ""
}

func GetSeller(a *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry) (string, string) {
	if a.FromUserId != nil {
		return *a.FromUserId.Username, fmt.Sprintf("%s/@%s", rootVerseUrl, *a.FromUserId.Username)
	}

	if a.FromEthereumAddress != nil {
		return a.FromEthereumAddress.Address, fmt.Sprintf("%s/%s", openSeaUrl, a.FromEthereumAddress.Address)
	}

	return "", ""
}

func SaleType(a ActivityEntryType) string {
	switch a {
	case ActivityEntryTypePmSale:
		return "PRIMARY"
	default:
		return "SECONDARY"
	}
}

func GetImageUrlFromStaticAsset(s *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntryAssetAssetEditionStaticAsset) string {
	switch staticAsset := (*s).(type) {
	case *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntryAssetAssetEditionStaticAssetImageAsset:
		return ImageUrlFromS3Key(*staticAsset.S3Key)
	case *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntryAssetAssetEditionStaticAssetVideoAsset:
		return ImageUrlFromS3Key(staticAsset.PreviewS3Key)
	case *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntryAssetAssetEditionStaticAssetIFrameAsset:
		return ImageUrlFromS3Key(staticAsset.PreviewS3Key)
	case *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntryAssetAssetEditionStaticAssetSVGAsset:
		return ImageUrlFromS3Key(staticAsset.PreviewS3Key)
	default:
		return ""
	}
}

func GetMarketplace(a ActivityEntryType) string {
	switch a {
	case ActivityEntryTypePmSale, ActivityEntryTypeSmListed, ActivityEntryTypeSmDelisted, ActivityEntryTypeSmSale, ActivityEntryTypeSmOffer, ActivityEntryTypeSmGlobalOffer:
		return "Verse"
	case ActivityEntryTypeOsListed, ActivityEntryTypeOsSale, ActivityEntryTypeOsOffer, ActivityEntryTypeOsCollectionOffer:
		return "OpenSea"
	}
	return ""
}

func GetArtistName(a *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry) string {
	switch asset := a.Asset.Asset.(type) {
	case *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntryAssetAssetEdition:
		return *asset.Artwork.Artist.Name
	}
	return ""
}

func GetArtistSlug(a *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry) string {
	return a.Artwork.Artist.Slug
}

func ContainsCollaborators(a *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry, desiredCollaborators *[]string) bool {
	// passthrough all if not specified
	if desiredCollaborators == nil || len(*desiredCollaborators) == 0 {
		return true
	}

	projectCollaborators := make([]string, 0)

	if a.Asset == nil {
		return false
	}

	switch asset := a.Asset.Asset.(type) {
	case *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntryAssetAssetEdition:
		for _, person := range asset.Artwork.Collection.Persons {
			projectCollaborators = append(projectCollaborators, person.Slug)
		}
	}

	for _, desiredCollaborator := range *desiredCollaborators {
		if slices.Contains(projectCollaborators, desiredCollaborator) {
			return true
		}
	}

	return false
}

func ContainsArtist(a *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry, desiredArtists *[]string) bool {
	// passthrough all if not specified
	if desiredArtists == nil || len(*desiredArtists) == 0 {
		return true
	}

	artistName := GetArtistSlug(a)

	if slices.Contains(*desiredArtists, artistName) {
		return true
	}

	return false
}

func ContainsEvents(a *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry, desiredEvents *[]string) bool {
	// passthrough all if not specified
	if desiredEvents == nil || len(*desiredEvents) == 0 {
		return true
	}

	if slices.Contains(*desiredEvents, string(a.EntryType)) {
		return true
	}

	return false
}

func ContainsCollections(a *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntry, desiredCollections *[]string) bool {
	// passthrough all if not specified
	if desiredCollections == nil || len(*desiredCollections) == 0 {
		return true
	}

	if a.Asset == nil {
		return false
	}

	switch asset := a.Asset.Asset.(type) {
	case *ActivityPageActivityPageActivityEntryConnectionNodesActivityEntryAssetAssetEdition:
		if slices.Contains(*desiredCollections, *&asset.Artwork.Collection.Slug) {
			return true
		}
	}

	return false
}
