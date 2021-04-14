package main

import (
	"adverto/internal/api"
	"adverto/internal/repository"
	gorilla "adverto/internal/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

var schema = `
CREATE TABLE products (
    id INT NOT NULL AUTO_INCREMENT,
    article text,
    name text,
    amount integer,
	address text,
PRIMARY KEY (id)
);`

func main() {
db, err := sqlx.Connect("mysql", "root:helloworld@tcp(localhost:3308)/testapp")
	if err != nil {
		log.Fatalln(err)
	}
	repo := repository.NewLinksRepository(db)
	app := api.NewApplication(repo)
	router := gorilla.RegisterRouter()
	gorilla.RegisterHandlers(router, app)
	//_, err = db.Exec(schema)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	log.Fatal(http.ListenAndServe(":8080", router))

}
