package main

import (
	"context"
	"database/sql"
	chapter1 "fitness.dev/app/gen"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Укажите строку подключения к базе данных
	dbURI := "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"

	// Открываем базу данных
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	// Проверка подключения
	if err := db.Ping(); err != nil {
		log.Fatalln("Error from database ping:", err)
	}

	// Создаем хранилище (store)
	st := chapter1.New(db)

	// Вставляем пользователя в базу
	user, err := st.CreateUsers(context.Background(), chapter1.CreateUsersParams{
		UserName:     "testuser",
		PassWordHash: "hash",
		Name:         "test",
	})
	if err != nil {
		log.Fatalln("Error creating user:", err)
	}

	// Если все успешно
	fmt.Printf("User created successfully: %+v\n", user)
}
