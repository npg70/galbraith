package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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

	name := fmt.Sprintf("opr%03d-%02d0-%04d-%04d.jpg", par1, par2, ref1, ref2)
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
