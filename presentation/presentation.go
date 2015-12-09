package presentation

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"text/template"
)

type Presentation struct {
	Slides []string
	Config string
	Theme  string
}

func NewFromFile(sourceFile *os.File) (*Presentation, error) {
	p := Presentation{}

	b, err := ioutil.ReadAll(sourceFile)
	if err != nil {
		return nil, err
	}
	slides := regexp.MustCompile("\n---\n").Split(string(b), -1)

	for _, slide := range slides {
		p.Slides = append(p.Slides, markdown(slide))
	}

	return &p, nil
}

func (p Presentation) template() (*template.Template, error) {
	assetFile := fmt.Sprintf("templates/%s.html", p.Theme)
	tpl, err := Asset(assetFile)
	if err != nil {
		return nil, err
	}

	t, err := template.New("").Parse(string(tpl))
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (p Presentation) WriteWithConfig(wr io.Writer, config *os.File) error {
	if config != nil {
		b, err := ioutil.ReadAll(config)
		if err != nil {
			return err
		}
		p.Config = string(b)
	}

	t, err := p.template()
	if err != nil {
		return err
	}

	if err = t.Execute(wr, p); err != nil {
		return err
	}

	return nil
}
