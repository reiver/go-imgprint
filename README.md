# go-imgprint

Package **imgprint** implements tools for printing text on images, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-imgprint

[![GoDoc](https://godoc.org/github.com/reiver/go-imgprint?status.svg)](https://godoc.org/github.com/reiver/go-imgprint)

## Example

Here a simple example of how to use **imgprint**:

```golang
// This is the image that we print the text on.
var img image.Image // = ...

// This is where on the image where we start printing the text string.
x,y = 10,50


imgprint.ImgPrint(img, x, y, "Hello world")
```

## Import

To import package **imgprint** use `import` code like the follownig:
```
import "github.com/reiver/go-imgprint"
```

## Installation

To install package **imgprint** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-imgprint
```

## Author

Package **imgprint** was written by [Charles Iliya Krempeaux](http://changelog.ca)
