package channelbis

// Create a channel to let Ping and Pong communicate and start both routines
// The function should return a channel which will stop the ping routine when closed
func PingPong() chan bool {
}

// Method which sends 'true' every second until the channel stop is closed
func Ping(c chan<- bool, stop chan bool) {
}

// Method which reads from channel and display "PONG"
// When channel is stopped, should display "END"
func Pong(c <-chan bool) {
}
