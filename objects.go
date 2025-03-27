package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/npg70/ssg"
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

type Person struct {
	parent     *Person
	generation int      // generation number 1,2,3..
	counter    int      // entry counter 1...
	refs       []string // footnote coutner per person
	lastUpdate time.Time

	root Root // not clear why I need this

	ID         string
	Name       []string
	Gender     GenderType
	Events     map[string]*Event
	Partners   []*Person
	Children   []*Person
	Footnotes  Footnotes
	External   ExternalSource
	Intro      string
	Body       []string
	Confusions []ConfusedWith
	Todos      []Todo
	Notes      []string
	Tags       []string
	Ydna       string // format TBD
}

// HasTag - simple test if a tag exists for the current person
// Case insensitive
func (p *Person) HasTag(s string) bool {
	s = strings.ToLower(s)
	for _, t := range p.Tags {
		if s == strings.ToLower(t) {
			return true
		}
	}
	return false
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

func shortEventYear(e *Event) string {
	if e == nil || e.Date.year == 0 {
		return ""
	}
	about := ""
	if e.Date.qualifier != DATE_EXACT {
		about = "~"
	}
	return fmt.Sprintf("%s%d", about, e.Date.year)
}

// BirthYearString gets the birth year from the birth or baptism event
// with qualifier (about/before/after, etc)
func (p *Person) BirthYearString() string {
	e := p.Events["birth"]
	if e == nil || e.Date.year == 0 {
		e = p.Events["baptism"]
	}
	return shortEventYear(e)
}
func (p *Person) DeathYearString() string {
	return shortEventYear(p.Events["death"])
}
func (p *Person) MarriageYearString() string {
	return shortEventYear(p.Events["marriage"])
}

func (p *Person) BirthLocation() string {
	e := p.Events["birth"]
	if e == nil || e.Location == "" {
		e = p.Events["baptism"]
	}
	if e == nil {
		return ""
	}
	return e.Location
}
func (p *Person) DeathLocation() string {
	e := p.Events["death"]
	if e == nil {
		return ""
	}
	return e.Location
}

func (p *Person) AllChildren() []*Person {
	out := []*Person{}
	for _, spouse := range p.Partners {
		for _, kid := range spouse.Children {
			out = append(out, kid)
		}
	}
	return out
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
		case "gender":
			gtype := strings.ToLower(args[1])
			switch gtype {
			case "m", "male":
				p.Gender = Male
			case "f", "female":
				p.Gender = Female
			case "u", "unknown":
				p.Gender = Unknown
			default:
				return fmt.Errorf("Got unknown gender: %s", gtype)
			}
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
		case "parent":
			// child personid reference
			if len(args) == 2 && body == "" {
				newp, err := p.root.Load(args[1])
				if err != nil {
					return err
				}
				p.parent = newp
			} else {
				newp := &Person{root: p.root}
				if err := newp.UnmarshalText([]byte(body)); err != nil {
					return err
				}
				p.parent = newp
			}
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
			if len(args) == 2 && body == "" {
				newp, err := p.root.Load(args[1])
				if err != nil {
					return err
				}
				p.Partners = append(p.Partners, newp)
			} else {
				newp := &Person{root: p.root}
				if err := newp.UnmarshalText([]byte(body)); err != nil {
					return err
				}
				p.Partners = append(p.Partners, newp)
			}
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
			txt := strings.TrimSpace(strings.Join(args[1:], " "))
			p.Body = append(p.Body, txt)
			if body != "" {
				p.Body = append(p.Body, body)
			}
		case "note":
			p.Notes = append(p.Notes, body)
		case "todo":
			switch len(args) {
			case 1:
				p.Todos = append(p.Todos, Todo{"", body})
			case 2:
				p.Todos = append(p.Todos, Todo{args[1], body})
			default:
				return fmt.Errorf("todo must be [todo tag|todo tag {body} | todo {body}")
			}
		case "confused-with":
			switch len(args) {
			case 1:
				p.Confusions = append(p.Confusions, ConfusedWith{"", body})
			case 2:
				p.Confusions = append(p.Confusions, ConfusedWith{args[1], body})
			default:
				return fmt.Errorf("confused-with takes a person id and optional body.  Got %q", args)
			}
		case "ydna":
			p.Ydna = strings.TrimSpace(strings.Join(args[1:], " "))
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

	// set map BEFORE loading in data
	// this prevents circular loops
	r[id] = newp

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

func EventDatePlaceCompact(p *Person, e *Event, commonLocation string) string {
	parts := []string{}
	if e.Date.String() != "" {
		parts = append(parts, "$date{"+e.Date.String()+"}")
	}
	normalizedLocation := TitleCompound(e.Location)
	if e.Location != "" && normalizedLocation != commonLocation {
		parts = append(parts, normalizedLocation)
	}
	// standardize common abbreviations
	note := e.Note
	switch note {
	case "dy", "dsp", "dspm", "d.y.", "d.s.p.", "d.s.p.m.", "young", "died young":
		// died in infancy or childhood implying
		// didn't marry or have children
		note = "young."

		// TODO NOTES ON dying unmarried
	case "d.um", "dum", "unm", "umn.":
		note = "unmarried."
	}
	if note != "" {
		parts = append(parts, note)
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
}
func WriteChildBioInline(w io.StringWriter, parent *Person, child *Person, ignorePlace string) {
	// no linkage, print mini-bio
	w.WriteString(WriteChildName(child.FullName()))
	birth := child.Events["birth"]
	bp := child.Events["baptism"]
	switch {
	case birth != nil && bp == nil:
		// since this child doesn't have it's own page,
		// the footnote is attached to the parent, not the child
		w.WriteString(", b.$ent[nbsp]" + EventDatePlaceCompact(parent, birth, ignorePlace))
	case birth == nil && bp != nil:
		// since this child doesn't have it's own page,
		// the footnote is attached to the parent, not the child
		w.WriteString(", bp.$ent[nbsp]" + EventDatePlaceCompact(parent, bp, ignorePlace))
	case birth != nil && bp != nil:
		// both specified.. it gets complicated
		// if both same location -- do something smart
		// if both have date -- do something smart

		w.WriteString(", b.$ent[nbsp]" + EventDatePlaceCompact(parent, birth, ignorePlace))
		w.WriteString(", bp.$ent[nbsp]" + EventDatePlaceCompact(parent, bp, ignorePlace))
	}
	death := child.Events["death"]
	if death != nil {
		w.WriteString("; ")
		// special case if child is assumed to have died infancy or
		// died young.
		txt := EventDatePlaceCompact(parent, death, "")
		w.WriteString(" d.$ent[nbsp]" + txt)
	}
	if len(child.Partners) == 1 {
		// ; m. name
		w.WriteString("; m.$ent[nbsp]" + WriteChildPartnerName(child.Partners[0]))
	} else {
		// ; m.(1) name; m.(2) name
		for i, p := range child.Partners {
			w.WriteString(fmt.Sprintf("; m.(%d)$ent[nbsp]%s", i+1, WriteChildPartnerName(p)))
		}
	}
	//w.WriteString(".")
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
		w.WriteString(", b.$ent[nbsp]" + EventDatePlaceCompact(child, birth, ignorePlace))

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
		w.WriteString(",$ent[nbsp]bp. " + EventDatePlaceCompact(child, bp, ignorePlace))

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

		w.WriteString(", b.$ent[nbsp]" + EventDatePlaceCompact(parent, birth, ignorePlace))

		// hack -- assume baptism has the reference
		oldref := bp.Ref
		// here the footnote is actually in the child page
		w.WriteString(",$ent[nbsp]bp. " + EventDatePlaceCompact(child, bp, ignorePlace))

		// if we have reference copy it into the parent
		if oldref != "" {
			if parent.Footnotes == nil {
				parent.Footnotes = make(Footnotes)
			}
			// if we have reference copy it into the parent
			parent.Footnotes[bp.Ref] = child.Footnotes[oldref]
			bp.Ref = oldref
		}
	}
	if len(child.Partners) == 1 {
		// ; m. name
		w.WriteString("; m.$ent[nbsp]" + WriteChildPartnerName(child.Partners[0]))
	} else {
		// ; m.(1) name; m.(2) name
		for i, p := range child.Partners {
			w.WriteString(fmt.Sprintf("; m.(%d)$ent[nbsp]%s", i+1, WriteChildPartnerName(p)))
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
	w.WriteString(ChildPlural(len(partner.Children)) + " of ")

	// swap order if primary is female
	if first.Gender == Female {
		first, partner = partner, first
	}
	w.WriteString(fmt.Sprintf("%s and %s (%s) %s",
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
	if p.ID == "" {
		return "$child-partner-name{" + p.FullName() + "}"
	}
	return "$child-link[" + p.ID + "]{" + p.FullName() + "}"
}

func WriteBirthDeath(p *Person) string {
	by := p.BirthYearString()
	bl := TitleCompound(p.BirthLocation())
	dy := p.DeathYearString()
	dl := TitleCompound(p.DeathLocation())

	switch {
	case by == "" && bl == "" && dy == "" && dl == "":
		{
			// birth and death locations unknown
			return ""
		}
	case by == "" && bl == "" && dy == "" && dl != "":
		{
			return "d. " + dl
		}
	case by == "" && bl == "" && dy != "" && dl == "":
		{
			return "d. " + dy
		}
	case by == "" && bl == "" && dy != "" && dl != "":
		{
			return "d. " + dy + " " + dl
		}
	case by != "" && bl != "" && dy == "" && dl == "":
		{
			return "b. " + by + " " + bl
		}
	case by != "" && bl == "" && dy == "" && dl == "":
		{
			return "b. " + by
		}
	case by != "" && bl != "" && dy == "" && dl == "":
		{
			return "b. " + by + " " + bl
		}
	case by != "" && bl != "" && dy != "" && dl == "":
		{
			return "b. " + by + " " + bl + "; d. " + dy
		}
	case by != "" && bl == "" && dy != "" && dl != "":
		{
			return "b. " + by + "; " + "d. " + dy + " " + dl
		}
	case by != "" && bl == "" && dy != "" && dl == "":
		{
			// SPECIAL CASE .. just birth and death date, no location
			return by + "-" + dy
			//return "b. " + by + "; d. " + dy
		}
	case by != "" && bl == dl && dy != "":
		// Special case  same location, have birth and death years
		{
			return by + "-" + dy + " of " + bl
		}
	case by != "" && bl != dl && dy != "":
		{
			return "b. " + by + " " + bl + "; d. " + dy + " " + dl
		}

	case by != "" && bl != dl && dy == "":
		{
			return "b. " + by + " " + bl + "; d. " + dl
		}
	}
	// if BY == DY
	log.Fatalf("Unreachable - %q %q %q %q", by, bl, dy, dl)
	return "FAIL"
}

// plain text version
func WriteTitle(p *Person) string {
	out := p.FullName()
	by := p.BirthYearString()
	if by != "" {
		out += " b. " + by
	}
	for _, spouse := range p.Partners {
		out += " m. " + spouse.FullName()
	}
	return out
}

// write a string that can be full text indexed.
func WriteFulltextDoc(p *Person) string {
	out := p.FullName()
	for _, spouse := range p.Partners {
		out += " "
		out += spouse.FullName()
	}
	return out
}

// Write MiniBio -- displayable to humans
func WriteBio(p *Person, style int) string {
	fn := p.FullName()
	switch style {
	case 0:
		// just as is
	case 1:
		fn = strings.ToUpper(fn)
	case 2:
		fn = fmt.Sprintf("<a class='fw-bold text-smallcaps' href='/galbraith/people/%s'>%s</a>", p.ID, p.FullName())
	case 3:
		fn = fmt.Sprintf("$child-link[%s]{%s}", p.ID, p.FullName())
	default:
		panic("Unknown style")
	}
	out := fn
	out += " " + WriteBirthDeath(p)
	if len(p.Partners) == 0 {
		return out
	}
	for _, spouse := range p.Partners {
		out += ";"
		out += " m. "
		if m := spouse.Events["marriage"]; m != nil {
			if ds := shortEventYear(m); ds != "" {
				out += ds
				if m.Location != "" {
					out += " " + TitleCompound(m.Location)
				}
				out += " to "
			} else {
				if m.Location != "" {
					out += " " + TitleCompound(m.Location)
					out += " to "
				}
			}
		}
		fn = spouse.FullName()
		switch style {
		case 0:
			// NOP
		case 1:
			fn = strings.ToUpper(fn)
		case 2:
			fn = "<span class='fw-bold text-smallcaps'>" + fn + "</span>"
		case 3:
			fn = "$partner-name{" + fn + "}."
		}
		out += fn
		//out += WriteBirthDeath(spouse)
	}
	return out
}

func WriteTitleLink(p *Person) string {
	out := "$child-link[" + p.ID + "]{" + p.FullName() + "}"
	by := p.BirthYearString()
	if by != "" {
		out += " b. " + by
	}
	for _, spouse := range p.Partners {
		out += " m. " + spouse.FullName()
	}
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

func (r Root) generateOne(primary string, outputFile string) (ssg.ContentSourceConfig, []string) {
	first, err := r.Load(primary)
	if err != nil {
		panic(err)
	}

	out := &bytes.Buffer{}

	out.WriteString(fmt.Sprintf("$person[id=%s generation=%d counter=%d]{",
		primary, first.generation, first.counter))

	out.WriteString("$banner{")
	out.WriteString("$headline{" + WriteTitle(first) + "}\n")
	out.WriteString("}\n")

	// all of person body text
	out.WriteString("$person-body{")

	// lineage, tags, externals
	out.WriteString("$person-secondary{")

	// get lineage from newest to oldest
	lineage := []string{}
	for p := first; p != nil; p = p.parent {
		if p.parent == nil {
			break
		}
		mother := FindMother(p)
		if mother == nil {
			log.Printf("Cant find mother for " + p.FullName())
			break
		}
		lineage = append(lineage, fmt.Sprintf("$ancestor[counter=%d generation=%d mother=%q year=%q]{%s}\n",
			p.parent.counter,
			p.generation,
			mother.FullName(),
			p.parent.BirthYearString(),
			WriteLineageNameLink(p.parent)))
	}
	out.WriteString("$lineage{")
	// and reverse
	for i := len(lineage) - 1; i >= 0; i-- {
		out.WriteString(lineage[i])
	}
	out.WriteString("}\n")

	out.WriteString("$externals{")
	for _, val := range first.External.List() {
		out.WriteString("$external[" + val.URL + "]{" + val.Name + "}\n")
	}
	out.WriteString("}\n")

	/*
		out.WriteString("$pagemeta{")
		out.WriteString(fmt.Sprintf("$elink[https://github.com/npg70/galbraith/blob/main/people/%s.sh]{Source} updated on ", first.ID))
		out.WriteString("$date{" + ParseTime(first.lastUpdate).String() + "}")
		out.WriteString("}\n")
	*/

	if len(first.Tags) > 0 {
		out.WriteString("$tags[")
		for _, tag := range first.Tags {
			out.WriteString(fmt.Sprintf("%q ", tag))
		}
		out.WriteString("]\n")
	}
	out.WriteString("}\n") // end person-secondary

	out.WriteString("$person-main{")

	// TODO: seems like this should be in person-main
	out.WriteString("$div[class='ms-5 me-5 mb-3']{")

	// this logic can be moved out
	if len(first.Intro) > 0 {
		out.WriteString("$intro{" + first.Intro + "}\n")

	}
	out.WriteString("}\n")

	out.WriteString("$person-bio{")

	out.WriteString("$p{")
	out.WriteString(fmt.Sprintf("$primary-number{%d}", first.counter))
	out.WriteString(WritePrimaryName(first))

	// Birth and Baptism
	birth := first.Events["birth"]
	bapt := first.Events["baptism"]
	if birth != nil && bapt != nil {
		out.WriteString(" was born " + EventDatePlace(first, birth))
		out.WriteString(" and baptized " + EventDatePlace(first, bapt))
	}
	if birth == nil && bapt != nil {
		out.WriteString(" was baptized " + EventDatePlace(first, bapt))
	}
	if birth != nil && bapt == nil {
		out.WriteString(" was born " + EventDatePlace(first, birth))
	}
	out.WriteString(". ")

	// Death and Burial
	death := first.Events["death"]
	if death != nil {
		out.WriteString(fmt.Sprintf("%s died %s. ",
			first.Gender.Pronoun(),
			EventDatePlace(first, death)))
	}
	out.WriteString("}") // end of paragraph

	// SPOUSES
	for i, partner := range first.Partners {
		marriage := partner.Events["marriage"]
		out.WriteString("$p{")

		ordinal := ""
		if len(first.Partners) > 1 {
			ordinal = Ordinal(i + 1)
		}
		// is marriage known or not?
		out.WriteString(first.Gender.Pronoun() + " married " + ordinal + " ")
		if marriage != nil {
			out.WriteString(EventDatePlace(first, marriage) + " to $partner-name{" + partner.FullName() + "}")
		} else {
			out.WriteString("$partner-name{" + partner.FullName() + "}")
		}
		if partner.parent != nil {
			if partner.parent.ID != "" {
				out.WriteString(", the " + partner.Gender.Child() + " of $child-link[" + partner.parent.ID + "]{" + partner.parent.FullName() + "}")
				if len(partner.parent.Partners) == 1 {
					out.WriteString(" and " + partner.parent.Partners[0].FullName())
				}
			} else {
				out.WriteString(", the " + partner.Gender.Child() + " of " + partner.parent.FullName())
				if len(partner.parent.Partners) == 1 {
					out.WriteString(" and " + partner.parent.Partners[0].FullName())
				}
			}
		}
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
		for _, pb := range partner.Body {
			out.WriteString(pb + " ")
		}
		out.WriteString("}")
	}

	if len(first.Confusions) > 0 {
		out.WriteString("$todos{\n")
		for _, n := range first.Confusions {
			out.WriteString("$todo{")
			if n.Id != "" {
				out.WriteString("Confused with " + WriteBio(r[n.Id], 3))
			}
			out.WriteString(n.Body)
			out.WriteString("}\n")
		}
		out.WriteString("}\n")
	}
	if len(first.Todos) > 0 {
		out.WriteString("$todos{\n")
		for _, n := range first.Todos {
			tag := "general"
			if n.Name != "" {
				tag = n.Name
			}
			out.WriteString("$todo{[" + tag + "] " + n.Body + "}\n")
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

	// TODO: this is prob not right
	for _, p := range first.Body {
		out.WriteString("$p{" + strings.TrimSpace(p) + "}\n")
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

	page := make(ssg.ContentSourceConfig)
	page["OutputFile"] = outputFile
	page["TemplateName"] = "baseof.html"
	page["title"] = WriteTitle(first)
	page["Content"] = out.Bytes()
	return page, cid
}
