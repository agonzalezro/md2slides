md2slides
=========

Transform markdown files to your favourite presentation tool.

At the moment the only supported format is [reveal.js](http://lab.hakim.se/reveal-js) but soon there will be more available.

This is a pre-alpha, you should just use it if you have any idea, so I will explain the dev usage only.

Dependencies
------------

We use glide for this:

    glide install

Build
-----

There are some templates that should be added to the binary file, for doing that you will need [go-bindata](https://github.com/jteeuwen/go-bindata) and run the following:

    go generate

After that you are ready to build:

    go build

Usage
-----

Take a look to the output of `-h`.

Slides
------

How should I write slides? It's just markdown adding a `---` to change from slide to slide. Here is a small example:

    md2slides
    =========

    ---

    This is an awesome tool!
