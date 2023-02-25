package domain

const (
	// 3 levels of authentication
	//
	// High > Medium > Low
	Low = iota + 1
	Medium
	High
)

type AuthKey struct {
	Key   string `db:"key"`
	Name  string `db:"name"`
	Level int    `db:"level"`
}
