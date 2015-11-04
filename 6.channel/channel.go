package channel

// Create a channel to let Ping and Pong communicate and start both routines
func PingPong() {
}

// Method which sends 'true' every second
func Ping(c chan<- bool) {
}

// Method which reads from channel and display "PONG"
func Pong(c <-chan bool) {
}
