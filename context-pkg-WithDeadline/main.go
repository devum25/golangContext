package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 5 * time.Second

func main(){
   d := time.Now().Add(shortDuration)

   ctx,cancel := context.WithDeadline(context.Background(),d)

    v,ok := ctx.Deadline()
	if ok{
		fmt.Printf("time started:%v\n",time.Now())
		fmt.Printf("time given to execute:%v\n",v)
	}

   defer cancel()

   select{
   case <-time.After(2 * time.Second):
	      fmt.Println("overslept")
   case <-ctx.Done():
	   fmt.Println("Done channel closed")
	   fmt.Println(ctx.Err())
   }
}