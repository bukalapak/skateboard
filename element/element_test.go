package element_test

import (
	"os"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/subosito/drafter-go"
	rc "github.com/subosito/drafter-go/rpc-plugin/client"
	"github.com/subosito/skateboard/element"
)

func TestFromJSON(t *testing.T) {
	c := qt.New(t)

	client := rc.New(rc.Config{Name: "skateboard-rpc"})
	defer client.Close()

	dc, err := client.Dispense()
	c.Assert(err, qt.IsNil)

	options := drafter.Options{
		Format: drafter.JSON,
	}

	f, err := os.Open("./testdata/Real World API.md")
	c.Assert(err, qt.IsNil)
	defer f.Close()

	b, err := dc.Parse(f, options)
	c.Assert(err, qt.IsNil)
	c.Assert(string(b), qt.Contains, "parseResult")

	el, err := element.FromJSON(b)
	c.Assert(err, qt.IsNil)
	c.Assert(el.Title, qt.Equals, "Real World API")
	c.Assert(el.Description, qt.Contains, "This API Blueprint demonstrates a real world example")

	f2, err := os.Open("./testdata/Parameters.md")
	c.Assert(err, qt.IsNil)
	defer f.Close()

	b2, err := dc.Parse(f2, options)
	c.Assert(err, qt.IsNil)
	c.Assert(string(b2), qt.Contains, "parseResult")

	el2, err := element.FromJSON(b2)
	c.Assert(err, qt.IsNil)
	c.Assert(el2.Title, qt.Equals, "Parameters API")

	hv := el2.ResourceGroups[0].Resources[1].Transitions[0].Href.Variables[0]
	c.Assert(hv.Required, qt.IsFalse)
	c.Assert(hv.Key, qt.Equals, "limit")
	c.Assert(hv.Title, qt.Equals, "number")
	c.Assert(hv.Value, qt.Equals, "20")
	c.Assert(hv.Description, qt.Equals, "The maximum number of results to return.")
}
