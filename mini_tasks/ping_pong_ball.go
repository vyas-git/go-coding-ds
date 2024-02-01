package main

// Ping pong game using channels and routines.
// There are two players playing ping pong game.
// This game has to go on for 2 minutes.
// You get the ball, wait for random time (0-10 seconds) and serve back.
// Each player serves at random intervals. At the end of 2 minutes, whoever has the ball is the winner.
// Please use context and routines
import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ball := make(chan int, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// start game
	ball <- 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				{
					break
				}
			case <-ball:
				{
					fmt.Println("Ball is in player 1 sending...")
					time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
					ball <- 1
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				{
					break
				}
			case <-ball:
				{
					fmt.Println("Ball is player 2 sending...")
					time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
					ball <- 2
				}
			}
		}
	}()

	select {
	case <-ctx.Done():
		{
			fmt.Println("Game Over Winner Is ", <-ball)
		}

	}
}
