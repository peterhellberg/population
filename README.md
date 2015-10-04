***Warning: This library should not be used just yet***

# population

Go library for the [World Population API](http://api.population.io/)

[![Build Status](https://travis-ci.org/peterhellberg/population.svg?branch=master)](https://travis-ci.org/peterhellberg/population)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/population)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/population#license-mit)

## Installation

    go get -u github.com/peterhellberg/population

## Usage

**ListCountries**

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/peterhellberg/population"
)

func main() {
	p := population.NewClient()

	list, err := p.ListCountries()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	json.NewEncoder(os.Stdout).Encode(list)
}
```

**MortalityDistributionTable**

```go
package main

import (
	"encoding/json"
	"os"

	"github.com/peterhellberg/population"
)

func main() {
	p := population.NewClient()

	table, err := p.MortalityDistributionTable("Sweden", "male", "32y")
	if err == nil {
		json.NewEncoder(os.Stdout).Encode(table)
	}
}
```

## Status

 - [x] [countries](http://api.population.io/#!/countries)
 - [x] [wp-rank](http://api.population.io/#!/wp-rank)
 - [ ] [life-expectancy](http://api.population.io/#!/life-expectancy)
 - [ ] [population](http://api.population.io/#!/population)
 - [x] [mortality-distribution](http://api.population.io/#!/mortality-distribution)

## License (MIT)

Copyright (c) 2015 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
