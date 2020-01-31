#### LL1

Builds LL(1) parser from given grammar and checks
if a string can be generated by this grammar.

Example:
```go
package main

import (
	"fmt"
	"github.com/tlaunhl/ll1"
)

func main() {
	config := `
	{
		"Terminals":	["+", "*", "(", ")", "α"],
		"Nonterminals":	["E", "E'", "T", "T'", "F"],
		"Start":		"E",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"E":	[["T", "E'"]],
			"E'":	[["+", "T", "E'"], ["ε"]],
			"T":	[["F", "T'"]],
			"T'":	[["*", "F", "T'"], ["ε"]],
			"F":	[["(", "E", ")"], ["α"]]
		}
	}`

	if p, err := ll1.NewParser(config); err != nil {
		fmt.Println(err)
	} else {
		if err := p.Parse("(α+α)*α"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("OK")
		}
	}
}
```
