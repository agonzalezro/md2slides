md2slides
=========

Transform markdown files to your favourite presentation tool.

At the moment the only supported format is [reveal.js](http://lab.hakim.se/reveal-js) but soon there will be more available.

Install
-------

If you are not a Go dev, the easiest way is to get the latest release from the [releases tab](https://github.com/agonzalezro/md2slides/releases).

Write your slides
-----------------

How should I write slides? It's just markdown adding a `---` to change from slide to slide. Here is a small example:

    md2slides
    =========

    ---

    This is an awesome tool!

Quick start
-----------

    md2slides -d examples/reveal/main.md

Usage
-----

    usage: md2slides [<flags>] <source>

    Flags:
      -h, --help                Show context-sensitive help (also try --help-long and --help-man).
      -o, --output=/dev/stdout  output file where to write the HTML.
      -d, --daemon              start a simple HTTP serving your slides.
          --port=8080           port where to run the server.
      -c, --theme-config=THEME-CONFIG
                                configuration for the theme (JS file)
          --theme="reveal"      Choose one: [reveal].

    Args:
      <source>  Markdown source file.

Development
-----------

### Dependencies

We use glide for this:

    glide install

### Build

There are some templates that should be added to the binary file, for doing that you will need [go-bindata](https://github.com/jteeuwen/go-bindata) and run the following:

    go generate

After that you are ready to build:

    go build
