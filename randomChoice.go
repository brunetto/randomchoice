package main

import( 
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main () () {
	if len(os.Args) < 3 {
		log.Println("Please provide two choice with quotes between you need to choose!")
		log.Fatal("For example, ask 'randomChoice \"to study\" \"not to study\"")
	}
	var (
		extractions, sum, choice float64
	)
	extractions = 10000000.
	rand.Seed(time.Now().Unix())
	sum = 0.
	for i:=0.; i<extractions;i++ {
		sum = sum + float64((rand.Intn(2)))
	}
	choice = sum / extractions
	if choice < 0.5 {
		fmt.Print(os.Args[1])
	} else {
		fmt.Print(os.Args[2])
	}
	fmt.Println(" over ", extractions, " extractions")
}


