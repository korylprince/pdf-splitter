# About

This is a simple utility to split a PDF file into separate pages. The pages are named with a value obtained from each page with the given regular expression. It is expected that the regular expression has one capture group, and that group is used as the value.

# Usage

    Usage of pdf-splitter:
      -debug
            output extracted text for each page
      -in string
            input PDF
      -out string
            directory for outputing PDFs
      -re string
            regular expression for value in PDF page content

# Example

    pdf-splitter -in "input.pdf" -out "/tmp/output" -re "Name: ([a-zA-Z ]+)"

# License

This utility relies heavily on the [UniDoc](https://github.com/unidoc/unidoc) library. This library uses a vendored version of UniDoc that removes the licensing code. This modification is done under their provided AGPLv3 license. Therefore this code is also licensed under AGPLv3.
