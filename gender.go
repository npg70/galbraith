package main

type GenderType int

const (
	Unset GenderType = iota
	Male
	Female
	Unknown
)

func (g GenderType) Child() string {
	switch g {
	case Unset, Male:
		return "son"
	case Female:
		return "daughter"
	case Unknown:
		return "child"
	default:
		panic("should never happen")
	}
}

func (g GenderType) Pronoun() string {
	switch g {
	case Unset, Male:
		return "He"
	case Female:
		return "She"
	case Unknown:
		return "He"
	default:
		panic("should never happen")
	}
}
