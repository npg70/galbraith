package main

import (
	"encoding/csv"
	"strings"
)

// takes csv and makes an HTML table
func CsvTableHTML(args []string, body string) string {
	out := strings.Builder{}
	out.WriteString("<div class=table-responsive style='max-width:768px'>")
	out.WriteString("<table class='table table-borderless table-sm small ms-3'>")
	r := csv.NewReader(strings.NewReader(body))

	/* optional caption that probably needs work */
	caption := strings.Join(args[1:], " ")
	if caption != "" {
		out.WriteString("<caption>" + caption + "</caption>")
	}
	// read header row
	row, _ := r.Read()
	out.WriteString("<thead><tr>")
	for _, col := range row {
		// note: bootstrap tables don't inherit colors
		out.WriteString("<th class='text-body'>" + strings.TrimSpace(col) + "</th>")
	}
	out.WriteString("</tr></thead>\n")

	for {
		row, err := r.Read()
		if err != nil {
			break
		}
		out.WriteString("<tr>")
		for _, col := range row {
			// note: bootstrap tables don't inherit colors
			out.WriteString("<td class='text-body'>" + col + "</td>")
		}
		out.WriteString("</tr>\n")
	}

	out.WriteString("</table>")
	out.WriteString("</div>\n")
	return out.String()
}
