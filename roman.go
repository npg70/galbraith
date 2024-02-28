package main

// number to roman numeral. Doesn't go beyond 20, and doesn't handle errors.
// fine for intended purpose.

func toRoman(num int) string {
	return []string{"", "i", "ii", "iii", "iv", "v", "vi", "vii", "viii", "ix",
		"x", "xi", "xii", "xiii", "xiv", "xv", "xvi", "xvii", "xviii", "xix", "xx"}[num]
}
