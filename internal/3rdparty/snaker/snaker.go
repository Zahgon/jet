package snaker

// Package snaker provides methods to convert CamelCase names to snake_case and back.
// It considers the list of allowed initialsms used by github.com/golang/lint/golint (e.g. ID or HTTP)

// CamelToSnake converts a given string to snake case
func CamelToSnake(s string) string { _ = "STUB: not implemented"; return "" }

// append the last word

// SnakeToCamel returns a string converted from snake case to uppercase
func SnakeToCamel(s string, firstLetterUppercase ...bool) string {
	_ = "STUB: not implemented"
	return ""
}

func snakeToCamel(s string, upperCase bool) string { _ = "STUB: not implemented"; return "" }

// lowerCase and i == 0

// startsWithInitialism returns the initialism if the given string begins with it
func startsWithInitialism(s string) string { _ = "STUB: not implemented"; return "" }

// the longest initialism is 5 char, the shortest 2

func toLowerFirstLetter(s string) string { _ = "STUB: not implemented"; return "" }

func camelizeWord(word string, force bool) string { _ = "STUB: not implemented"; return "" }

// already camelCase

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/206c0f020eba0f7fbcfbc467a5eb808037df2ed6/lint.go#L731
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"ETA":   true,
	"GPU":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"OS":    true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
	"OAuth": true,
}

// add exceptions here for things that are not automatically convertable
var snakeToCamelExceptions = map[string]string{
	"oauth": "OAuth",
}
