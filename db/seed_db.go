package db

import (
	"context"
)

func SeedDB() {
	ctx := context.Background()
	conn, err := ConnectToDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	sql := `
DROP TABLE IF EXISTS users;
	
CREATE TABLE users
	(
		id serial,
		name text,
		balance integer,
		group_id integer,
		PRIMARY KEY (id)
	);

INSERT INTO users (name, balance, group_id)
VALUES ('Bob', 100, 1),
       ('Alice', 100, 1),
       ('Eve', 100, 2),
       ('Mallory', 100, 2),
       ('Trent', 100, 3);
`
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		panic(err)
	}
}
