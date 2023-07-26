package journal

import (
	"fmt"
	"github.com/edsrzf/mmap-go"
	"os"
	"sync"
)

type Persistence struct {
	FileLocker sync.Mutex
}

func (p *Persistence) SaveToFile(journal *Journal, filename string) error {
	file, openFileErr := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if openFileErr != nil {
		return openFileErr
	}

	// TODO: Serialize journal struct
	//_, writeError := file.Write(journal)

	defer file.Close()

	mMap, mMapErr := mmap.Map(file, mmap.RDWR, 0)
	if mMapErr != nil {
		return mMapErr
	}

	fmt.Printf("Showing map:\n")
	fmt.Println(string(mMap))
	fmt.Printf("Map showed:\n\n")

	if err := mMap.Unmap(); err != nil {
		return err
	}

	err := mMap.Flush()
	if err != nil {
		return err
	}
	return nil
}
