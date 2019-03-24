// simple forensic image file use stlib Golang
// test gif, png, jpg

// package main

// import (
// 	"fmt"
// 	"image"
// 	_ "image/gif"
// 	_ "image/jpeg"
// 	_ "image/png"
// 	"os"
// )

// func main() {
// 	if len(os.Args) == 0 {
// 		fmt.Printf("Use: %v fileUnknown fileUnknown fileUnknown ...\n", os.Args[0])
// 	} else {
// 		for _, file := range os.Args[1:] {
// 			f, err := os.Open(file)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "file: %v\n", err)
// 				continue
// 			}
// 			defer f.Close()

// 			_, kind, err := image.Decode(f)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "%v: %v\n", f.Name(), err)
// 				continue
// 			}
// 			fmt.Fprintln(os.Stderr, f.Name(), "Format =", kind)
// 		}
// 	}
// }

// Example 2
// Schedular example in golang

package main

import (
	"fmt"
	"time"
)

const INTERVAL_PERIOD time.Duration = 24 * time.Hour

const HOUR_TO_TICK int = 15
const MINUTE_TO_TICK int = 02
const SECOND_TO_TICK int = 00

type jobTicker struct {
	t *time.Timer
}

func getNextTickDuration() time.Duration {
	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), HOUR_TO_TICK, MINUTE_TO_TICK, SECOND_TO_TICK, 0, time.Local)
	if nextTick.Before(now) {
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	return nextTick.Sub(time.Now())
}

func NewJobTicker() jobTicker {
	fmt.Println("new tick here")
	return jobTicker{time.NewTimer(getNextTickDuration())}
}

func (jt jobTicker) updateJobTicker() {
	fmt.Println("next tick here")
	jt.t.Reset(getNextTickDuration())
}

func main() {
	jt := NewJobTicker()
	for {
		<-jt.t.C
		fmt.Println(time.Now(), "- just ticked")
		jt.updateJobTicker()
	}
}
