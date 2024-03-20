package domain

type Category struct {
	ID      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Preview string `db:"preview" json:"preview"`
	Created string `db:"created" json:"created"`
}

type Quiz struct {
	ID      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Preview string `db:"preview" json:"preview"`
}

type Option struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Preview  string `db:"preview" json:"preview"`
	Priority int    `db:"priority" json:"priority"`
}

type Lang string

type Translations struct {
	VarName string `db:"var_name"`
	En      Lang   `db:"en"`
	Ru      Lang   `db:"ru"`
	Es      Lang   `db:"es"`
}

const (
	RU Lang = "ru"
	EN Lang = "en"
	ES Lang = "es"
)
