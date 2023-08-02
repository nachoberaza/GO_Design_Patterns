package journal

import (
	"fmt"
)

var entryCount = 0

type Journal struct {
	Entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d %s", entryCount, text)
	j.Entries = append(j.Entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(_ int) {
	fmt.Println("Not Implemented yet!")
}
