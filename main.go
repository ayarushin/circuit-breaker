package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	circuit := Breaker(func(ctx context.Context) (string, error) {
		if time.Now().Second()%2 == 0 {
			return "", errors.New("simulated error")
		}

		return "Success", nil
	}, 2)

	for i := 0; i < 5; i++ {
		response, err := circuit(context.Background())
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Response: %v\n", response)
		}

		time.Sleep(time.Second)
	}

}
