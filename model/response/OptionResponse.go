package response

type OptionResponse struct {
	Key     string `json:"OptionKey" form:"OptionKey"`
	Content string `json:"OptionContent" form:"OptionContent"`
}
