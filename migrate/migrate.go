package main

import (
	"github.com/gitlord/REST_API_WITH_GORM/initializers"
	"github.com/gitlord/REST_API_WITH_GORM/models"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}
func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}