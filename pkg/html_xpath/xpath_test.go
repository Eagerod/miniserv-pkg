package html_xpath

import (
	"fmt"
	"os"
	"path"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestTextAttributeFromQueryUrl(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	urlString := fmt.Sprintf("file://%s", path.Join(cwd, "test", "test.html"))

	items, err := TextAttributeFromQueryUrl(urlString, "/html/body/a", "href")
	assert.NoError(t, err)

	assert.Equal(t, items, []string{"some legit link"})
}

func TestTextAttributeFromQueryUrlMultipleResults(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	urlString := fmt.Sprintf("file://%s", path.Join(cwd, "test", "test.html"))

	items, err := TextAttributeFromQueryUrl(urlString, "//a", "href")
	assert.NoError(t, err)

	assert.Equal(t, items, []string{"some other link", "some legit link"})
}

func TestTextAttributeFromQueryUrlNoElements(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	urlString := fmt.Sprintf("file://%s", path.Join(cwd, "test", "test.html"))

	items, err := TextAttributeFromQueryUrl(urlString, "//div", "id")
	assert.NoError(t, err)

	assert.Equal(t, items, []string{})
}

func TestTextAttributeFromQueryUrlNoAttributes(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	urlString := fmt.Sprintf("file://%s", path.Join(cwd, "test", "test.html"))

	items, err := TextAttributeFromQueryUrl(urlString, "//a", "id")
	assert.NoError(t, err)

	assert.Equal(t, items, []string{})
}

func TestTextAttributeFromQueryUrlEmptyAttributes(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	urlString := fmt.Sprintf("file://%s", path.Join(cwd, "test", "test.html"))

	items, err := TextAttributeFromQueryUrl(urlString, "//h1", "hidden")
	assert.NoError(t, err)

	assert.Equal(t, items, []string{""})
}
