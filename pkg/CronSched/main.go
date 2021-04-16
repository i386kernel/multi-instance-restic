package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
)

func main(){
	log.Println("Starting....")
	c := cron.New()
	_, err := c.AddFunc("* * * * *", func(){
		log.Println("Hello World")
	})
	if err != nil{
	    log.Println(err)
	}
	c.Start()
	// c1 := cron.New()
	_, err = c.AddFunc("* * * * *", func(){
		anotherfunc()
	})
	if err != nil{
	    log.Println(err)
	}

	select {}
}

func anotherfunc(){
	fmt.Println("THis is another func")
}