package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Ref struct {
	ID   string
	Body string
}

func importFootnoteID(p *Person, ref string) string {

	if ref == "" {
		return ""
	}
	if strings.HasPrefix(ref, "/") {
		return ref
	}
	return "/" + p.ID + "/" + ref
}

func makeFootnoteID(p *Person, ref string) string {
	if ref == "" {
		return ""
	}
	if strings.HasPrefix(ref, "/") {
		return ref
	}
	//
	return ref
	//return "/" + p.ID + "/" + ref
}
func makeFootnoteRef(p *Person, ref string) string {
	if ref == "" {
		return "" // no reference
	}

	// if ref is already an absolute ID
	if strings.HasPrefix(ref, "/") {
		return "$ref[" + ref + "]"
	}

	return "$ref[" + makeFootnoteID(p, ref) + "]"
}

func makeFootnote(p *Person, id string, body string) string {
	if id == "" {
		return ""
	}
	return "$footnote[" + makeFootnoteID(p, id) + "]{" + body + "}\n"
}

func makeFootnotes(p *Person) string {
	if len(p.Footnotes) == 0 {
		return ""
	}
	s := strings.Builder{}
	s.WriteString("$footnotes{\n")
	for id, body := range p.Footnotes {
		s.WriteString(makeFootnote(p, id, body))
	}
	s.WriteString("}\n")
	return s.String()
}

type Footnotes map[string]string

func (n Footnotes) UnmarshalText(text []byte) error {
	scan := NewScanner(text)
	for {
		args, body, err := scan.Next()
		if args == nil && err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if len(args) != 1 {
			return fmt.Errorf("Inside Ref.%s expected one arg, got %v", args[0], args)
		}
		n[args[0]] = body
	}
}

type Event struct {
	Type     string
	Date     Date
	Location string
	Name     string
	Ref      string
	Note     string
}

func (e *Event) UnmarshalText(text []byte) error {
	scan := NewScanner(text)
	for {
		args, _, err := scan.Next()
		if args == nil && err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if len(args) != 2 {
			return fmt.Errorf("Inside Event.%s expected two args, got %v", args[0], args)
		}
		switch args[0] {
		case "date":
			e.Date.UnmarshalText([]byte(args[1]))
		case "location":
			e.Location = args[1]
		case "ref":
			e.Ref = args[1]
		case "note":
			e.Note = args[1]
		case "Name":
			e.Note = args[1]
		default:
			return fmt.Errorf("unknown Event comment %v", args)
		}
	}
}

func (e Event) DatePlace() string {
	return fmt.Sprintf("%s at %s", e.Date, e.Location)
}

type SourceLink struct {
	order    int // for ordering
	RecordID string
	URL      string
	Name     string // for display
}

type ExternalSource map[string]SourceLink

func (e ExternalSource) List() []SourceLink {
	out := make([]SourceLink, len(e))
	i := 0
	for _, val := range e {
		out[i] = val
		i++
	}
	//TODO Sort by ordering
	return out
}

func (e ExternalSource) UnmarshalText(text []byte) error {
	scan := NewScanner(text)
	for {
		args, _, err := scan.Next()
		if args == nil && err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if len(args) != 2 {
			return fmt.Errorf("Expected two args, got %v", args)
		}
		id := strings.ToLower(args[0])
		switch id {
		case "familysearch":
			e[id] = SourceLink{
				order:    10,
				RecordID: args[1],
				URL:      "https://www.familysearch.org/tree/person/details/" + args[1],
				Name:     "FamilySearch",
			}
		case "findagrave":
			e[id] = SourceLink{
				order:    20,
				RecordID: args[1],
				URL:      "https://www.findagrave.com/memorial/" + args[1],
				Name:     "Find a Grave",
			}
		case "billiongraves":
			e[id] = SourceLink{
				order:    50,
				RecordID: args[1],
				URL:      "https://billiongraves.com/grave/" + args[1],
				Name:     "Billion Graves",
			}
		case "wikitree":
			e[id] = SourceLink{
				order:    40,
				RecordID: args[1],
				URL:      "https://www.wikitree.com/wiki/" + args[1],
				Name:     "WikiTree",
			}
		case "geni":
			e[id] = SourceLink{
				order:    60,
				RecordID: args[1],
				URL:      "https://www.geni.com/people/" + args[1],
				Name:     "Geni",
			}
		case "ancestry":
			e[id] = SourceLink{
				order:    30,
				RecordID: args[1],
				URL:      "https://www.ancestry.com/family-tree/person/" + strings.Trim(args[1], "/") + "/",
				Name:     "Ancestry",
			}
		default:
			return fmt.Errorf("Unknown external source: %q", args[0])
		}
	}
}

type Person struct {
	parent     *Person
	generation int      // generation number 1,2,3..
	counter    int      // entry counter 1...
	refs       []string // footnote coutner per person
	lastUpdate time.Time

	root Root // not clear why I need this

	ID        string
	Name      []string
	Events    map[string]*Event
	Partners  []*Person
	Children  []*Person
	Footnotes Footnotes
	External  ExternalSource
	Intro     string
	Body      []string
	Notes     []string
	Todos     []string
	Tags      []string
}

func (p *Person) getRef(s string) int {
	for i, val := range p.refs {
		if s == val {
			return i + 1
		}
	}
	return 0
}
func (p *Person) addRef(s string) int {
	if idx := p.getRef(s); idx != 0 {
		return idx
	}
	p.refs = append(p.refs, s)
	return len(p.refs)
}

func (p *Person) FullName() string {
	return strings.Join(p.Name, " ")
}
func (p *Person) FirstName() string {
	return p.Name[0]
}
func (p *Person) LastName() string {
	return p.Name[len(p.Name)-1]
}
func (p *Person) UnmarshalText(text []byte) error {
	scan := NewScanner(text)
	for {
		args, body, err := scan.Next()
		if args == nil && err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if len(args) == 0 {
			log.Fatalf("Person function had 0 args: body = %q", string(text))
		}
		switch args[0] {
		case "tags":
			p.Tags = args[1:]
		case "external":
			e := make(ExternalSource)
			if err := e.UnmarshalText([]byte(body)); err != nil {
				return err
			}
			p.External = e
		case "footnotes":
			fn := make(Footnotes)
			if err := fn.UnmarshalText([]byte(body)); err != nil {
				return err
			}
			p.Footnotes = fn
		case "child":
			// child personid reference
			if len(args) == 2 && body == "" {
				newp, err := p.root.Load(args[1])
				if err != nil {
					return err
				}
				p.Children = append(p.Children, newp)
			} else {
				newp := &Person{root: p.root}
				if err := newp.UnmarshalText([]byte(body)); err != nil {
					return err
				}
				p.Children = append(p.Children, newp)
			}
		case "partner":
			newp := &Person{root: p.root}
			if err := newp.UnmarshalText([]byte(body)); err != nil {
				return err
			}
			p.Partners = append(p.Partners, newp)
		case "name":
			p.Name = args[1:]
		case "birth", "baptism", "death", "burial", "marriage":
			if p.Events == nil {
				p.Events = make(map[string]*Event)
			}
			e := &Event{}

			// inline version
			if len(args) > 1 {
				fs := flag.NewFlagSet("Event "+args[0], flag.ExitOnError)
				fs.TextVar(&e.Date, "date", &e.Date, "")
				fs.StringVar(&e.Location, "location", "", "")
				fs.StringVar(&e.Ref, "ref", "", "")
				fs.StringVar(&e.Note, "note", "", "")
				fs.StringVar(&e.Name, "name", "", "")
				if err := fs.Parse(args[1:]); err != nil {
					return err
				}
			}
			// Normal embedded form
			if len(body) > 0 {
				if err := e.UnmarshalText([]byte(body)); err != nil {
					return err
				}
			}
			p.Events[args[0]] = e
		case "intro":
			p.Intro = strings.TrimSpace(body)
		case "body":
			p.Body = append(p.Body, strings.Join(args[1:], " "))
			if body != "" {
				p.Body = append(p.Body, body)
			}
		case "note":
			p.Notes = append(p.Notes, body)
		case "todo":
			p.Todos = append(p.Todos, body)
		default:
			return fmt.Errorf("Unknown command %q", args[0])
		}
	}
}

type Root map[string]*Person

func (r Root) Load(id string) (*Person, error) {
	val, ok := r[id]
	if ok {
		return val, nil
	}
	newp := &Person{
		root: r,
		ID:   id,
	}

	sourcedir := "people"
	filesuffix := ".sh"

	fname := id
	if !strings.HasPrefix(id, sourcedir) {
		fname = filepath.Join(sourcedir, id)
	}
	if !strings.HasSuffix(id, filesuffix) {
		fname += filesuffix
	}
	log.Printf("reading %q\n", fname)
	bstring, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	if err = newp.UnmarshalText(bstring); err != nil {
		return nil, err
	}
	/*
		f2 := make(Footnotes)
		for id, body := range newp.Footnotes {
			f2[id] = body
		}
		newp.Footnotes = f2
	*/
	r[id] = newp
	return newp, nil
}

func EventDatePlace(p *Person, e *Event) string {
	out := fmt.Sprintf("$date{%s}", e.Date.String())

	if e.Location != "" {
		out += " at " + TitleCompound(e.Location)
	}
	if e.Ref != "" {
		out += makeFootnoteRef(p, e.Ref)
	}
	return out
}

func EventDatePlaceCompact(p *Person, e *Event, ignorePlace string) string {
	parts := []string{}
	if e.Date.String() != "" {
		parts = append(parts, "$date{"+e.Date.String()+"}")
	}
	if e.Location != "" && !strings.EqualFold(e.Location, ignorePlace) {
		parts = append(parts, TitleCompound(e.Location))
	}
	if e.Note != "" {
		parts = append(parts, e.Note)
	}
	return strings.Join(parts, ", ") + makeFootnoteRef(p, e.Ref)
}

func WriteLineageNameLink(p *Person) string {
	return fmt.Sprintf("$child-link[%s]{%s}", p.ID, WriteLineageName(p))
}

func WriteLineageName(p *Person) string {
	return fmt.Sprintf("$lineage-name{%s$gen{%d} %s}",
		p.FirstName(), p.generation, p.LastName())
}
func WritePrimaryName(p *Person) string {
	return fmt.Sprintf("$primary-name{%s$gen{%d} %s}",
		p.FirstName(), p.generation, p.LastName())
}

func WriteChildName(name string) string {
	return "$child-name{" + name + "}"
}

func WritePersonLink(p *Person) string {
	return fmt.Sprintf("$child-link[%s]{%s}", p.ID, p.FullName())
}

// child link to new page/entry/person
func WriteChildLink(w io.StringWriter, p *Person) {
	val := fmt.Sprintf("$child-link[%s]{%s}", p.ID, p.FullName())
	w.WriteString(WriteChildName(val))

	/*
		w.WriteString(WriteChildName(fmt.Sprintf("<child-link href=%q>%s</child-link>",
			p.ID, p.FullName())))
	*/
}
func WriteChildBioInline(w io.StringWriter, parent *Person, child *Person, ignorePlace string) {
	// no linkage, print mini-bio
	w.WriteString(WriteChildName(child.FullName()))
	birth := child.Events["birth"]
	bp := child.Events["baptism"]
	if birth != nil && bp == nil {
		// since this child doesn't have it's own page,
		// the footnote is attached to the parent, not the child
		w.WriteString(", b.&nbsp;" + EventDatePlaceCompact(parent, birth, ignorePlace))
	}
	if birth == nil && bp != nil {
		// since this child doesn't have it's own page,
		// the footnote is attached to the parent, not the child
		w.WriteString(", bp.&nbsp;" + EventDatePlaceCompact(parent, bp, ignorePlace))
	}
	if birth != nil && bp != nil {
		// both specified.. it gets complicated
		// if both same location -- do something smart
		// if both have date -- do something smart

		w.WriteString(", b.&nbsp;" + EventDatePlaceCompact(parent, birth, ignorePlace))
		w.WriteString(", bp.&nbsp;" + EventDatePlaceCompact(parent, bp, ignorePlace))
	}
	death := child.Events["death"]
	if death != nil {
		w.WriteString("; d.&nbsp;" + EventDatePlaceCompact(parent, death, ""))
	}
	if len(child.Partners) == 1 {
		// ; m. name
		w.WriteString("; m.&nbsp;" + WriteChildPartnerName(child.Partners[0]))
	} else {
		// ; m.(1) name; m.(2) name
		for i, p := range child.Partners {
			w.WriteString(fmt.Sprintf("; m.(%d)&nbsp;%s", i+1, WriteChildPartnerName(p)))
		}
	}
	w.WriteString(".")
	for i, b := range child.Body {
		child.Body[i] = strings.TrimSpace(b)
	}
	body := strings.Join(child.Body, "\n\n")
	if len(body) > 0 {
		w.WriteString("\n\n" + body)
	}
}

// Mini child bio with link to new page/person/entry
func WriteChildBioLinked(w io.StringWriter, parent *Person, child *Person, ignorePlace string) {
	// link, print basics
	WriteChildLink(w, child)
	birth := child.Events["birth"]
	bp := child.Events["baptism"]
	if birth != nil && bp == nil {
		oldRef := birth.Ref
		birth.Ref = importFootnoteID(child, oldRef)

		// here the footnote is actually in the child page
		w.WriteString(", b.&nbsp;" + EventDatePlaceCompact(child, birth, ignorePlace))

		if oldRef != "" {
			// if we have reference copy it into the parent
			parent.Footnotes[birth.Ref] = child.Footnotes[oldRef]
			birth.Ref = oldRef
		}
	}

	// if ONLY baptism
	if birth == nil && bp != nil {
		oldref := bp.Ref
		// here the footnote is actually in the child page
		w.WriteString(",&nbsp;bp. " + EventDatePlaceCompact(child, bp, ignorePlace))

		// if we have reference copy it into the parent
		if oldref != "" {
			// if we have reference copy it into the parent
			if parent.Footnotes == nil {
				parent.Footnotes = make(Footnotes)
			}
			parent.Footnotes[bp.Ref] = child.Footnotes[oldref]
			bp.Ref = oldref
		}
	}

	// both birth and baptism
	if birth != nil && bp != nil {
		// both specified.. it gets complicated
		// if both same location -- do something smart
		// if both have date -- do something smart

		w.WriteString(", b.&nbsp;" + EventDatePlaceCompact(parent, birth, ignorePlace))

		// hack -- assume baptism has the reference
		oldref := bp.Ref
		// here the footnote is actually in the child page
		w.WriteString(",&nbsp;bp. " + EventDatePlaceCompact(child, bp, ignorePlace))

		// if we have reference copy it into the parent
		if oldref != "" {
			// if we have reference copy it into the parent
			parent.Footnotes[bp.Ref] = child.Footnotes[oldref]
			bp.Ref = oldref
		}
	}
	if len(child.Partners) == 1 {
		// ; m. name
		w.WriteString("; m.&nbsp;" + WriteChildPartnerName(child.Partners[0]))
	} else {
		// ; m.(1) name; m.(2) name
		for i, p := range child.Partners {
			w.WriteString(fmt.Sprintf("; m.(%d)&nbsp;%s", i+1, WriteChildPartnerName(p)))
		}
	}
	w.WriteString(".")
}

func WriteChild(w io.StringWriter, parent *Person, child *Person, birthOrder int, common string) {
	w.WriteString(fmt.Sprintf("$child[counter=%d generation=%d birth-order=%d]{",
		child.counter,
		child.generation,
		birthOrder))
	if child.ID == "" {
		WriteChildBioInline(w, parent, child, common)
	} else {
		WriteChildBioLinked(w, parent, child, common)
	}
	w.WriteString("}\n")
}

func WriteChildrenIntro(w io.StringWriter, first *Person, partner *Person, common string, ctype string) {
	w.WriteString("$children-intro{")
	w.WriteString(fmt.Sprintf("%s of %s and %s (%s) %s",
		ChildPlural(len(partner.Children)),
		first.FirstName(),
		partner.FirstName(),
		partner.LastName(),
		first.LastName()))
	if common != "" && ctype != "" {
		w.WriteString(", " + ctype)
	}
	w.WriteString(":")
	w.WriteString("}\n")
}

func WriteChildPartnerName(p *Person) string {
	return "$child-partner-name{" + p.FullName() + "}"
}

func WriteTitle(p *Person) string {
	out := p.FullName()

	e := p.Events["birth"]
	if e == nil || e.Date.year == 0 {
		e = p.Events["baptism"]
	}
	if e != nil && e.Date.year != 0 {
		about := ""
		if e.Date.qualifier != DATE_EXACT {
			about = "~"
		}
		out += fmt.Sprintf(" b. %s%d", about, e.Date.year)
	}
	for _, spouse := range p.Partners {
		out += " m. " + spouse.FullName()
	}
	return out
}

// writeTitleBlock emits a block in the form of
// Full Name b. YEAR m. LastName
func WriteTitleBlock(p *Person) string {
	out := "$front{"

	out += "$title{"
	out += WriteTitle(p)
	out += "}\n"

	out += "}\n"
	return out
}

func (r Root) LoadAll(root string) {
	first, err := r.Load(root)
	if err != nil {
		panic(err)
	}

	// counter is the "numbered person" counter in registry style
	// TODO total count of all people/children in XXX style

	counter := 1
	first.counter = counter
	first.generation = 1
	q := []*Person{first}

	for len(q) > 0 {
		first = q[0]
		q = q[1:]

		// generate next generation
		for _, p := range first.Partners {
			for _, c := range p.Children {
				if c.ID != "" {
					kid, err := r.Load(c.ID)
					if err != nil {
						panic(err)
					}
					counter++
					kid.parent = first
					kid.counter = counter
					kid.generation = first.generation + 1
					q = append(q, kid)
				}
			}
		}
	}
}

func FindMother(child *Person) *Person {
	p := child.parent
	// we know the parent (likely father).. but need to find the right spouse
	for _, spouse := range p.Partners {
		for _, c := range spouse.Children {
			if c.ID == child.ID {
				return spouse
			}
		}
	}
	return nil
}

func (r Root) loadOne(primary string) []string {
	first, err := r.Load(primary)
	if err != nil {
		log.Fatalf("Error loading %s: %s", primary, err)
	}

	cid := []string{}
	// generate next generation
	for _, p := range first.Partners {
		for _, c := range p.Children {
			if c.ID != "" {
				cid = append(cid, c.ID)
			}
		}
	}
	return cid
}

func (r Root) generateOne(primary string) (string, []string) {
	first, err := r.Load(primary)
	if err != nil {
		panic(err)
	}

	out := &strings.Builder{}
	out.WriteString(WriteTitleBlock(first))

	out.WriteString(fmt.Sprintf("$person[id=%s generation=%d counter=%d]{",
		primary, first.generation, first.counter))

	out.WriteString("$banner{")
	out.WriteString("$headline{" + WriteTitle(first) + "}\n")
	out.WriteString("}\n")

	out.WriteString("$person-body{")
	out.WriteString("$person-secondary{")
	out.WriteString("$externals{")
	for _, val := range first.External.List() {
		out.WriteString("$external[" + val.URL + "]{" + val.Name + "}\n")
	}
	out.WriteString("}\n")

	out.WriteString("$lineage{")
	for p := first; p != nil; p = p.parent {
		if p.parent == nil {
			break
		}
		mother := FindMother(p)

		// this isn't quite right yet.
		// might want generation number here and to break apart father's first and last names.

		out.WriteString(fmt.Sprintf("$ancestor[counter=%d generation=%d mother=%q]{%s}\n",
			p.parent.counter, p.generation, mother.FullName(), WriteLineageNameLink(p.parent)))
	}
	out.WriteString("}\n") /* lineage */

	out.WriteString("$pagemeta{")
	out.WriteString(fmt.Sprintf("<a href=https://github.com/npg70/galbraith/blob/main/people/%s.sh>%s</a> updated on ", first.ID, "source"))
	out.WriteString("$date{" + ParseTime(first.lastUpdate).String() + "}")
	out.WriteString("}\n") /* page meta */

	out.WriteString("}\n")

	out.WriteString("$person-main{")

	out.WriteString("<div class='ms-5 me-5 mb-3'>")
	if len(first.Intro) > 0 {
		out.WriteString("$intro{" + first.Intro + "}\n")

	}
	if len(first.Tags) > 0 {
		out.WriteString("$tags[")
		for _, tag := range first.Tags {
			out.WriteString(fmt.Sprintf("%q ", tag))
		}
		out.WriteString("]\n")
	}
	out.WriteString("</div>\n")

	out.WriteString("$person-bio{<p>")
	out.WriteString(fmt.Sprintf("$primary-number{%d}", first.counter))
	out.WriteString(WritePrimaryName(first))
	birth := first.Events["birth"]
	if birth != nil {
		out.WriteString(" was born " + EventDatePlace(first, birth))
	}
	bapt := first.Events["baptism"]
	if bapt != nil {
		out.WriteString(" was baptized " + EventDatePlace(first, bapt))
	}
	out.WriteString(". ")
	death := first.Events["death"]
	if death != nil {
		out.WriteString("Died " + EventDatePlace(first, death))
	}
	out.WriteString(".")
	out.WriteString("</p>")
	for i, partner := range first.Partners {
		marriage := partner.Events["marriage"]
		if marriage != nil {
			ordinal := ""
			if len(first.Partners) > 1 {
				ordinal = Ordinal(i + 1)
			}
			out.WriteString("<p>")
			out.WriteString("He married " + ordinal + " " + EventDatePlace(first, marriage) + " to $partner-name{" + partner.FullName() + "}")

			birth := partner.Events["birth"]
			if birth != nil {
				out.WriteString(", born " + EventDatePlace(partner, birth))
			}
			bapt := partner.Events["baptism"]
			if bapt != nil {
				out.WriteString(", baptized " + EventDatePlace(partner, bapt))
			}
			death := partner.Events["death"]
			if death != nil {
				out.WriteString(", died " + EventDatePlace(partner, death))
			}
			burial := partner.Events["burial"]
			if burial != nil && burial.Name != "" {
				out.WriteString(", and buried at " + burial.Name)
				if burial.Ref != "" {
					out.WriteString("$ref[" + burial.Ref + "]")
				}
			}
			out.WriteString(".")
			if len(partner.Body) > 0 {
				out.WriteString(strings.Join(partner.Body, " "))
			}
		}
		out.WriteString("</p>\n")
	}

	if len(first.Todos) > 0 {
		out.WriteString("$todos{\n")
		for _, n := range first.Todos {
			out.WriteString("$todo{" + n + "}\n")
		}
		out.WriteString("}\n")
	}

	if len(first.Notes) > 0 {
		out.WriteString("$notes{\n")
		for _, n := range first.Notes {
			out.WriteString("$note{" + n + "}\n")
		}
		out.WriteString("}\n")
	}

	for _, p := range first.Body {
		out.WriteString("<p>")
		out.WriteString(p)
		out.WriteString("</p>")
	}

	out.WriteString("$children{")
	birthOrder := 1
	for _, p := range first.Partners {
		if len(p.Children) == 0 {
			continue
		}

		// if everyone was born or baptized in the same place
		places := []string{}
		for _, child := range p.Children {
			bp := child.Events["baptism"]
			if bp == nil {
				places = append(places, "")
			} else {
				places = append(places, bp.Location)
			}
		}
		bapCommon, bapType := AllMostSome(places)

		places = []string{}
		for _, child := range p.Children {
			b := child.Events["birth"]
			if b == nil {
				places = append(places, "")
			} else {
				places = append(places, b.Location)
			}
		}
		bCommon, bType := AllMostSome(places)

		common := ""
		commonEvent := ""
		commonType := ""

		switch {
		case bCommon != "" && bapCommon == "":
			commonEvent = "born"
			commonType = bType
			common = bCommon
		case bCommon == "" && bapCommon != "":
			commonEvent = "baptized"
			commonType = bapType
			common = bapCommon
		}

		commonPhrase := ""
		switch commonType {
		case "all":
			commonPhrase = "all " + commonEvent + " in " + common
		case "except":
			commonPhrase = "all " + commonEvent + " in " + common + " unless noted"
		}
		//
		// could wrap this in another container per partner
		WriteChildrenIntro(out, first, p, common, commonPhrase)
		out.WriteString("$child-list{")
		for _, child := range p.Children {
			WriteChild(out, first, child, birthOrder, common)
			birthOrder++
		}
		out.WriteString("}") /* child-list */
	}
	out.WriteString("}") /* children */
	out.WriteString("}") /* person-bio */

	out.WriteString("}") /* person-main */
	out.WriteString("}") /* person-body */
	out.WriteString(makeFootnotes(first))

	out.WriteString("}") /* person */

	cid := []string{}
	// generate next generation
	for _, p := range first.Partners {
		for _, c := range p.Children {
			if c.ID != "" {
				cid = append(cid, c.ID)
			}
		}
	}
	return out.String(), cid
}
