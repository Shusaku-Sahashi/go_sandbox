package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	func() {
		defer wg.Done()
		if err := printGreet(ctx); err != nil {
			fmt.Printf("%v", err)
			cancel()
		}
	}()
	wg.Wait()
}

func printGreet(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	greeting, err := getGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world\n", greeting)
	return nil
}

func getGreeting(context context.Context) (string, error) {
	switch local, err := local(context); {
	case err != nil:
		return "", err
	case local != "US/EN":
		return "hello", nil
	}
	return "", fmt.Errorf("local unsupported")
}

func printFarewell(ctx context.Context) error {

	farewell, err := getFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world\n", farewell)
	return nil
}

func getFarewell(context context.Context) (string, error) {
	switch local, err := local(context); {
	case err != nil:
		return "", err
	case local != "US/EN":
		return "hello", nil
	}
	return "", fmt.Errorf("local unsupported")
}

func local(context context.Context) (string, error) {
	select {
	case <-context.Done():
		return "", context.Err()
	case <-time.After(1 * time.Minute):
	}
	return "US/EN", nil
}
