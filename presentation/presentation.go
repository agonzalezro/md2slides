package presentation

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"text/template"
)

type Presentation struct {
	Slides []string
	Config string
	Theme  string

	RawContent string

	sourceFile   string
	sourceConfig string
}

func NewFromFileWithConfig(sourceFile, sourceConfig string) (*Presentation, error) {
	p := Presentation{sourceFile: sourceFile, sourceConfig: sourceConfig}
	if err := p.Load(); err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Presentation) Load() error {
	b, err := ioutil.ReadFile(p.sourceFile)
	if err != nil {
		return err
	}
	p.RawContent = string(b)
	slides := regexp.MustCompile("\n---\n").Split(p.RawContent, -1)

	p.Slides = []string{}
	for _, slide := range slides {
		p.Slides = append(p.Slides, markdown(slide))
	}

	// Configuration is optional
	if p.sourceConfig != "" {
		b, err := ioutil.ReadFile(p.sourceConfig)
		if err != nil {
			return err
		}
		p.Config = string(b)
	}

	return nil
}

func (p Presentation) template() (*template.Template, error) {
	assetFile := fmt.Sprintf("templates/%s.tmpl", p.Theme)
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

func (p Presentation) Write(wr io.Writer) error {
	t, err := p.template()
	if err != nil {
		return err
	}

	if err = t.Execute(wr, p); err != nil {
		return err
	}

	return nil
}
