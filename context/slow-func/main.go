package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "userID", 10)
	val, err := fetchUserData(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result: ", val)
	fmt.Println("took: ", time.Since(start))

}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context) (int, error) {
	userID := ctx.Value("userID").(int)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow(userID)
		respch <- Response{value: val, err: err}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fecthing data from third party took to long")
		case resp := <-respch:
			return resp.value, resp.err
		}

	}

}

func fetchThirdPartyStuffWhichCanBeSlow(userID int) (int, error) {
	time.Sleep(time.Millisecond * 150)

	return 666, nil
}
