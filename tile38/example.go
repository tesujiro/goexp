package main

// https://github.com/tidwall/tile38/wiki/Go-example-(redigo)1:w

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":9851")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer c.Close()

	ret, _ := c.Do("SET", "fleet", "truck1", "POINT", "33", "-115")
	fmt.Printf("%s\n", ret)

	ret, _ = c.Do("GET", "fleet", "truck1")
	fmt.Printf("%s\n", ret)

}
