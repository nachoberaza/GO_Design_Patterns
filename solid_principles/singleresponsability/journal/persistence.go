package journal

import (
	"encoding/json"
	"os"
)

type Persistence struct {
}

func (p *Persistence) SaveToFile(journal *Journal, filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer file.Close()

	serializedJournal, err := json.MarshalIndent(journal, "", "")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, serializedJournal, 0644)
	if err != nil {
		return err
	}

	return nil
}
