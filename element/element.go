package element

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs/v2"
)

var ErrUnsupportedElement = errors.New("unsupported element")

func FromJSON(b []byte) (*API, error) {
	c, err := gabs.ParseJSON(b)
	if err != nil {
		return nil, err
	}

	if str, ok := c.Path("element").Data().(string); !ok || str != "parseResult" {
		return nil, ErrUnsupportedElement
	}

	return fromContainer(c)
}

func fromContainer(c *gabs.Container) (*API, error) {
	a := new(API)

	for _, child := range c.Path("content").Children() {
		if isElement(child, "category") {
			a.Title = toTitle(child)
			a.Description = toDescription(child)
			a.Metadata = toMetadata(child)
			a.ResourceGroups = toResourceGroups(child)
			a.Resources = toResources(child)
		}
	}

	return a, nil
}

func toTitle(c *gabs.Container) string {
	return toString(c.Path("meta.title.content"))
}

func toDescription(c *gabs.Container) string {
	var ss []string

	for _, child := range filterElement(c, "copy") {
		ss = append(ss, toString(child.Path("content")))
	}

	return strings.Join(ss, "\n")
}

func toMetadata(c *gabs.Container) []Metadata {
	children := c.Path("attributes.metadata.content").Children()
	metadata := make([]Metadata, len(children))

	for i, child := range children {
		metadata[i] = Metadata{
			Key:   toString(child.Path("content.key.content")),
			Value: toString(child.Path("content.value.content")),
		}
	}

	return metadata
}

func toResourceGroups(c *gabs.Container) []ResourceGroup {
	var children []*gabs.Container

	for _, child := range filterElement(c, "category") {
		if hasClass(child, "resourceGroup") {
			children = append(children, child)
		}
	}

	groups := make([]ResourceGroup, len(children))

	for i, child := range children {
		groups[i] = ResourceGroup{
			Title:       toTitle(child),
			Description: toDescription(child),
			Resources:   toResources(child),
		}
	}

	return groups
}

func toResources(c *gabs.Container) []Resource {
	children := filterElement(c, "resource")
	resources := make([]Resource, len(children))

	for i, child := range children {
		resources[i] = Resource{
			Title:       toTitle(child),
			Description: toDescription(child),
			Href:        toHref(child),
			Transitions: toTransitions(child),
		}
	}

	return resources
}

func toTransitions(c *gabs.Container) []Transition {
	children := filterElement(c, "transition")
	transitions := make([]Transition, len(children))

	for i, child := range children {
		transitions[i] = Transition{
			Title:        toTitle(child),
			Description:  toDescription(child),
			Href:         toHref(child),
			Transactions: toTransactions(child),
		}
	}

	return transitions
}

func toTransactions(c *gabs.Container) []Transaction {
	children := filterElement(c, "httpTransaction")
	transactions := make([]Transaction, len(children))

	for i, child := range children {
		for _, grandChild := range child.Path("content").Children() {
			if isElement(grandChild, "httpRequest") {
				transactions[i].Request = toRequest(grandChild)
			}

			if isElement(grandChild, "httpResponse") {
				transactions[i].Response = toResponse(grandChild)
			}
		}
	}

	return transactions
}

func toRequest(c *gabs.Container) Request {
	return Request{
		Description: toDescription(c),
		Method:      toString(c.Path("attributes.method.content")),
		Headers:     toHeaders(c),
		Body:        toAsset(c, "messageBody"),
		Schema:      toAsset(c, "messageBodySchema"),
	}
}

func toResponse(c *gabs.Container) Response {
	return Response{
		Description: toDescription(c),
		StatusCode:  toStatusCode(c),
		Headers:     toHeaders(c),
		Body:        toAsset(c, "messageBody"),
		Schema:      toAsset(c, "messageBodySchema"),
	}
}

func toAsset(c *gabs.Container, s string) Asset {
	for _, child := range filterElement(c, "asset") {
		if hasClass(child, s) {
			return Asset{
				ContentType: toString(child.Path("attributes.contentType.content")),
				Body:        toString(child.Path("content")),
			}
		}
	}

	return Asset{}
}

func toStatusCode(c *gabs.Container) (n int) {
	if s := toString(c.Path("attributes.statusCode.content")); s != "" {
		if v, err := strconv.Atoi(s); err == nil {
			n = v
		}
	}

	return
}

func toHref(c *gabs.Container) Href {
	return Href{
		Path:      toString(c.Path("attributes.href.content")),
		Variables: toHrefVariables(c),
	}
}

func toHrefVariables(c *gabs.Container) []HrefVariable {
	children := c.Path("attributes.hrefVariables.content").Children()
	items := make([]HrefVariable, len(children))

	for i, child := range children {
		items[i] = HrefVariable{
			Title:       toTitle(child),
			Description: toDescription(child),
			Required:    isRequired(child),
			Key:         toString(child.Path("content.key.content")),
			Value:       toString(child.Path("content.value.content")),
		}
	}

	return items
}

func toHeaders(c *gabs.Container) []Header {
	children := c.Path("attributes.headers.content").Children()
	headers := make([]Header, len(children))

	for i, child := range children {
		headers[i] = Header{
			Key:   toString(child.Path("content.key.content")),
			Value: toString(child.Path("content.value.content")),
		}
	}

	return headers
}

func toString(c *gabs.Container) string {
	if str := c.String(); str != "null" {
		if s, err := strconv.Unquote(str); err == nil {
			return s
		}
	}

	return ""
}

func isRequired(c *gabs.Container) bool {
	return toString(c.Path("attributes.typeAttributes.content.0.content")) == "required"
}

func filterElement(c *gabs.Container, s string) []*gabs.Container {
	var children []*gabs.Container

	for _, child := range c.Path("content").Children() {
		if isElement(child, s) {
			children = append(children, child)
		}
	}

	return children
}

func hasClass(c *gabs.Container, s string) bool {
	for _, child := range c.Path("meta.classes.content").Children() {
		return toString(child.Path("content")) == s
	}

	return false
}

func isElement(c *gabs.Container, s string) bool {
	return toString(c.Path("element")) == s
}
