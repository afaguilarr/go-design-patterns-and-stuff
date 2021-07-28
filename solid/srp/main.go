package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// journal responsibility starts
var entryCount = 0

type journal struct {
	entries []string
}

func (j *journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount,
		text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// journal responsibility ends, just logic bound to the journal abstraction manipulation

// The responsibility of storing/persisting the journal data
// shouldn't be bound to the journal abstraction itself
func SaveToFile(c string, filename string) {
	ioutil.WriteFile(filename, []byte(c), 0644)
}

func main() {
	j := journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug")
	// separate function
	SaveToFile(j.String(), "journal.txt")
}
