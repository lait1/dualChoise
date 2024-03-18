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

type Translations struct {
	VarName string `db:"var_name"`
	En      string `db:"en"`
	Ru      string `db:"ru"`
	Es      string `db:"es"`
}
