package main

import (
	"context"
	"fmt"

	db "github.com/krishnamdhanush/isolation-levels/db"
)

type Scenario struct {
	name            string
	f               func(string)
	isolationLevels []string
}

func main() {
	fmt.Println("Hello world")

	scenarios := make([]*Scenario, 0)
	scenarios = append(scenarios, &Scenario{f: db.DirtyRead, name: "Dirty Reads", isolationLevels: []string{"READ COMMITTED", "READ UNCOMMITTED"}})
	scenarios = append(scenarios, &Scenario{f: db.NonRepeatableRead, name: "Non Repeatable Reads", isolationLevels: []string{"READ COMMITTED", "REPEATABLE READ"}})
	scenarios = append(scenarios, &Scenario{f: db.PhantomReads, name: "Phantom Reads", isolationLevels: []string{"READ COMMITTED", "REPEATABLE READ"}})
	scenarios = append(scenarios, &Scenario{f: db.SerializableAnomaly, name: "Phantom Reads", isolationLevels: []string{"REPEATABLE READ", "SERIALIZABLE"}})

	for _, scenario := range scenarios {
		fmt.Println(scenario.name + "\n")
		for _, isolationLevel := range scenario.isolationLevels {
			fmt.Println(isolationLevel)
			db.SeedDB()
			scenario.f(isolationLevel)
			printTable()
		}
	}
}

func printTable() {
	ctx := context.Background()
	conn, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Final table state:\n")
	rows, _ := conn.Query(ctx, "SELECT id, name, balance, group_id FROM users ORDER BY id")
	for rows.Next() {
		var name []byte
		var id, balance, group_id int
		rows.Scan(&id, &name, &balance, &group_id)
		fmt.Printf("%2d | %10s | %5d | %d\n", id, name, balance, group_id)
	}
}
