package presentation

import (
	"bytes"
	"fmt"

	"github.com/russross/blackfriday"
)

type Renderer struct {
	blackfriday.Renderer
}

// Image will use the alt of the image to create the class attribute
func (r *Renderer) Image(out *bytes.Buffer, link []byte, _ []byte, alt []byte) {
	out.WriteString(fmt.Sprintf(`<img src="%s" class="%s">`, link, alt))
}

func markdown(content string) string {
	// set up the HTML renderer
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	renderer := &Renderer{blackfriday.HtmlRenderer(htmlFlags, "", "")}

	// set up the parser
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	html := blackfriday.Markdown([]byte(content), renderer, extensions)
	return string(html)
}
