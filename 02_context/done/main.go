package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})

	defer close(done)

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreeting(done); err != nil {
			fmt.Errorf("%v", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printFarewell(done); err != nil {
			fmt.Errorf("%v", err)
			return
		}
	}()

	wg.Wait()
}

func printGreeting(done <-chan struct{}) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func genGreeting(done <-chan struct{}) (string, error) {
	switch local, err := local(done); {
	case err != nil:
		return "", err
	case local == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func printFarewell(done <-chan struct{}) error {
	farewell, err := genFarewell(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genFarewell(done <-chan struct{}) (string, error) {
	switch local, err := local(done); {
	case err != nil:
		return "", err
	case local == "EN/US":
		return "goodby", nil
	}
	return "", fmt.Errorf("unsuported local")
}

func local(done <-chan struct{}) (string, error) {
	select {
	case <-done:
		return "", fmt.Errorf("calceled")
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}
