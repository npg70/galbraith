package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParishRef(id1, id2 string) string {
	id2 = strings.TrimLeft(id2, "0")
	if id2 == "" {
		return id1
	}
	return id1 + "/" + id2
}

func ParishName(id1, id2 string) string {

	parishmap := map[string]string{
		"074":    "Lochalsh, Ross and Cromarty",
		"108":    "Barra",
		"114":    "Raasay, Inverness",
		"152":    "Enzie,Banff",
		"228":    "Old Deer, Aberdeen",
		"280":    "Elgin",
		"301":    "Aberdeen",
		"387":    "Perth",
		"424":    "Dunfermline, Fife",
		"466":    "Clackmannan",
		"472":    "Balfron, Stirling",
		"475":    "Campsie, Stirling",
		"479":    "Falkirk",
		"487":    "Polmont",
		"488":    "St Ninians",
		"496":    "Dumbarton",
		"497":    "Kilmaronock",
		"500":    "New or East Kilpatrick, Dunbarton",
		"501":    "Old or West Kilpatrick, Dunbarton",
		"502/2":  "Kilcreggan and Cove, Dunbarton",
		"507":    "Campbeltown, Argyll",
		"507/1":  "Campbeltown (Burgh), Argyll",
		"507/2":  "Campbeltown (Landward), Argyll",
		"510":    "Dunoon, Argyll",
		"512":    "Inveraray and Glenaray",
		"516":    "Kilcalmonell and Kilberry, Argyll",
		"516/1":  "Kilcalmonell, Argyll",
		"516/2":  "Kilberry, Argyll",
		"518":    "Kilfinan, Argyll",
		"519":    "Killean and Kilchenzie, Argyll",
		"523":    "Kilmore and Kilbride, Argyll",
		"526":    "Lochgilphead, Argyll",
		"531":    "Saddell and Skipness, Argyll",
		"531/1":  "Saddell, Argyll",
		"531/2":  "Skipness, Argyll",
		"532":    "Southend, Argyll",
		"535":    "Tarbert, Argyll",
		"536":    "Campbeltown, Argyll", // not clear what this is.. on in Statutory records
		"537":    "Gigha and Cara, Argyll",
		"539/2":  "Colonsay and Oronsay",
		"554":    "Kimory, Bute",
		"555":    "Kingarth, Bute",
		"557":    "North Bute",
		"558":    "Rothesay, Bute",
		"559":    "Abbey (Paisley), Renfrew",
		"559/3":  "Johnstone and Elderslie, Renfrew",
		"560":    "Cathcart",
		"564":    "Greenock",
		"564/1":  "Greenock New or Middle",
		"564/2":  "Greenock West",
		"564/3":  "Greenock Old or West",
		"568":    "Kilbarchan, Renfrew",
		"571":    "Mearns",
		"573":    "Paisley",
		"573/1":  "Paisley High Church",
		"573/2":  "Paisley Low Church",
		"573/3":  "Paisley Burgh or Low",
		"574":    "Port Glasgow",
		"575":    "Renfrew",
		"575/1":  "Renfrew, Renfrew",
		"576":    "Ardrossan, Ayr",
		"578":    "Ayr, Ayr",
		"578/1":  "Ayr, Ayr",
		"590":    "Dundonald, Ayr",
		"594":    "Girvan, Ayr",
		"595":    "Irvine, Ayr",
		"597":    "Kilmarnock, Ayr",
		"598":    "Kilmaurs, Ayr",
		"602":    "Largs, Ayr",
		"611":    "Riccarton, Ayr",
		"611/1":  "Riccarton, Ayr",
		"620":    "West Kilbride, Ayr",
		"622":    "Barony",
		"640":    "Inverclyde",
		"642":    "Port Glasgow",
		"644":    "Glasgow",
		"644/1":  "Glasgow",
		"644/2":  "Gorbals",
		"644/3":  "Calton or Dennistoun, Glasgow",
		"644/7":  "Milton, Glasgow",
		"644/8":  "Blythswood",
		"644/9":  "Kelvin",
		"644/13": "Hillhead, Glasgow",
		"644/14": "Kinning Park, Lanark",
		"644/17": "Govan, Glasgow",
		"646":    "Govan, Lanark",
		"646/1":  "Govan, Lanark",
		"646/2":  "Govan, Lanark",
		"663":    "Bo'ness",
		"664":    "Carriden",
		"684":    "Duddingston, Midlothian",
		"685":    "Edinburgh",
		"685/1":  "Edinburgh",
		"685/2":  "St. Cuthberts",
		"685/5":  "Newington",
		"685/8":  "Leith",
		"694":    "Mid Calder, Midlothian",
		"706":    "Dunbar",
		"713":    "North Berwick",
		"743":    "Greenlaw",
		"775":    "Galashiels, Selkirk",
		"788":    "Inverness",
		"790":    "Hobkirk, Roxburgh",
	}

	ref := ParishRef(id1, id2)
	if name, ok := parishmap[ref]; ok {
		return name
	}
	log.Printf("UNKNOWN PARISH %s / %s", id1, id2)
	return ref
}

func SPLink(imageid string) string {
	base := "https://storage.googleapis.com/galbraith-research/scotlandspeople"
	return base + "/" + imageid + ".png"
}

func makeImageID(refid string, imageNum string) string {
	// CONFUSION ALERT
	// so the image id is the redif with the last field swapped for the imageNum
	// ie. refid =  d-1867-507-00-0148  imageNum = 50 -->
	// imageId== d-1867-507-00-0050

	if len(imageNum) == 0 || len(imageNum) > 4 {
		log.Fatalf("got a refid of %s, and imagenum of %q", refid, imageNum)
	}

	// pade left with zeros
	for len(imageNum) < 4 {
		imageNum = "0" + imageNum
	}

	// we assume here the refid is 4 chars so we can just cut off the old and add the new
	return refid[:len(refid)-4] + imageNum
}

func SPLinkHTML(refText string, refid string, imageNum string) string {
	target := SPLink(makeImageID(refid, imageNum))
	return fmt.Sprintf("<a href='%s'>%s</a>", target, refText)
}

// converts a full opr id into something that
// is more human readable and may be similar to what
// scotlands people uses
func oprref(parts []string) string {
	if len(parts) != 6 {
		log.Fatalf("Invalid OPR ID: %s", strings.Join(parts, "-"))
	}
	pnum := strings.TrimLeft(parts[2], "0")
	pnum2 := strings.TrimLeft(parts[3], "0")
	vol := strings.TrimLeft(parts[4], "0")
	page := strings.TrimLeft(parts[5], "0")

	return fmt.Sprintf("%s/%s %s %s", pnum, pnum2, vol, page)
}

func statref(parts []string) string {
	if len(parts) < 5 {
		log.Fatalf("bad ref: %s", strings.Join(parts, "-"))
	}
	year := parts[1]
	pnum := strings.TrimLeft(parts[2], "0")
	pnum2 := strings.TrimLeft(parts[3], "0")
	entry := strings.TrimLeft(parts[4], "0")

	return fmt.Sprintf("%s %s/%s %s", year, pnum, pnum2, entry)
}

func oprlink(parts []string) string {
	base := "https://storage.googleapis.com/galbraith-research/scotlandspeople"
	return base + "/opr-" + strings.Join(parts[2:], "-") + ".png"
}
func oprhtml(parts []string, link bool) string {
	text := oprref(parts)
	if link {
		return fmt.Sprintf("<a href=%s>%s</a>", oprlink(parts), text)
	}
	return text
}

func OPRText(refid, name, name2 string, link bool) string {
	parts := strings.Split(refid, "-")
	// args[0] --> type
	// args[1] --> year, 4 digits
	// args[2] --> parish 3 digits
	// args[3] --> subparish
	// args[4] --> volume
	// args[5] --> page, 4 digits
	if len(parts) != 6 {
		log.Fatalf("OPRText: expected 6 parts but got %s", refid)
	}
	base := fmt.Sprintf("%s. OPR of %s, %s",
		parts[1],
		ParishName(parts[2], parts[3]),
		oprhtml(parts, link))
	switch parts[0] {
	case "b":
		return fmt.Sprintf("Baptism of %s, %s", name, base)
	case "m":
		return fmt.Sprintf("Marriage of %s and %s, %s", name, name2, base)
	case "d":
		return fmt.Sprintf("Death of %s, %s", name, base)
	}
	log.Fatalf("Unknown OPR ref: %s", refid)
	return ""
}

func SPText(refid, imageNum, name, name2 string) string {
	parts := strings.Split(refid, "-")
	if len(parts) != 5 {
		log.Fatalf("SPText: expected 5 parts got %s", refid)
	}
	refText := fmt.Sprintf("%s %s %s", parts[1], ParishRef(parts[2], parts[3]),
		strings.TrimLeft(parts[4], "0"))
	if imageNum != "" {
		refText = SPLinkHTML(refText, refid, imageNum)
	}
	switch parts[0] {
	case "d":
		// d-4YEAR-3PARISH-2SUBPARISH-4RECORD
		return fmt.Sprintf("Death of %s, %s Statutory Records of %s, Reference %s",
			name, parts[1], ParishName(parts[2], parts[3]), refText)
	case "b":
		// b-4YEAR-3PARISH-2SUBPARISH-3PAGE
		return fmt.Sprintf("Birth of %s, %s Statutory Records of %s, Reference %s",
			name, parts[1], ParishName(parts[2], parts[3]), refText)
	case "m":
		// M-4YEAR-3PARISH-2SUBPARISH-3PAGE
		return fmt.Sprintf("Marriage of %s and %s, %s Statutory Records of %s, Reference %s",
			name, name2, parts[1], ParishName(parts[2], parts[3]), refText)
	default:
		log.Fatalf("Unknown SP record type: %s", refid)
	}
	return ""
}

type OPRBaptism struct {
	Surname       string
	Forename      string
	Gender        string
	Baptism       string
	Parish        string
	Ref           string
	ParishName    string
	Father        string
	Mother        string
	Transcription string
}

func (db *OPRBaptism) ReferenceImage() string {

	par := strings.Split(db.Parish, "/")
	if len(par) == 1 {
		par = append(par, "0")
	}
	ref := strings.Split(db.Ref, "/")

	par1, _ := strconv.Atoi(strings.TrimSpace(par[0]))
	par2, _ := strconv.Atoi(strings.TrimSpace(par[1]))
	ref1, _ := strconv.Atoi(strings.TrimSpace(ref[0]))
	ref2, _ := strconv.Atoi(strings.TrimSpace(ref[1]))
	if par2 == 0 {
		return fmt.Sprintf("Parish %d Volume %d Page %d", par1, ref1, ref2)
	}
	return fmt.Sprintf("Parish %d/%d Volume %d Page %d", par1, par2, ref1, ref2)
}

func (db *OPRBaptism) ReferenceLink() string {
	par := strings.Split(db.Parish, "/")
	if len(par) == 1 {
		par = append(par, "0")
	}
	ref := strings.Split(db.Ref, "/")

	par1, _ := strconv.Atoi(strings.TrimSpace(par[0]))
	par2, _ := strconv.Atoi(strings.TrimSpace(par[1]))
	ref1, _ := strconv.Atoi(strings.TrimSpace(ref[0]))
	ref2, _ := strconv.Atoi(strings.TrimSpace(ref[1]))

	name := fmt.Sprintf("opr-%03d-%02d0-%04d-%04d.png", par1, par2, ref1, ref2)
	base := "https://storage.googleapis.com/galbraith-research/scotlandspeople"
	return base + "/" + name
}

type OPRBaptismDB map[string]OPRBaptism

func LoadOPRBaptisms() (OPRBaptismDB, error) {
	fname := "./records/opr-baptisms.csv"
	fo, _ := os.Open(fname)
	r := csv.NewReader(fo)
	_, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("Unable to read header: %s", err)
	}
	db := make(OPRBaptismDB)
	for {
		args, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		uuid := args[0]
		record := OPRBaptism{
			Surname:       strings.Title(strings.ToLower(args[1])),
			Forename:      strings.Title(strings.ToLower(args[2])),
			Gender:        args[4],
			Baptism:       args[5],
			Parish:        args[6],
			Ref:           args[7],
			ParishName:    strings.Title(strings.ToLower(args[8])),
			Father:        args[9],
			Mother:        args[10],
			Transcription: args[11],
		}
		db[uuid] = record
	}
	return db, nil
}
