package regex

import "regexp"

var AC, _ = regexp.Compile(`(?P<year>\d{4})(?P<symbol1>[-/]*)(?P<month>\d{2})(?P<symbol2>[-/]*)(?P<day>\d{2})`)
// 西元年: 20210318 | 2021-03-18 | 2021/03/18
var RepublicEra, _ = regexp.Compile(`(?P<year>\d{1,3})(?P<symbol1>[-/]*)(?P<month>\d{2})(?P<symbol2>[-/]*)(?P<day>\d{2})`)
// 民國年: 1100318 | 110-03-18 | 110/03/18