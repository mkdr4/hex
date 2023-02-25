package domain

type Payment struct {
	Id             string `db:"id"`
	Store          string `db:"store"`
	Signature      string `db:"signature"`
	Fiat           string `db:"fiat"`
	FiatCurrency   int    `db:"fiat_currency"`
	Crypto         string `db:"crypto"`
	CryptoCurrency int    `db:"crypto_currency"`
	PayDescription string `db:"pay_description"`
	RedirectUrl    string `db:"redirect_url"`
	WebHookUrl     string `db:"webhook_url"`
	ExpirationTime string `db:"expiration_time"`
	Status         string `db:"status"`
}
