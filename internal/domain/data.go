package domain

type Cryptocurrencies struct {
	Name   string `db:"name,omitempty"`
	Image  string `db:"image_link,omitempty"`
	Course int    `db:"course,omitempty"`
	Symbol string `db:"symbol,omitempty"`
}

type Wallet struct {
	Name  string `db:"name"`
	Image string `db:"image_link"`
}

type Networks struct {
	Name  string `db:"name"`
	Image string `db:"image_link"`
}
