package fontawesome

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
)

var styleMap = map[string]string{
	"fab": "brands",
	"fal": "light",
	"far": "regular",
	"fas": "solid",
	"fad": "duotone",
	"fat": "thin",
}

// Library is a container for Font Awesome icons.
type Library struct {
	Path  string
	icons map[string]Icon
}

// New returns a new Font Awesome Library loaded with data from the given path.
func New(path string) (fa *Library, err error) {
	var library Library

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return &library, err
	}

	if err := json.Unmarshal(dat, &library.icons); err != nil {
		return &library, err
	}

	library.Path = path
	return &library, nil
}

// Icon returns the Font Awesome icon with the given name.
func (fa *Library) Icon(name string) (Icon, error) {
	icon, ok := fa.icons[name]
	if !ok {
		return Icon{}, fmt.Errorf("Font Awesome icon doesn't exist with name: '%v'", name)
	}

	return icon, nil
}

// SVG returns the given icon as a raw SVG element.
// Edited from original to only return template.HTML and Println any errors
func (fa *Library) SVG(prefix, name string) (template.HTML) {
	icon, err := fa.Icon(name)
	if err != nil {
		fmt.Println("FA Error:", err)
		return template.HTML("")
	}

	style, ok := styleMap[prefix]
	if !ok {
		fmt.Println("FA Error: No such icon style:", prefix)
		return template.HTML("")
	}

	svg, ok := icon.SVG[style]
	if !ok {
		fmt.Println("FA Error:", name, "icon is missing style", prefix)
		return template.HTML("")
	}

	return template.HTML(svg.Raw)
}
