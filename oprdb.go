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
	if id2 == "" || id2 == "00" || id2 == "000" {
		return id1
	}
	return id1 + "/" + strings.TrimLeft(id2, "0")
}

func ParishName(id1, id2 string) string {
	parishmap := map[string]string{
		"108":   "Barra",
		"280":   "Elgin",
		"301":   "Aberdeen",
		"472":   "Balfron",
		"479":   "Falkirk",
		"487":   "Polmont",
		"488":   "St Ninians",
		"497":   "KILMARONOCK",
		"500":   "NEW OR EAST KILPATRICK",
		"507":   "Campbeltown, Argyll",
		"512":   "INVERARAY AND GLENARAY",
		"516":   "KILCALMONELL AND KILBERRY",
		"516/2": "Kilberry",
		"526":   "Lochgilphead, Argyll",
		"537":   "Gigha and Cara, Argyll",
		"539/2": "COLONSAY AND ORONSAY",
		"554":   "Kimory, Bute",
		"560":   "CATHCART",
		"564/1": "Greenock New or Middle",
		"564/2": "Greenock West",
		"564/3": "Greenock Old or West",
		"573/2": "Paisley Low Church",
		"575":   "Renfrew",
		"574":   "Port Glasgow",
		"590":   "Dundonald, Ayr",
		"595":   "Irvine, Ayr",
		"597":   "Kilmarnock, Ayr",
		"602":   "Largs, Ayr",
		"611":   "Riccarton, Ayr",
		"611/1": "Riccarton, Ayr",
		"622":   "Barony",
		"640":   "Inverclyde",
		"644/1": "Glasgow",
		"644/2": "Gorbals",
		"644/3": "Calton, Glasgow",
		"646/2": "Govan, Lanark",
		"663":   "Bo'ness",
		"664":   "Carriden",
		"706":   "Dunbar",
		"743":   "Greenlaw",
		"788":   "INVERNESS",
	}

	ref := ParishRef(id1, id2)
	if name, ok := parishmap[ref]; ok {
		return name
	}
	return ref
}

func SPLink(pageid string) string {
	base := "https://storage.googleapis.com/galbraith-research/scotlandspeople"
	return base + "/" + pageid + ".png"
}

func SPLinkHTML(ref string, pageid string) string {
	return fmt.Sprintf("<a href='%s'>%s</a>", SPLink(pageid), ref)
}

func SPText(pageid, name, name2 string, link bool) string {
	parts := strings.Split(pageid, "-")
	switch parts[0] {
	case "d":
		if len(parts) != 5 {
			log.Fatalf("Expected 5 parts, got %q", pageid)
		}
		// D-4YEAR-3PARISH-2SUBPARISH-3PAGE
		ref := fmt.Sprintf("%s %s %s", parts[1], ParishRef(parts[2], parts[3]),
			strings.TrimLeft(parts[4], "0"))
		if link {
			ref = SPLinkHTML(ref, pageid)
		}
		return fmt.Sprintf("Death of %s, %s Statutory Records of %s, Scotlands People, Reference %s",
			name, parts[1], ParishName(parts[2], parts[3]), ref)
	case "b":
		// D-4YEAR-3PARISH-2SUBPARISH-3PAGE

		ref := fmt.Sprintf("%s %s %s", parts[1], ParishRef(parts[2], parts[3]),
			strings.TrimLeft(parts[4], "0"))
		if link {
			ref = SPLinkHTML(ref, pageid)
		}
		return fmt.Sprintf("Birth of %s, %s Statutory Records of %s, Scotlands People, Reference %s",
			name, parts[1], ParishName(parts[2], parts[3]), ref)
	case "m":
		// M-4YEAR-3PARISH-2SUBPARISH-3PAGE
		ref := fmt.Sprintf("%s %s %s", parts[1], ParishRef(parts[2], parts[3]),
			strings.TrimLeft(parts[4], "0"))
		if link {
			ref = SPLinkHTML(ref, pageid)
		}
		return fmt.Sprintf("Marriage of %s and %s, %s Statutory Records of %s, Scotlands People, Reference %s",
			name, name2, parts[1], ParishName(parts[2], parts[3]), ref)
	default:
		log.Fatalf("Unknown SP record type: %s", pageid)
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
