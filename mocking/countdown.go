package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"
const sleep = "sleep"
const write = "write"

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type SpySleeper struct {
	Calls int
}
type DefaultSleeper struct {
}

type CountdownOperationSpy struct {
	Calls []string
}

func (c *CountdownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{1 * time.Second, time.Sleep})
}

func fizzBuzz(n int) []string {
	list := []string{}

	for i := 0; i < n; i++ {
		switch {
		case i%3 == 0 && i%5 != 0:
			list = append(list, "Fizz")
		case i%3 != 0 && i%5 == 0:
			list = append(list, "Buzz")
		case i%3 == 0 && i%5 == 0:
			list = append(list, "FizzBuzz")
		default:
			list = append(list, strconv.Itoa(i))
		}
	}

	return list
}
