package main

import (
	"fmt"

	"github.com/StarGazer500/Department-Management-WebApp/internals/database"
	
	"github.com/StarGazer500/Department-Management-WebApp/internals/server"
// "github.com/joho/godotenv"
	"github.com/StarGazer500/Department-Management-WebApp/internals/config"
)


func init() {

    err1 := config.LoadEnv()
	if err1!=nil{
		fmt.Errorf("error loading Env: %v", err1)
		
	}

	// err1 := godotenv.Load("../.env")
	// if err1 != nil {
	// 	fmt.Println("Error loading .env file")
	// 	// return errors.New("failed to load environment variables") // Corrected error creation
	// }

	err := database.ConnectTODb()
	if err!=nil{
		fmt.Errorf("error gconnecting to the database: %v", err)
		
	}
}

func main(){
	defer database.Dbinstance.Close()
	server.NewServer().Run()
	
	
}
