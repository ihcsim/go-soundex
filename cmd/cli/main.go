package main

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/ihcsim/go-soundex"
	"github.com/jessevdk/go-flags"
)

type opts struct {
	Info   bool `short:"i" long:"info" description:"Show Soundex algorithm information"`
	Sample bool `short:"s" long:"sample" description:"Show the sample Soundex chart"`
}

func main() {
	o := opts{}
	p := flags.NewParser(&o, flags.Default)
	args, err := p.Parse()
	if err != nil {
		if flagErr, ok := err.(*flags.Error); ok && flagErr.Type == flags.ErrHelp {
			showHelp(p)
		}
		log.Fatal(err)
	}
	p.Usage = "[OPTIONS] <name>"

	if o.Info {
		showInfo()
	}

	if o.Sample {
		showSample()
	}

	if len(args) == 0 {
		showHelp(p)
	}

	for _, arg := range args {
		if _, err := strconv.Atoi(arg); err == nil {
			fmt.Println("Unable to encode a number")
			continue
		}

		if arg[0] == '-' {
			continue
		}

		fmt.Printf("Soundex code of %s:%s\n", arg, soundex.Encode(arg))
	}
}

func showHelp(p *flags.Parser) {
	p.WriteHelp(os.Stdout)
	os.Exit(0)
}

func showInfo() {
	fmt.Println(`
American Soundex Algorithm
==========================

The Soundex value of a name is determined by:

1. Retain the first letter of the name and drop all other occurrences of a, e, i, o, u, y, h, w.
2. Replace consonants with digits as follows (after the first letter):
  * b, f, p, v → 1
  * c, g, j, k, q, s, x, z → 2
  * d, t → 3
  * l → 4
  * m, n → 5
  * r → 6
3. If two or more letters with the same number are adjacent in the original name (before step 1), only retain the first letter; also two letters with the same number separated by 'h' or 'w' are coded as a single number, whereas such letters separated by a vowel are coded twice. This rule also applies to the first letter.
4. If you have too few letters in your word that you can't assign three numbers, append with zeros until there are three numbers. If you have more than 3 letters, just retain the first 3 numbers.`)

	os.Exit(0)
}

func showSample() {
	fmt.Println(`
Soundex Code Sample
===================

The following is a list of sample names with their respective Soundex code.

Soundex Code | Name
------------ | ----
A261         | Ashcraft
A261         | Ashcroft
B620         | Burroughs
B620         | Burrows
C532         | Ciondecks
E460         | Ellery
E460         | Euler
E251         | Example
G200         | Gauss
G200         | Ghosh
H416         | Heilbronn
H416         | Hilbert
K530         | Kant
K530         | Knuth
L300         | Ladd
L222         | Lissajous
L300         | Lloyd
L222         | Lukasiewicz
O600         | O'Hara
R163         | Robert
R150         | Rubin
R163         | Rupert
S532         | Soundex
T522         | Tymczak
W350         | Wheaton
`)

	os.Exit(0)
}
