package db

import (
	"context"
	"fmt"
)

func DirtyRead(isolationLevel string) {
	ctx := context.Background()
	conn1, err := ConnectToDB()
	if err != nil {
		panic(err)
	}
	defer conn1.Close(ctx)
	conn2, err := ConnectToDB()
	if err != nil {
		panic(err)
	}
	defer conn2.Close(ctx)
	tx, err := conn1.Begin(ctx)
	if err != nil {
		panic(err)
	}
	tx.Exec(ctx, "SET TRANSACTION ISOLATION LEVEL "+isolationLevel)

	_, err = tx.Exec(ctx, "UPDATE users SET balance = 256 WHERE name='Bob'")
	if err != nil {
		fmt.Printf("Failed to update Bob balance in tx: %v\n", err)
	}

	var balance int
	row := tx.QueryRow(ctx, "SELECT balance FROM users WHERE name='Bob'")
	row.Scan(&balance)
	fmt.Printf("Bob balance from main transaction after update: %d\n", balance)

	row = conn2.QueryRow(ctx, "SELECT balance FROM users WHERE name='Bob'")
	row.Scan(&balance)
	fmt.Printf("Bob balance from concurrent transaction: %d\n\n", balance)

	if err := tx.Commit(ctx); err != nil {
		fmt.Printf("Failed to commit: %v\n", err)
	}

}
