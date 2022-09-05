package main

import (
	"fmt"
)

// Option 1: using range in consumer
func producerOption1(stream Stream) chan *Tweet {
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

func consumerOption1(tw chan *Tweet) {
	for tweet := range tw {
		if tweet.IsTalkingAboutGo() {
			fmt.Println(tweet.Username, "\ttweets about golang")
			continue
		}

		fmt.Println(tweet.Username, "\tdoes not tweet about golang")
	}
}

// Option 2: using select in consumer
func producerOption2(done chan struct{}, stream Stream) chan *Tweet {
	tw := make(chan *Tweet)
	go func() {
		for {
			tweet, err := stream.Next()
			if err == ErrEOF {
				close(tw)
				close(done)
				return
			}
			tw <- tweet
		}
	}()
	return tw
}

func consumerOption2(done chan struct{}, tw chan *Tweet) {
	for {
		select {
		case tweet := <-tw:
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

// Task
func producer(stream Stream) (tweets []*Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return tweets
		}
		// TODO: use channel here
		tweets = append(tweets, tweet)
	}
}

func consumer(tweets []*Tweet) {
	// TODO: use channel here
	for _, t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
			continue
		}

		fmt.Println(t.Username, "\tdoes not tweet about golang")
	}
}

func main() {
	// TODO ... stub
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
