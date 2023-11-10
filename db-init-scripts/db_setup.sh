psql -U postgres -d postgres -c "\c postgres";
psql -U postgres -d postgres -c "CREATE table if not exists users(id  serial primary key ,name text, balance integer, group_id integer)";
psql -U postgres -d postgres -c "INSERT INTO users(name, balance, group_id) VALUES ('Bob', 100, 1), ('Alice', 100, 1), ('Eve', 100, 2), ('Mallory', 100, 2), ('Trent', 100, 3);"; 
psql -U postgres -d postgres -c "SELECT * FROM users;"; 
