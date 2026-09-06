module github.com/npg70/galbraith

go 1.26.2

require (
	github.com/client9/cmdconfig v0.0.0-20260122131647-96eb45f98417
	github.com/client9/mdrec v0.0.0-20260831124222-17e6915a600a
	github.com/client9/ssg v0.0.0-20250629033452-a2320abf5d10
	github.com/client9/ssg/render/shortcode v0.0.0-00010101000000-000000000000
	github.com/client9/tagfunctions v0.0.0-20250629035250-66a093778e25
	golang.org/x/net v0.49.0
	golang.org/x/text v0.33.0
)

require (
	github.com/clipperhouse/displaywidth v0.6.2 // indirect
	github.com/clipperhouse/stringish v0.1.1 // indirect
	github.com/clipperhouse/uax29/v2 v2.3.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.19 // indirect
	github.com/olekukonko/cat v0.0.0-20250911104152-50322a0618f6 // indirect
	github.com/olekukonko/errors v1.1.0 // indirect
	github.com/olekukonko/ll v0.1.4-0.20260115111900-9e59c2286df0 // indirect
	github.com/olekukonko/tablewriter v1.1.3 // indirect
	github.com/yuin/goldmark/v2 v2.0.0 // indirect
	golang.org/x/sys v0.40.0 // indirect
)

replace github.com/client9/mdrec => ../mdrec/

replace github.com/client9/ssg/render/shortcode => ../ssg/render/shortcode
