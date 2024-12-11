package main

import (
	"fmt"

	postgrest "github.com/supabase-community/postgrest-go"
)

var (
	REST_URL = `http://localhost:3000`
)

func main() {
	client := postgrest.NewClient(REST_URL, "", nil)
	if client.ClientError != nil {
		panic(client.ClientError)
	}

	result, err := client.Rpc("add_them", "", map[string]int{"a": 9, "b": 3})
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
