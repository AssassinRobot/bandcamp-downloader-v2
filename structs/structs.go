package structs

type Mp3struct struct {
	Artist               string
	Title                string
	AlbumTitle           string
	ArtID                int
	Image                string
	AlbumArtwork         string
	BaseFilepath         string
	AlbumArtworkFilepath string
}

type TrackData struct {
	IsBandMember       interface{} `json:"is_band_member"`
	ClientIDSig        string      `json:"client_id_sig"`
	IsPreorder         bool        `json:"is_preorder"`
	AlbumIsPreorder    bool        `json:"album_is_preorder"`
	LicensedVersionIds interface{} `json:"licensed_version_ids"`
	FreeDownloadPage   interface{} `json:"freeDownloadPage"`
	HasVideo           interface{} `json:"has_video"`
	DefaultPrice       float64     `json:"defaultPrice"`
	Current            struct {
		Type                string      `json:"type"`
		SellingBandID       int         `json:"selling_band_id"`
		BandID              int         `json:"band_id"`
		AutoRepriced        interface{} `json:"auto_repriced"`
		SetPrice            float64     `json:"set_price"`
		DownloadPref        int         `json:"download_pref"`
		PublishDate         string      `json:"publish_date"`
		DownloadDescID      interface{} `json:"download_desc_id"`
		MinimumPriceNonzero float64     `json:"minimum_price_nonzero"`
		About               string      `json:"about"`
		PurchaseURL         interface{} `json:"purchase_url"`
		NewDescFormat       int         `json:"new_desc_format"`
		MinimumPrice        float64     `json:"minimum_price"`
		ModDate             string      `json:"mod_date"`
		Credits             string      `json:"credits"`
		Killed              interface{} `json:"killed"`
		PurchaseTitle       interface{} `json:"purchase_title"`
		ReleaseDate         string      `json:"release_date"`
		Artist              interface{} `json:"artist"`
		RequireEmail        interface{} `json:"require_email"`
		NewDate             string      `json:"new_date"`
		IsSetPrice          interface{} `json:"is_set_price"`
		RequireEmail0       interface{} `json:"require_email_0"`
		Private             interface{} `json:"private"`
		Title               string      `json:"title"`
		Audit               int         `json:"audit"`
		FeaturedTrackID     int         `json:"featured_track_id"`
		Upc                 string      `json:"upc"`
		ID                  int64       `json:"id"`
		ArtID               int         `json:"art_id"`
	} `json:"current"`
	TralbumSubscriberOnly bool        `json:"tralbum_subscriber_only"`
	LastSubscriptionItem  interface{} `json:"last_subscription_item"`
	HasAudio              bool        `json:"hasAudio"`
	InitialTrackNum       interface{} `json:"initial_track_num"`
	Trackinfo             []struct {
		VideoSourceType interface{} `json:"video_source_type"`
		Lyrics          interface{} `json:"lyrics"`
		File            struct {
			Mp3128 string `json:"mp3-128"`
		} `json:"file"`
		AltLink           interface{} `json:"alt_link"`
		HasInfo           bool        `json:"has_info"`
		VideoSourceID     interface{} `json:"video_source_id"`
		TrackNum          int         `json:"track_num"`
		EncodingError     interface{} `json:"encoding_error"`
		IsDownloadable    bool        `json:"is_downloadable"`
		LicenseType       int         `json:"license_type"`
		VideoID           interface{} `json:"video_id"`
		VideoMobileURL    interface{} `json:"video_mobile_url"`
		AlbumPreorder     bool        `json:"album_preorder"`
		TrackLicenseID    interface{} `json:"track_license_id"`
		EncodingPending   interface{} `json:"encoding_pending"`
		HasFreeDownload   interface{} `json:"has_free_download"`
		VideoPosterURL    interface{} `json:"video_poster_url"`
		UnreleasedTrack   bool        `json:"unreleased_track"`
		TrackID           int         `json:"track_id"`
		PlayCount         int         `json:"play_count"`
		FreeAlbumDownload bool        `json:"free_album_download"`
		Private           interface{} `json:"private"`
		EncodingsID       int64       `json:"encodings_id"`
		Title             string      `json:"title"`
		VideoCaption      interface{} `json:"video_caption"`
		TitleLink         string      `json:"title_link"`
		IsCapped          bool        `json:"is_capped"`
		IsDraft           bool        `json:"is_draft"`
		SizeofLyrics      int         `json:"sizeof_lyrics"`
		Duration          float64     `json:"duration"`
		ID                int         `json:"id"`
		VideoFeatured     interface{} `json:"video_featured"`
		Streaming         int         `json:"streaming"`
		HasLyrics         bool        `json:"has_lyrics"`
	} `json:"trackinfo"`
	PlayingFrom                string      `json:"playing_from"`
	PackageAssociatedLicenseID interface{} `json:"package_associated_license_id"`
	Artist                     string      `json:"artist"`
	AlbumReleaseDate           string      `json:"album_release_date"`
	ItemsPurchased             interface{} `json:"items_purchased"`
	PlayCapData                struct {
		StreamingLimit         int  `json:"streaming_limit"`
		StreamingLimitsEnabled bool `json:"streaming_limits_enabled"`
	} `json:"play_cap_data"`
	IsBonus          interface{} `json:"is_bonus"`
	ItemType         string      `json:"item_type"`
	UseExpandoLyrics bool        `json:"use_expando_lyrics"`
	IsPrivateStream  interface{} `json:"is_private_stream"`
	FREE             int         `json:"FREE"`
	Packages         []struct {
		AlbumID                 int64       `json:"album_id"`
		DownloadHasAudio        bool        `json:"download_has_audio"`
		SellingBandID           int         `json:"selling_band_id"`
		BandID                  int         `json:"band_id"`
		LiveEventID             interface{} `json:"live_event_id"`
		SubscriberOnly          interface{} `json:"subscriber_only"`
		SubscriberOnlyPublished bool        `json:"subscriber_only_published"`
		AlbumPrivate            interface{} `json:"album_private"`
		FulfillmentDays         int         `json:"fulfillment_days"`
		DownloadURL             string      `json:"download_url"`
		OptionsTitle            string      `json:"options_title"`
		Country                 interface{} `json:"country"`
		DescPt1                 string      `json:"desc_pt1"`
		DownloadArtID           int         `json:"download_art_id"`
		DownloadType            string      `json:"download_type"`
		LimitedCheckout         bool        `json:"limited_checkout"`
		AlbumReleaseDate        string      `json:"album_release_date"`
		AlbumPublishDate        string      `json:"album_publish_date"`
		AlbumArtist             interface{} `json:"album_artist"`
		AlbumTitle              string      `json:"album_title"`
		QuantitySold            interface{} `json:"quantity_sold"`
		LiveEventOver           interface{} `json:"live_event_over"`
		NewDescFormat           int         `json:"new_desc_format"`
		DescPt2                 interface{} `json:"desc_pt2"`
		TypeName                string      `json:"type_name"`
		QuantityAvailable       int         `json:"quantity_available"`
		DownloadArtist          string      `json:"download_artist"`
		Origins                 []struct {
			QuantitySold      interface{} `json:"quantity_sold"`
			OptionID          int         `json:"option_id"`
			QuantityAvailable int         `json:"quantity_available"`
			Quantity          interface{} `json:"quantity"`
			PackageID         int64       `json:"package_id"`
			ID                int         `json:"id"`
		} `json:"origins"`
		TypeID            int         `json:"type_id"`
		LiveEventTimezone interface{} `json:"live_event_timezone"`
		NewDate           string      `json:"new_date"`
		ReleaseDate       interface{} `json:"release_date"`
		DownloadID        int64       `json:"download_id"`
		Currency          string      `json:"currency"`
		Price             float64     `json:"price"`
		GridIndex         int         `json:"grid_index"`
		FeaturedDate      interface{} `json:"featured_date"`
		EditionSize       interface{} `json:"edition_size"`
		AlbumArt          interface{} `json:"album_art"`
		Options           []struct {
			QuantitySold interface{} `json:"quantity_sold"`
			Origins      []struct {
				QuantitySold      int   `json:"quantity_sold"`
				OptionID          int64 `json:"option_id"`
				QuantityAvailable int   `json:"quantity_available"`
				Quantity          int   `json:"quantity"`
				PackageID         int64 `json:"package_id"`
				ID                int   `json:"id"`
			} `json:"origins"`
			QuantityAvailable int         `json:"quantity_available"`
			Quantity          interface{} `json:"quantity"`
			Title             string      `json:"title"`
			Sku               string      `json:"sku"`
			Index             int         `json:"index"`
			ID                int64       `json:"id"`
		} `json:"options"`
		IsSetPrice              interface{} `json:"is_set_price"`
		AlbumArtID              int         `json:"album_art_id"`
		AlbumUpc                string      `json:"album_upc"`
		LiveEventReplaysEnabled interface{} `json:"live_event_replays_enabled"`
		IsLiveTicket            interface{} `json:"is_live_ticket"`
		Private                 interface{} `json:"private"`
		Title                   string      `json:"title"`
		URL                     string      `json:"url"`
		QuantityLimits          int         `json:"quantity_limits"`
		ShippingExceptionMode   interface{} `json:"shipping_exception_mode"`
		Label                   string      `json:"label"`
		Upc                     string      `json:"upc"`
		Sku                     string      `json:"sku"`
		Description             string      `json:"description"`
		DownloadTitle           string      `json:"download_title"`
		Arts                    []struct {
			Height   int    `json:"height"`
			FileName string `json:"file_name"`
			Width    int    `json:"width"`
			Crc      int    `json:"crc"`
			ImageID  int    `json:"image_id"`
			Index    int    `json:"index"`
			ID       int    `json:"id"`
		} `json:"arts"`
		TaxRate                     interface{} `json:"tax_rate"`
		LiveEventURL                interface{} `json:"live_event_url"`
		ID                          int64       `json:"id"`
		LiveEventScheduledStartDate interface{} `json:"live_event_scheduled_start_date"`
		AssociatedLicenseID         interface{} `json:"associated_license_id"`
		CertifiedSeller             int         `json:"certified_seller"`
		QuantityWarning             bool        `json:"quantity_warning"`
		URLForApp                   string      `json:"url_for_app"`
	} `json:"packages"`
	PreorderCount   interface{} `json:"preorder_count"`
	ForTheCurious   string      `json:"for the curious"`
	IsPurchased     interface{} `json:"is_purchased"`
	ID              int64       `json:"id"`
	FeaturedTrackID int         `json:"featured_track_id"`
	URL             string      `json:"url"`
	HasDiscounts    bool        `json:"has_discounts"`
	PAID            int         `json:"PAID"`
	ArtID           int         `json:"art_id"`
}
