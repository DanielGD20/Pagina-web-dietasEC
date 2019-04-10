package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	qr, err := db.Query("INSERT INTO dietasec.recetas( 'idRecetas', nombre_receta, descripcion) VALUES (001,'Pollo', 'se prepara el arroz con pollo con pollo');")
	/*qr, err := db.Query("INSERT INTO dietasec.ingredientes(nombre_ingrediente, cantidad_ingrediente, 'idRecetas') VALUES ('pollo', 1, 001);")
	qr, err := db.Query("INSERT INTO ingredientes (idRecetas, nombre_ingrediente, cantidad_ingrediente) VALUES (001, 'sopa instantanea', 1);")
	qr, err := db.Query("INSERT INTO ingredientes (idRecetas, nombre_ingrediente, cantidad_ingrediente) VALUES (001, 'tomate', 2);")*/

	fmt.Printf("Query Result: %v\n", qr)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("paginas/index.html")
	router.Static("/css", "css")
	router.Static("/js", "js")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(":" + port)
}

/*func enviar(w http.ResponseWritter, r*http.Request){
	fmt.Println(r.Method)
	if r.Method == "GET"{
		t,_:=template.ParseFiles("index.html")
		t.Execute(w,nil)

	}else{
		fmt.Println(r.Method)
		r.ParseForm()
		fmt.Println(r.Form[])
		fmt.Println(r.Form[])
	}
}*/
