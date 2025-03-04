package compression

import (
	"binvault/pkg/clients/filesystem"
	"log"
	"sync"
)

func Init(workers int) {
	log.Printf("Starting compression workers [%d]", workers)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for filePath := range filesystem.FileQueue {
				runWorker(filePath)
			}
		}()
	}
}
