package main

import (
	"context"
	"fmt"
)

func main(){

	gen := func(ctx context.Context) <-chan int{
		dst := make(chan int)
		n := 1

		go func() {
			for{
				select{
				case <-ctx.Done():
					fmt.Println("I am not returning")
				case dst <-n:
					n++
				}
			}
		}()
		return dst
	}

	ctx,cancel := context.WithCancel(context.Background())
	
	
	defer cancel()

	

	for n := range gen(ctx){
		fmt.Println(n)
		if n == 5{
			break
		}
	}
}