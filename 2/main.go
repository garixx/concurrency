// Hint 1: time.Ticker can be used to cancel function
// Hint 2: to calculate time-diff for Advanced lvl use:
//  start := time.Now()
//	// your work
//	t := time.Now()
//	elapsed := t.Sub(start) // 1s or whatever time has passed

package main

import (
	//"sync"
	"time"
)

const (
	almostTenSeconds = 9_900
	tenSeconds       = 10_900
	tenMillis        = 10
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	//mu        sync.Mutex
	ID        int
	IsPremium bool  // can be used for 2nd level task. Premium users won't have 10 seconds limit.
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if !u.IsPremium {
		select {
		case <-time.After(10 * time.Second):
			return false
		}
	}
	// TODO: you need to modify only this function and implement logic that will return false for 2 levels of tasks.
	process()
	return true
}

// HandleRequest2 runs the processes requested by users. Returns false if process had to be killed
func HandleRequest2(process func(), u *User) bool {
	if !u.IsPremium {
		c := time.Tick(10 * time.Millisecond)
		//select {
		//case <-c:
		//	if u.TimeUsed > almostTenSeconds {
		//		return false
		//	}
		//	u.TimeUsed += tenMillis
		//default:
		//}
		for _ = range c {
			if u.TimeUsed >= tenSeconds { //> almostTenSeconds {
				return false
			}
			u.TimeUsed += tenMillis
		}
	}
	// TODO: you need to modify only this function and implement logic that will return false for 2 levels of tasks.
	process()
	return true
}

func main() {
	RunMockServer()
}
