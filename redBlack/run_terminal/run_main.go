package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	runMain(10)
}

func runMain(number_times int) {

	for i := 0; i < number_times; i++ {
		cmd := exec.Command("go", "run", "../main.go")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(out.String())
	}
	fmt.Println("Finished running main function", number_times, "times.")
}
