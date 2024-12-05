package main

type PrinterInfo struct {
	pageOrderingRules []PageOrderingRule
	updates           []Update
}

type PageOrderingRule struct {
	from int
	to   int
}

type Update struct {
	pages []int
}
