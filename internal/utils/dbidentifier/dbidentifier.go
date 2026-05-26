package dbidentifier

// ToGoIdentifier converts database identifier to Go identifier.
func ToGoIdentifier(databaseIdentifier string) string { _ = "STUB: not implemented"; return "" }

// ToGoFileName converts database identifier to Go file name.
func ToGoFileName(databaseIdentifier string) string { _ = "STUB: not implemented"; return "" }

func replaceInvalidChars(identifier string) string { _ = "STUB: not implemented"; return "" }

func needsCharReplacement(identifier string) (increase int, needs bool) {
	_ = "STUB: not implemented"
	return 0, false
}

var asciiCharacterReplacement = map[rune]string{
	'!':  "exclamation",
	'"':  "quotation",
	'#':  "number",
	'$':  "dollar",
	'%':  "percent",
	'&':  "ampersand",
	'\'': "apostrophe",
	'(':  "opening_parentheses",
	')':  "closing_parentheses",
	'*':  "asterisk",
	'+':  "plus",
	',':  "comma",
	'-':  "_",
	'.':  "_",
	'/':  "slash",
	':':  "colon",
	';':  "semicolon",
	'<':  "less",
	'=':  "equal",
	'>':  "greater",
	'?':  "question",
	'@':  "at",
	'[':  "opening_bracket",
	'\\': "backslash",
	']':  "closing_bracket",
	'^':  "caret",
	'`':  "accent",
	'{':  "opening_braces",
	'|':  "vertical_bar",
	'}':  "closing_braces",
	'~':  "tilde",
}
