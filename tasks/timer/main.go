package main

// Trying out the select statement.

// 1. Create a timer that fires after 1 second.

// 2. Create a goroutine that passes the values 1, 2, 3 on a int channel
// successively - with a 200ms sleep after each result.

// 3. Create a select statement (wrapped in a for loop) that has three cases:
// one for the timer, one for the channel sending int (1, 2, 3) and one for a
// `time.Tick` channel, that fires every e.g. 100ms.

func main() {

}
