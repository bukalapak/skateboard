package element_test

import (
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/subosito/drafter-go"
	rc "github.com/subosito/drafter-go/rpc-client"
	"github.com/subosito/skateboard/element"
)

func TestFromJSON(t *testing.T) {
	c := qt.New(t)

	client := rc.New()
	defer client.Close()

	dc, err := client.Dispense()
	c.Assert(err, qt.IsNil)

	options := drafter.Options{
		Format: drafter.JSON,
	}

	b, err := dc.Parse(strings.NewReader("# API"), options)
	c.Assert(err, qt.IsNil)
	c.Assert(string(b), qt.Contains, "parseResult")

	el, err := element.FromJSON(b)

	c.Assert(err, qt.IsNil)
	c.Assert(el.Title, qt.Equals, "API")
}
