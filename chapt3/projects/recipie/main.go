package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

const (
    dbUser = "yoshua"
    dbPass = "password"
    dbHost = "localhost"
    dbPort = "3306"
    dbName = "Go"
)

type Recipe struct {
	Name, Ingredients, Instructions string 
	PrepTime int

}

var db *sql.DB

func servePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	var current_recipe Recipe
	err := db.QueryRow("SELECT recipe_name, ingredients, instructions, prep_time FROM recipes WHERE slug =? OR recipe_id =?", slug, slug).Scan(&current_recipe.Name, &current_recipe.Ingredients, &current_recipe.Instructions, &current_recipe.PrepTime)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}else {
		html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>%s</title>
		</head>
		<body>
			<main>
				<h1 style="text-align:center">%s</h1>
				<h4>Ingredients: %s</h4>
				<p>Instructions: %s</p>
				<p>Preparation Time: %d minutes</p>
			</main>
		</body>
		</html>
	`, current_recipe.Name, current_recipe.Name, current_recipe.Ingredients, current_recipe.Instructions, current_recipe.PrepTime)

fmt.Fprintln(w, html)

	}

}

func main() {
	PORT := ":3000"
	var err error
	con_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = sql.Open("mysql", con_string)
	if err != nil {
		fmt.Println("Cant Open db",  err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/recipe/{slug}", servePage)
	fmt.Println("Server running on port =", PORT)
	http.ListenAndServe(PORT, r)
	
	
}