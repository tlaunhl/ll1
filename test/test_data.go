package test

const (
	EpsilonProductionConfg = `
	{
		"Nonterminals": ["S"],
		"Epsilon":      "ε",
		"Productions":	{
			"S":	[["S", "ε"]]
		}
	}`

	ProductionNotFoundConfig = `
	{
		"Nonterminals":	["S", "T"],
		"Productions":	{
			"S":	[["T"]]
		}
	}`

	LeftRecursiveGrammarConfig = `
	{
		"Nonterminals": ["S", "T", "R"],
		"Productions":	{
			"S":	[["T"]],
			"T":	[["R"]],
			"R":	[["S"]]
		}
	}`

	AmbiguousGrammarConfig = `
	{
		"Terminals":	["α"],
		"Nonterminals": ["S"],
		"Start":		"S",
		"Productions":	{
			"S":	[["α"], ["α"]]
		}
	}`

	UnmatchedSymbolsConfig = `
	{
		"Terminals":	["(", ")"],
		"Nonterminals": ["S"],
		"Start":		"S",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"S":	[["(", "S", ")", "S"], ["ε"]]
		}
	}`

	ParsingTableEntryConfig = `
	{
		"Terminals":	["α"],
		"Nonterminals": ["S"],
		"Start":		"S",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"S":	[["α"]]
		}
	}`

	BracketsSequenceConfig = `
	{
		"Terminals":	["(", ")"],
		"Nonterminals": ["S"],
		"Start":		"S",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"S":	[["(", "S", ")", "S"], ["ε"]]
		}
	}`

	ArithmeticExpressionsConfig = `
	{
		"Terminals":	["+", "*", "(", ")", "α", "β"],
		"Nonterminals":	["E", "E'", "T", "T'", "F"],
		"Start":		"E",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"E":	[["T", "E'"]],
			"E'":	[["+", "T", "E'"], ["ε"]],
			"T":	[["F", "T'"]],
			"T'":	[["*", "F", "T'"], ["ε"]],
			"F":	[["(", "E", ")"], ["α"], ["β"]]
		}
	}`

	ExtendedArithmeticExpressionsConfig = `
	{
		"Terminals":	["+", "-", "*", "/", "^", "(", ")", "α", "β", "γ"],
		"Nonterminals":	["E", "E'", "T", "T'", "F", "F'", "R"],
		"Start":		"E",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"E":	[["T", "E'"]],
			"E'":	[["+", "T", "E'"], ["-", "T", "E'"], ["ε"]],
			"T":	[["F", "T'"]],
			"T'":	[["*", "F", "T'"], ["/", "F", "T'"], ["ε"]],
			"F":	[["R", "F'"]],
			"F'":	[["^", "R", "F'"], ["ε"]],
			"R":	[["(", "E", ")"], ["α"], ["β"], ["γ"]]
		}
	}`

	NumericArithmeticExpressionsConfig = `
	{
		"Terminals":	["+", "-", "*", "/", "^", "(", ")", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."],
		"Nonterminals":	["E", "E'", "T", "T'", "F", "F'", "R", "N", "D"],
		"Start":		"E",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"E":	[["T", "E'"]],
			"E'":	[["+", "T", "E'"], ["-", "T", "E'"], ["ε"]],
			"T":	[["F", "T'"]],
			"T'":	[["*", "F", "T'"], ["/", "F", "T'"], ["ε"]],
			"F":	[["R", "F'"]],
			"F'":	[["^", "R", "F'"], ["ε"]],
			"R":	[["(", "E", ")"], ["N"]],
			"N":	[["D", "N"], ["ε"]],
			"D":	[["0"], ["1"], ["2"], ["3"], ["4"], ["5"], ["6"], ["7"], ["8"], ["9"], ["."]]
		}
	}`

	ABCDConfig = `
	{
		"Terminals":	["a", "b", "c", "d"],
		"Nonterminals":	["S", "A", "B", "C", "D"],
		"Start":		"S",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"S":	[["A", "a", "A", "b"], ["B", "b", "B", "a"]],
			"A":	[["C", "d"], ["ε"]],
			"B":	[["D", "c"], ["ε"]],
			"C":	[["ε"]],
			"D":	[["ε"]]
		}
	}`
)
