package main
import(
	"pressureClient/internal/api"
	"fmt"


)


func main(){
	client,err := api.NewClient()
	if err != nil{
		fmt.Printf("Failed to create client")
	}
	
	measurementsFromFile, err := loadMeasurements(client.DataFolder)
	_, err = client.PostMeasurementList(measurementsFromFile)
	

}


