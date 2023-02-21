package main

import (
	"encoding/csv"
	"strings"
)

// takes csv and makes an HTML table
func CsvTableHTML(args []string, body string) string {
	out := strings.Builder{}
	out.WriteString("<table class=xtable>")
	r := csv.NewReader(strings.NewReader(body))

	// read header row
	row, _ := r.Read()
	out.WriteString("<tr>")
	for _, col := range row {
		out.WriteString("<th>" + strings.TrimSpace(col) + "</th>")
	}
	out.WriteString("</tr>\n")

	for {
		row, err := r.Read()
		if err != nil {
			break
		}
		out.WriteString("<tr>")
		for _, col := range row {
			out.WriteString("<td>" + col + "</td>")
		}
		out.WriteString("</tr>\n")
	}

	out.WriteString("</table>\n")
	return out.String()
}
