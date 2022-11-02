package main

import (
	"fmt"
)

// Option 1: using range in consumer
func producer1(stream Stream) chan *Tweet {
	tw := make(chan *Tweet)
	go func() {
		for {
			tweet, err := stream.Next()
			if err == ErrEOF {
				close(tw)
				return
			}
			tw <- tweet
		}
	}()
	return tw
}

func consumer1(tw chan *Tweet) {
	for tweet := range tw {
		if tweet.IsTalkingAboutGo() {
			fmt.Println(tweet.Username, "\ttweets about golang")
			continue
		}

		fmt.Println(tweet.Username, "\tdoes not tweet about golang")
	}
}

// Option 2: using select in consumer
func producer2(done chan struct{}, stream Stream) chan *Tweet {
	tw := make(chan *Tweet)
	go func() {
		for {
			tweet, err := stream.Next()
			if err == ErrEOF {
				close(tw)
				return
			}
			tw <- tweet
		}
	}()
	return tw
}

func consumer2(done chan struct{}, tw chan *Tweet) {
	for {
		select {
		case tweet := <-tw:
			// check default value got(channel closed)
			if tweet == nil {
				return
			}
			if tweet.IsTalkingAboutGo() {
				fmt.Println(tweet.Username, "\ttweets about golang")
			} else {
				fmt.Println(tweet.Username, "\tdoes not tweet about golang")
			}
		case done <- struct{}{}:
			return
		}
	}
}

func main() {
	// check main_test.go
}

// Original main
//
//func main() {
//	start := time.Now()
//	stream := GetMockStream()
//
//	// Modification starts from here
//	// Hint: this can be resolved via channels
//	// Producer
//	tweets := producer(stream)
//	// Consumer
//	consumer(tweets)
//
//	fmt.Printf("Process took %s\n", time.Since(start))
//}
