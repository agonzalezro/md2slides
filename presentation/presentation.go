package presentation

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"text/template"
)

type Presentation struct {
	Slides []string
	Theme  string
}

func NewFromFile(sourcePath string) (*Presentation, error) {
	p := Presentation{}

	b, err := ioutil.ReadFile(sourcePath)
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

func (p Presentation) Write(file *os.File) error {
	t, err := p.template()
	if err != nil {
		return err
	}

	if err = t.Execute(file, p); err != nil {
		return err
	}

	return nil
}
