package f8time

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	// Current date and time is:  2018-08-10 21:10:39.121597055 +0530 IST
	fmt.Println("Current date and time is: ", dt.String())

	//Format MM-DD-YYYY => 08-10-2018
	fmt.Println(dt.Format("01-02-2006"))

	//Format MM-DD-YYYY hh:mm:ss => 08-10-2018 21:11:58
	fmt.Println(dt.Format("01-02-2006 15:04:05"))

	//With short weekday (Mon) => 08-10-2018 21:11:58 Fri
	fmt.Println(dt.Format("01-02-2006 15:04:05 Mon"))

	//With weekday (Monday) => 08-10-2018 21:11:58 Friday
	fmt.Println(dt.Format("01-02-2006 15:04:05 Monday"))

	//Include micro seconds => 08-10-2018 21:11:58.880934
	fmt.Println(dt.Format("01-02-2006 15:04:05.000000"))

	//Include nano seconds => 08-10-2018 21:11:58.880934320
	fmt.Println(dt.Format("01-02-2006 15:04:05.000000000"))
}

// ============================================================================
//                        code in package:time
// ============================================================================
type Duration int64

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

var currentTime time.Time = time.Now()
var currentStamp int64 = time.Now().Unix()
var formatTime string = time.Now().Format("01-01-2006 15:03:22")

// ============================================================================

type TrafficMapCache struct {
	TrafficMap map[uint64]string
	LastUpdate time.Time
	Interval   time.Duration
}

var trafficMapCache = &TrafficMapCache{
	TrafficMap: make(map[uint64]string, 100),
	Interval:   time.Second * 10,
}

func Query() {
	if time.Since(trafficMapCache.LastUpdate) > trafficMapCache.Interval {
		trafficMapCache.LastUpdate = time.Now()
	}
}
