***Warning: This library should not be used just yet***

# population

Go library for the [World Population API](http://api.population.io/)

[![Build Status](https://travis-ci.org/peterhellberg/population.svg?branch=master)](https://travis-ci.org/peterhellberg/population)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/population)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/population#license-mit)

## Installation

    go get -u github.com/peterhellberg/population

## Public API

 - **[func NewClient(httpClients ...\*http.Client) \*Client](https://godoc.org/github.com/peterhellberg/population#NewClient)**
 - **[func (c \*Client) DateByRank(sex, country, dob string, rank int) (DateByRank, error)](https://godoc.org/github.com/peterhellberg/population#Client.DateByRank)**
 - **[func (c \*Client) ListCountries() (CountryList, error)](https://godoc.org/github.com/peterhellberg/population#Client.ListCountries)**
 - **[func (c \*Client) MortalityDistributionTable(sex, country, age string) (MortalityDistributionTable, error)](https://godoc.org/github.com/peterhellberg/population#Client.MortalityDistributionTable)**
 - **[func (c \*Client) RankByAge(sex, country, dob, age string) (RankByAge, error)](https://godoc.org/github.com/peterhellberg/population#Client.RankByAge)**
 - **[func (c \*Client) RankByDate(sex, country, dob, date string) (RankByDate, error)](https://godoc.org/github.com/peterhellberg/population#Client.RankByDate)**
 - **[func (c \*Client) RankInFuture(sex, country, dob, in string) (RankWithOffset, error)](https://godoc.org/github.com/peterhellberg/population#Client.RankInFuture)**
 - **[func (c \*Client) RankInPast(sex, country, dob, ago string) (RankWithOffset, error)](https://godoc.org/github.com/peterhellberg/population#Client.RankInPast)**
 - **[func (c \*Client) RankToday(sex, country, dob string) (RankToday, error)](https://godoc.org/github.com/peterhellberg/population#Client.RankToday)**
 - **[func (c \*Client) RemainingLifeExpectancy(sex, country, date, age string) (RemainingLifeExpectancy, error)](https://godoc.org/github.com/peterhellberg/population#Client.RemainingLifeExpectancy)**
 - **[func (c \*Client) Table(country string, year, age int) (Table, error)](https://godoc.org/github.com/peterhellberg/population#Client.Table)**
 - **[func (c \*Client) TableAllAges(country string, year int) (Table, error)](https://godoc.org/github.com/peterhellberg/population#Client.TableAllAges)**
 - **[func (c \*Client) TableAllYears(country string, age int) (Table, error)](https://godoc.org/github.com/peterhellberg/population#Client.TableAllYears)**
 - **[func (c \*Client) TotalLifeExpectancy(sex, country, dob string) (TotalLifeExpectancy, error)](https://godoc.org/github.com/peterhellberg/population#Client.TotalLifeExpectancy)**
 - **[func (c \*Client) TotalPopulationByDate(country, date string) (TotalPopulation, error)](https://godoc.org/github.com/peterhellberg/population#Client.TotalPopulationByDate)**
 - **[func (c \*Client) TotalPopulationTodayAndTomorrow(country string) (TotalPopulation, TotalPopulation, error)](https://godoc.org/github.com/peterhellberg/population#Client.TotalPopulationTodayAndTomorrow)**

## Usage

**TotalLifeExpectancy**

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

	t, err := p.TotalLifeExpectancy("male", "Sweden", "1983-04-28")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	json.NewEncoder(os.Stdout).Encode(t)
}
```

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

	table, err := p.MortalityDistributionTable("male", "Sweden", "32y")
	if err == nil {
		json.NewEncoder(os.Stdout).Encode(table)
	}
}
```

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
