package tripadvisor

type Places struct {
	Data []struct {
		LocationID string `json:"location_id"`
		Name       string `json:"name"`
		AddressObj struct {
			AddressString string `json:"address_string"`
		} `json:"address_obj"`
	} `json:"data"`
}

type Details struct {
	LocationID string `json:"location_id"`
	Name       string `json:"name"`
	WebURL     string `json:"web_url"`

	Rating string `json:"rating"`
}

type Photos struct {
	Data []struct {
		Images struct {
			// Thumbnail struct {
			// 	Height int    `json:"height"`
			// 	Width  int    `json:"width"`
			// 	URL    string `json:"url"`
			// } `json:"thumbnail"`
			// Small struct {
			// 	Height int    `json:"height"`
			// 	Width  int    `json:"width"`
			// 	URL    string `json:"url"`
			// } `json:"small"`
			// Medium struct {
			// 	Height int    `json:"height"`
			// 	Width  int    `json:"width"`
			// 	URL    string `json:"url"`
			// } `json:"medium"`
			// Large struct {
			// 	Height int    `json:"height"`
			// 	Width  int    `json:"width"`
			// 	URL    string `json:"url"`
			// } `json:"large"`
			Original struct {
				Height int    `json:"height"`
				Width  int    `json:"width"`
				URL    string `json:"url"`
			} `json:"original"`
		} `json:"images"`
	} `json:"data"`
}
