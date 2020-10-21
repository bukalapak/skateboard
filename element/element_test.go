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
}
