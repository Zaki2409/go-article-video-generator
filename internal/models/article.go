package models 

type Article struct {
	ID string `json:"id"`
	Title string `json:"titlr"`
	Content string `json:"content"`
}

type Summary struct {
	ID string `json:"id"`
	ArticleID string `json: "article_id"`
	Text string `json : "text"`
}

type VedioRequest struct {
	SummaryID string `json: "summary_id"`
	Title string `json : "title"`
	Text string `json : "text"`
}

type VedioResponse struct {
	ID string `json: "id"`
	VedioURL string `json : "vedio_url"`
}