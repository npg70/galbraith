package main

import (
	"fmt"
	"io"
)

// really anything thats not a whitespace, quote or brace
func isBareword(b byte) bool {
	if b < 32 {
		return false
	}
	if isSpace(b) {
		return false
	}
	if isQuote1(b) {
		return false
	}
	if isQuote2(b) {
		return false
	}
	if isLeftBrace(b) {
		return false
	}
	if isBackQuote(b) {
		return false
	}
	return true
}
func isQuote1(b byte) bool    { return b == '\'' }
func isQuote2(b byte) bool    { return b == '"' }
func isSpace(b byte) bool     { return b == ' ' || b == '\t' }
func isLeftBrace(b byte) bool { return b == '{' }
func isNewLine(b byte) bool   { return b == '\n' }
func isBackQuote(b byte) bool { return b == '`' }

type Scanner struct {
	s   []byte
	pos int
}

func NewScanner(in []byte) *Scanner {
	return &Scanner{
		s:   in,
		pos: 0,
	}
}
func (s *Scanner) parseBareword() (string, error) {
	out := ""
	i := s.pos
	for s.pos < len(s.s) {
		b := s.s[s.pos]
		switch {
		case isSpace(b):
			out += string(s.s[i:s.pos])
			return out, nil
		case isQuote1(b):
			out += string(s.s[i:s.pos])
			inner, err := s.parseQuote1()
			if err != nil {
				return out, err
			}
			i = s.pos
			out += inner
		case isQuote2(b):
			out += string(s.s[i:s.pos])
			inner, err := s.parseQuote2()
			if err != nil {
				return out, err
			}
			i = s.pos
			out += inner
		case b == '\n':
			return out + string(s.s[i:s.pos]), nil
		// TODO: backslash
		default:
			s.pos += 1
		}
	}
	if out == "" {
		// bareword till EOF
		return string(s.s[i:]), nil
	}
	return out, nil
}
func (s *Scanner) parseBackQuote() (string, error) {
	s.pos += 1
	// first char after initial quote1
	i := s.pos
	for s.pos < len(s.s) {
		b := s.s[s.pos]
		if b == '`' {
			out := string(s.s[i:s.pos])
			s.pos += 1
			return out, nil
		}
		s.pos += 1
	}
	return "", fmt.Errorf("got EOF in back quote")
}
func (s *Scanner) parseQuote1() (string, error) {
	s.pos += 1
	// first char after initial quote1
	i := s.pos
	for s.pos < len(s.s) {
		b := s.s[s.pos]
		switch b {
		case '\'':
			out := string(s.s[i:s.pos])
			s.pos += 1
			return out, nil
		case '\n':
			return "", fmt.Errorf("got newline in single quote")
		// TODO: backslash
		default:
			s.pos += 1
		}
	}
	return "", fmt.Errorf("got EOF in single quote")
}
func (s *Scanner) parseQuote2() (string, error) {
	s.pos += 1
	// first char after initial quote1
	i := s.pos
	for s.pos < len(s.s) {
		b := s.s[s.pos]
		switch b {
		case '"':
			out := string(s.s[i:s.pos])
			s.pos += 1
			return out, nil
		case '\n':
			return "", fmt.Errorf("got newline in single quote")
		// TODO: backslash,
		default:
			s.pos += 1
		}
	}
	return "", fmt.Errorf("got EOF in single quote")
}

// TODO - handle escaped characters in particular braces
func (s *Scanner) parseBrace() (string, error) {
	// skip brace
	s.pos += 1
	// first char after opening '{''
	i := s.pos
	stack := 1
	for s.pos < len(s.s) {
		b := s.s[s.pos]
		switch b {
		case '{':
			stack += 1
		case '}':
			stack -= 1
			if stack == 0 {
				out := string(s.s[i:s.pos])
				s.pos += 1
				return out, nil
			}
		}
		s.pos += 1
	}
	return "", fmt.Errorf("got EOF in opening brace")
}

// Next returns the arguments and the body, along with an error if any.
// foo bar { the body }
// --> []string{"foo", "bar"}, "the body"
func (s *Scanner) Next() ([]string, string, error) {
	args := []string{}
	body := ""

	arg := ""
	var err error

	for s.pos < len(s.s) {
		b := s.s[s.pos]
		switch {
		case isSpace(b):
			s.pos += 1
		case isBareword(b) || isQuote1(b) || isQuote2(b):
			// no error here possible
			arg, _ = s.parseBareword()
			args = append(args, arg)
		case isBackQuote(b):
			arg, err = s.parseBackQuote()
			if err != nil {
				return args, body, err
			}
			args = append(args, arg)
		case isLeftBrace(b):
			body, err = s.parseBrace()
			return args, body, err
		case isNewLine(b):
			s.pos += 1
			if len(args) > 0 {
				return args, body, nil
			}
		}
		// TODO: case blackslash
		// TODO: # comments
	}

	// nothing to do.. end of file
	if len(args) == 0 {
		return nil, "", io.EOF
	}
	return args, "", nil
}
