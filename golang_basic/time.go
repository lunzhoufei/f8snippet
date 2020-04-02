package f8time

import (
	"fmt"
	"time"
)


// 其中layout的时间必须是"2006-01-02 15:04:05"这个时间，
// 不管格式如何，时间点一定得是这个，
// 如："Jan 2, 2006 at 3:04pm (MST)"，"2006-Jan-02"等
// 如换一个时间解析出来的时间就不对了，要特别注意这一点。

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

// ============================================================================

package main

import (
	"fmt"
	"time"
)

func main() {
	// See the example for Time.Format for a thorough description of how
	// to define the layout string to parse a time.Time value; Parse and
	// Format use the same model to describe their input and output.

	// longForm shows by example how the reference time would be represented in
	// the desired layout.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	// shortForm is another way the reference time would be represented
	// in the desired layout; it has no time zone present.
	// Note: without explicit zone, returns time in UTC.
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	// Some valid layouts are invalid time values, due to format specifiers
	// such as _ for space padding and Z for zone information.
	// For example the RFC3339 layout 2006-01-02T15:04:05Z07:00
	// contains both Z and a time zone offset in order to handle both valid options:
	// 2006-01-02T15:04:05Z
	// 2006-01-02T15:04:05+07:00
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	fmt.Println(t)
	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error", err) // Returns an error as the layout is not a valid time value

}
