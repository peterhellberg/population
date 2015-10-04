***Warning: This library is brand new, and should not be considered stable just yet***

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

```js
{"sex":"male","country":"Sweden","dob":"1983-04-28","total_life_expectancy":85.3205032926867}
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

```js
{"countries":["Afghanistan","Albania","Algeria","Angola","Antigua and Barbuda","Azerbaijan","Argentina","Australia","Austria","The Bahamas","Bahrain","Bangladesh","Armenia","Barbados","Belgium","Bhutan","Bolivia","Bosnia and Herzegovina","Botswana","Brazil","Belize","Solomon Islands","Brunei Darussalam","Bulgaria","Myanmar","Burundi","Belarus","Cambodia","Cameroon","Canada","Cabo Verde","Central African Republic","Sri Lanka","Chad","Chile","China","Colombia","Comoros","Mayotte","Congo","Dem Rep of Congo","Costa Rica","Croatia","Cuba","Cyprus","Czech Republic","Benin","Denmark","Dominican Republic","Ecuador","El Salvador","Equatorial Guinea","Ethiopia","Eritrea","Estonia","Fiji","Finland","France","French Guiana","French Polynesia","Djibouti","Gabon","Georgia","The Gambia","West Bank and Gaza","Germany","Ghana","Kiribati","Greece","Grenada","Guadeloupe","Guam","Guatemala","Guinea","Guyana","Haiti","Honduras","Hong Kong SAR-China","Hungary","Iceland","India","Indonesia","Islamic Republic of Iran","Iraq","Ireland","Israel","Italy","Cote-d-Ivoire","Jamaica","Japan","Kazakhstan","Jordan","Kenya","Dem Peoples Rep of Korea","Rep of Korea","Kuwait","Kyrgyz Republic","Lao PDR","Lebanon","Lesotho","Latvia","Liberia","Libya","Lithuania","Luxembourg","Macao SAR China","Madagascar","Malawi","Malaysia","Maldives","Mali","Malta","Martinique","Mauritania","Mauritius","Mexico","Mongolia","Moldova","Montenegro","Morocco","Mozambique","Oman","Namibia","Nepal","The Netherlands","Curacao","Aruba","New Caledonia","Vanuatu","New Zealand","Nicaragua","Niger","Nigeria","Norway","Federated States of Micronesia","Pakistan","Panama","Papua New Guinea","Paraguay","Peru","Philippines","Poland","Portugal","Guinea-Bissau","Timor-Leste","Puerto Rico","Qatar","Reunion","Romania","Russian Federation","Rwanda","St-Lucia","St-Vincent and the Grenadines","Sao Tome and Principe","Saudi Arabia","Senegal","Serbia","Seychelles","Sierra Leone","Singapore","Slovak Republic","Vietnam","Slovenia","Somalia","South Africa","Zimbabwe","Spain","South Sudan","Sudan","Western Sahara","Suriname","Swaziland","Sweden","Switzerland","Syrian Arab Rep","Tajikistan","Thailand","Togo","Tonga","Trinidad and Tobago","United Arab Emirates","Tunisia","Turkey","Turkmenistan","Uganda","Ukraine","FYR Macedonia","Arab Rep of Egypt","United Kingdom","Channel Islands","Tanzania","United States","US Virgin Islands","Burkina Faso","Uruguay","Uzbekistan","RB-de-Venezuela","Samoa","Rep of Yemen","Zambia","World"]}
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

```js
{"mortality_distribution":[{"age":30,"mortality_percent":0},{"age":35,"mortality_percent":0.18301053314006066},{"age":40,"mortality_percent":0.3940349931720418},{"age":45,"mortality_percent":0.5754961805614758},{"age":50,"mortality_percent":0.8954022579575682},{"age":55,"mortality_percent":1.356341082175081},{"age":60,"mortality_percent":1.973425541247046},{"age":65,"mortality_percent":2.9697416640766847},{"age":70,"mortality_percent":4.56219813009874},{"age":75,"mortality_percent":7.130513701033149},{"age":80,"mortality_percent":11.41127901607065},{"age":85,"mortality_percent":17.125367814281475},{"age":90,"mortality_percent":21.00234914822359},{"age":95,"mortality_percent":17.8959645040995},{"age":100,"mortality_percent":9.52990677877034},{"age":105,"mortality_percent":2.665309500672174},{"age":110,"mortality_percent":0.3098041369864069},{"age":115,"mortality_percent":0.019271048826542916},{"age":120,"mortality_percent":0.0005743130947408195},{"age":125,"mortality_percent":9.547634164945612e-06},{"age":130,"mortality_percent":1.0787857398267798e-07}]}
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
