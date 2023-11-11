package stats

type GetStatsDB struct {
	SiteSafety string `json:"siteSafety"`
	Count      string `json:"count"`
}

type GetStatsResponse GetStatsDB

type GetStatsQuery struct {
	SiteUrl string `form:"siteUrl" binding:"required,http_url"`
}
