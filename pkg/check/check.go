package check

import (
	"fmt"
	"tender/storage"
	"time"
)

func StartTenderStatusUpdater(svc storage.Storage, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		err := svc.Client().CloseExpiredTenders()
		if err != nil {
			fmt.Printf("Error while updating expired tenders: %v\n", err)
		} else {
			fmt.Println("working")
		}
	}
}
