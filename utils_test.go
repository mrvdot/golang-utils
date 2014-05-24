package utils

import (
	"testing"
	. "launchpad.net/gocheck"
	"appengine/aetest"
)

type MySuite struct{}

var (
	_   = Suite(&MySuite{})
	ctx aetest.Context
)

// Hook up gocheck testing library to our usual testing tool
func Test(t *testing.T) {
	TestingT(t)
}

func (s *MySuite) TestGenerateSlug(c *C) {
	testString := "My awesome string"
	want := "my-awesome-string"
	slug := GenerateSlug(testString)
	c.Assert(slug, Equals, want)
}

func (s *MySuite) TestInChain(c *C) {
	chain := []string{"one", "two", "three"}
	yep := InChain("two", chain)
	c.Assert(yep, Equals, true)
	nope := InChain("four", chain)
	c.Assert(nope, Equals, false)
}
