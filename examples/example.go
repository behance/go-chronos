package main

import (
	"fmt"

	chronos "github.com/behance/go-chronos/chronos"
)

func main() {
	config := chronos.NewDefaultConfig()
	client, err := chronos.NewClient(config)

	fmt.Printf("Client: %+v\n", client)
	fmt.Printf("Err: %+v\n", err)

	jobs, err := client.Jobs()
	fmt.Printf("Jobs: %+v\n", jobs)
	fmt.Printf("Err: %+v\n", err)
}
