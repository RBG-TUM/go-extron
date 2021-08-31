package main

import (
	"fmt"
	goextron "github.com/RBG-TUM/go-extron"
	"time"
)

//main authenticates with a smp, mutes the video and audio, waits a second and unmutes it.
func main() {
	ge := goextron.New("http://1.2.3.4/", "admin", "123456")
	err := ge.GetAuthStatus()
	if err != nil {
		fmt.Printf("Authentication failure: %v\n", err)
		return
	}
	err = ge.SetMute(true)
	if err != nil {
		fmt.Printf("unable to mute: %v\n", err)
		return
	}
	time.Sleep(time.Second*5)
	err = ge.SetMute(false)
	if err != nil {
		fmt.Printf("unable to unmute: %v\n", err)
		return
	}
}
