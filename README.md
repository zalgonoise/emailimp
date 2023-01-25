## emailimp

*A package to count domain occurrences in email addresses, from CSV files*

_________

### Concept

This package will read a CSV file which contains a column for email addresses (named `email`), and extract the domain names from these addresses alongside a count for how many times the domain is present.

Here is the expected header of the CSV file:

```
first_name,last_name,email,gender,ip_address
```

________

### Approach

The task can be broken in three simple steps:
- parse the CSV data from the raw bytes in the file
- extract the domain from the email column data, and build a data structure with these occurrences (such as with a map / reduce technique)
- sort the results in alphabetical order, to organize the domain names

#### Breakdown

The steps above are first organized into three private functions `parseCSV`, `mapEmailRow` and `sortResults`.

`parseCSV` is the most straightforward to implement as it simply reads the file from the filesystem, and uses a (standard library) CSV Reader to extract all of the content as a string matrix (`[][]string`).

The second step, `mapEmailRow` will combine a map-reduce technique in one run. For this, I initialize a `map[string]int` representing a domain-to-count map or dictionary and iterate through all results only once:
- if the email field is the header (only says `email`), it's skipped
- if the extracted domain is already present in the map, increment the count value
- otherwise, add the domain to the entries map with count value of `1` (first occurrence)

The last step is to sort the results with `sortResults`. Since maps in Go do not keep order (the maps' keys are actually hash functions for comparable types), I need to extract all elements into a slice. Then, to sort the resulting slice by domain name.

Lastly, a (public) `Parse(string) ([]Entry, error)` function is written to allow this library to be imported and consumed.

#### Testing

Tests cover 95% of the statements; the remaining *uncovered* statements point to the validation done in `mapEmailRow` that verifies if the input (CSV) results are not empty and that the entries have (at least) 3 rows -- as this package was written with the CSV header on the top of this document in mind.

#### Benchmarks

The initial draft for this package benchmarked the processing actions (so, skipping the file read and CSV parsing) with the `BenchmarkMapAndSort` test. Initially this test pointed to about 5500 allocations per operation (which is surreal). 

I've profiled the package with `go test -bench . -benchmem -memprofile /tmp/mem.pprof`, which allowed me to look into the performance bottlenecks of the original implementation when running `go tool pprof -alloc_objects /tmp/mem.pprof`.

From here, it's a matter of looking into the steps that are slowing the app the most with `top10 -cum` and inspecting the lines of code slowing down the implementation with `list {path to func or method}`.

Here is a summary of some corrections that converted those 5500 allocations to 26 (in about 15 minutes of work time):

- `sortResults`: in the original implementation, the `sort.Slice()` call was iterating through all characters in one domain and comparing it to the next. Replacing with `strings.Compare()` gave a great performance boost.
- `mapEmailRow`: the domain was originally extracted with `strings.Split(domain, "@")[1]`. This was the next big flaw impacting performance. I implemented `extractDomain` to simplify this action and with it came a great performance boost.
- `mapEmailRow`: started with its output as a `map[string]*Entry`. Simplified to a map of domain-to-count instead, so replaced with `map[string]int`, where `sortResults` builds the Entry objects.
- `sortResults`: slice of Entry was originally a slice of pointers to Entry. Removing the pointer gets to the where the app is now (at 26 allocs/op for `BenchmarkMapAndSort`)

The other benchmark function (`BenchmarkParse`) shows how the overhead is present on both the I/O and the CSV parsing, especially the latter. To push further in terms of performance from here, would be to implement a lightweight CSV parser and lexer. I am saying this because I already have [the generic library to do so](https://github.com/zalgonoise/lex) :) Otherwise probing for different libraries with better performance that still ensure that the tests pass is a great option, too.

Benchmark results:

```
go test -bench . -benchmem -memprofile /tmp/mem.pprof | prettybench 

goos: linux
goarch: amd64
pkg: github.com/zalgonoise/emailimp
cpu: AMD Ryzen 3 PRO 3300U w/ Radeon Vega Mobile Gfx

PASS
benchmark               iter        time/iter   bytes alloc           allocs
---------               ----        ---------   -----------           ------
BenchmarkParse-4         409    2602.11 μs/op   697584 B/op   6063 allocs/op
BenchmarkMapAndSort-4   1618     681.74 μs/op    73721 B/op     26 allocs/op
ok      github.com/zalgonoise/emailimp  2.553s
```