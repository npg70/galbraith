package main

func toRoman(num int) string {
	return []string{"", "i", "ii", "iii", "iv", "v", "vi", "vii", "viii", "ix",
		"x", "xi", "xii", "xiii", "xiv", "xv"}[num]
}
