package main

func main() {
	c := CreateCollection("Go Collection", "Awesome description")

	c.AddItemGroup("This is a folder").AddItem(&Item{
		Name: "An item inside a folder",
	})

	c.AddItem(&Item{
		Name:    "This is a request",
		Request: NewRequest("http://www.google.fr", "GET"),
	})

	c.AddItemGroup("Empty folder")

	c.Write("postman_collection.json")
}
