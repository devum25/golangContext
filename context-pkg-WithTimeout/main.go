package main

import (
	"context"
	"fmt"
	"time"
)


const shortDuration = 1 * time.Second

func main(){

   ctx,cancel := context.WithTimeout(context.Background(),shortDuration)
   defer cancel()

   select{
   case <-time.After(2 * time.Second):
	   fmt.Println("overslept")
   case <-ctx.Done():
	   fmt.Println(ctx.Err())
	        
   }
}