package domain

type Category struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Preview string `db:"preview"`
	Created string `db:"created"`
}

type Quiz struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Preview string `db:"preview"`
}

type Option struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Preview  string `db:"preview"`
	Priority int    `db:"priority"`
}

type Translations struct {
	VarName string `db:"var_name"`
	En      string `db:"en"`
	Ru      string `db:"ru"`
	Es      string `db:"es"`
}
