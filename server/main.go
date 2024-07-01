package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Todo struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
	Body string `json:"body"`
}

func main() {
	
	fmt.Printf("Hello")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/healthcheck",func(c *fiber.Ctx) error{
		return c.SendString("Hello from Golang")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error{
		todo := &Todo{}
		err := c.BodyParser(todo)
		if err != nil{
			return c.Status(400).SendString("Input can't be parsed,Invalid format")
		}
		
		todo.Id = len(todos) +1
		todos = append(todos,*todo)
		fmt.Println("title:" , todo.Title , "body: " , todo.Body, todo.Id , todo.Status)
		return c.JSON(todo);
	})

	app.Patch("/api/:id/status",func(c *fiber.Ctx) error {
		id,err := c.ParamsInt("id")
		if err != nil{
			return c.Status(401).SendString("Invalid id")
		}
		for i,todo := range todos{
			if todo.Id== id {
				todos[i].Status = true
				fmt.Println("Patched: " , "title: " , todo.Title , "body: " , todo.Body, todo.Id , todo.Status)
				break
			}
		}
		
		return c.JSON(todos)
	})

	app.Get("/api/todos",func(c *fiber.Ctx) error{
		return c.JSON(todos)
	})


	panic(app.Listen((":8080")))
}