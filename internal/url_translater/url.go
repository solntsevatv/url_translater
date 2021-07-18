package url_translater

type URL struct {
	Id       int
	LongUrl  string
	ShortURL string
}

type ShortURL struct {
	Id      int    `json:"-" db:"id"`
	LinkUrl string `json:"url" db:"short_url" binding:"required"`
}

type LongURL struct {
	Id      int    `json:"-" db:"id"`
	LinkUrl string `json:"url" db:"long_url" binding:"required"`
}
