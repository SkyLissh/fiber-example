package schemas

type Url struct {
	Url string `json:"url"`
	TargetUrl string `json:"target_url" validate:"required,url"`
	IsActive bool `json:"is_active"`
	Clicks int `json:"clicks"`
}
