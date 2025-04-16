package htmx

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElementA(t *testing.T) {
	e := Element("a")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<a></a>", e.String())
}

func TestElementDiv(t *testing.T) {
	e := Element("div")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<div></div>", e.String())
}

func TestElementSpan(t *testing.T) {
	e := Element("span")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<span></span>", e.String())
}

func TestElementP(t *testing.T) {
	e := Element("p")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<p></p>", e.String())
}

func TestElementH1(t *testing.T) {
	e := Element("h1")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<h1></h1>", e.String())
}

func TestElementH2(t *testing.T) {
	e := Element("h2")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<h2></h2>", e.String())
}

func TestElementH3(t *testing.T) {
	e := Element("h3")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<h3></h3>", e.String())
}

func TestElementH4(t *testing.T) {
	e := Element("h4")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<h4></h4>", e.String())
}

func TestElementH5(t *testing.T) {
	e := Element("h5")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<h5></h5>", e.String())
}

func TestElementH6(t *testing.T) {
	e := Element("h6")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<h6></h6>", e.String())
}

func TestElementCustom(t *testing.T) {
	e := Element("custom")
	assert.NotNil(t, e)
	assert.Implements(t, (*fmt.Stringer)(nil), e)
	assert.Equal(t, ElementType, e.Type())
	assert.Equal(t, "<custom></custom>", e.String())
}

func ExampleElement_a() {
	_ = Element("a").Render(os.Stdout)
	// Output: <a></a>
}

func ExampleElement_div() {
	_ = Element("div").Render(os.Stdout)
	// Output: <div></div>
}
