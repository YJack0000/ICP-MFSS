package middlewares

import (
    "fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/YJack0000/ICP-MFSS/utils"
)

type ipRecord struct {
	lastTime time.Time
	mutex    sync.Mutex
}

func RateLimiter(delay time.Duration) gin.HandlerFunc {
	ipRecords := make(map[string]*ipRecord)
	var mutex sync.RWMutex

	return func(c *gin.Context) {
		ip := c.ClientIP()

		mutex.RLock()
		record, exists := ipRecords[ip]
		mutex.RUnlock()

		if exists {
			record.mutex.Lock()
            fmt.Println(time.Since(record.lastTime), delay)
			if time.Since(record.lastTime) < delay {
				c.Data(http.StatusAccepted, "text/html", utils.GetUploadErrorResponse("Too many reqsts"))
                c.Abort()
				record.mutex.Unlock()
				return
			}
			record.lastTime = time.Now()
			record.mutex.Unlock()
		} else {
			mutex.Lock()
			ipRecords[ip] = &ipRecord{lastTime: time.Now()}
			mutex.Unlock()
		}

		c.Next()
	}
}
