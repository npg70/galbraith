package main

import (
	"strings"

	tf "github.com/client9/tagfunctions"
	"golang.org/x/net/html"
)

//  foo     foo           --> yes
//  foo:bar foo           --> yes
//  foo:bar foo:bing      --> yes
//  foo:bar:bing foo:bar  --> yes
//
//  case x: foo:bar foo:bar:bing  --> no
//          foo     foo:bar       --> no
//
//  foo     foo:bar       --> no
//  junk    foo:bar       --> no
//  junk:a  crap:b        --> no

func labelSubset(a, b string) bool {
	// handle case where they are same
	if a == b {
		return true
	}
	aparts := strings.Split(a, ":")
	bparts := strings.Split(b, ":")

	min := len(aparts)
	if min == 0 {
		// "", "foo" --> add it
		return true
	}
	if len(bparts) < min {
		min = len(bparts)
	}
	if min == 0 {
		// "foo", "" --> weird case, don't add it
		return false
	}

	// min is not zero, both a and b contain stuff
	for i := 0; i < min; i++ {
		if aparts[i] == bparts[i] {
			continue
		}
		if i == 0 {
			return false
		}
		return true
	}

	// if we got here, a or b is complete subset of the other

	// b is subset of a
	if min == len(bparts) {
		// len(b) < len(a)
		return true
	}

	// len(b) > len(a).. b is not a subset of a
	return false
}

func inLabelPath(path []string, label string) bool {
	for _, p := range path {
		if labelSubset(p, label) {
			return true
		}
	}
	return false
}

func makeTagButton(path []string, body string) *html.Node {

	if len(path) == 1 && path[0] == "" {
		path = nil
	}
	if len(path) == 0 && !strings.HasPrefix(body, "all") {
		path = []string{body}
	}
	tag := body // TitleTag(body)

	taglink := "/galbraith/tags?q=" + tagsQueryString(tagsNormalize(path))

	return tf.Append(
		tf.NewElement("a",
			"class", "me-2",
			"href", taglink),
		tf.NewText("#"+tag))
}
