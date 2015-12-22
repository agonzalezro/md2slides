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

	sourceFile   *os.File
	sourceConfig *os.File
}

func NewFromFileWithConfig(sourceFile *os.File, sourceConfig *os.File) (*Presentation, error) {
	p := Presentation{sourceFile: sourceFile, sourceConfig: sourceConfig}
	if err := p.Load(); err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Presentation) Load() error {
	b, err := ioutil.ReadAll(p.sourceFile)
	if err != nil {
		return err
	}
	slides := regexp.MustCompile("\n---\n").Split(string(b), -1)

	p.Slides = []string{}
	for _, slide := range slides {
		p.Slides = append(p.Slides, markdown(slide))
	}

	// Configuration is optional
	if p.sourceConfig != nil {
		b, err := ioutil.ReadAll(p.sourceConfig)
		if err != nil {
			return err
		}
		p.Config = string(b)
	}

	return nil
}

func (p *Presentation) Reload() error {
	if _, err := p.sourceFile.Seek(0, 0); err != nil {
		return err
	}
	if p.sourceConfig != nil {
		if _, err := p.sourceConfig.Seek(0, 0); err != nil {
			return err
		}
	}
	return p.Load()
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
