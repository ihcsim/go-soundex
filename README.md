# go-soundex

A Soundex algorithm written in Go.

## Introduction

[Soundex](https://en.wikipedia.org/wiki/Soundex) algorithms are used for indexing human names by sound, as pronounced in English. By encoding homophones to the same representation, a Soundex algorithm is able to identify matching names despite minor differences in spelling.

This project provides an implementation of the [American Soundex](https://en.wikipedia.org/wiki/Soundex#American_Soundex). As described in Wikipedia, the Soundex value of a name is determined by:

1. Retain the first letter of the name and drop all other occurrences of a, e, i, o, u, y, h, w.
1. Replace consonants with digits as follows (after the first letter):
  * b, f, p, v → 1
  * c, g, j, k, q, s, x, z → 2
  * d, t → 3
  * l → 4
  * m, n → 5
  * r → 6
1. If two or more letters with the same number are adjacent in the original name (before step 1), only retain the first letter; also two letters with the same number separated by 'h' or 'w' are coded as a single number, whereas such letters separated by a vowel are coded twice. This rule also applies to the first letter.
1. If you have too few letters in your word that you can't assign three numbers, append with zeros until there are three numbers. If you have more than 3 letters, just retain the first 3 numbers.

Using this version of the Soundex algorithm, both "Robert" and "Rupert" return the same string "R163" while "Rubin" yields "R150". "Ashcraft" and "Ashcroft" both yield "A261" and not "A226" (the chars 's' and 'c' in the name would receive a single number of 2 and not 22 since an 'h' lies in between them). "Tymczak" yields "T522" not "T520" (the chars 'z' and 'k' in the name are coded as 2 twice since a vowel lies in between them). "Pfister" yields "P236" not "P123" (the first two letters have the same number and are coded once as 'P').

Soundex is the most widely known of all phonetic algorithms (in part because it is a standard feature of popular database software such as DB2, PostgreSQL, MySQL, Ingres, MS SQL Server and Oracle) and is often used as a synonym for "phonetic algorithm".

## Getting Started

```sh
$ go get github.com/ihcsim/go-soundex             # get the code
$ go test -v github.com/ihcsim/go-soundex         # test the code
$ go install $GOPATH/github.com/ihcsim/go-soundex # build the code
$ $GOPATH/bin/go-soundex                          # start the web server
```

1. Navigate to `http://localhost:7000/testsuite` to view a list of names and their respective soundex value.
1. Soundex value of specific name can be determined by specifying the name as a query parameter like this `http://localhost:7000?name=john`.

## LICENSE

This project is under Apache v2 License. See the [LICENSE](LICENSE) file for the full license text.
