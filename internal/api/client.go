package api 
import (
	"net/http"
	"github.com/joho/godotenv"
	"os"
	"time"
	"fmt"
	"github.com/google/uuid"
	"encoding/json"
	"bytes"
)

const(
	MeasurementsEndpoint = "/api/measurements"
)

type Client struct{
	BaseURL string
	BaseClient *http.Client
	ApiKey string
	DataFolder string
}

type CreateMeasurementJSON struct {
	MeasurementDate string `json:"measurement_date"`
	MeasurementTime string `json:"measurement_time"`
	Pressure1       float64	  `json:"pressure_1"`
	Pressure2      	float64	  `json:"pressure_2"`
	Temperature1    float64	`json:"temperature_1"`
	Temperature2    float64	  `json:"temperature_2"`
}
type MeasurementJSON struct {
	ID              uuid.UUID `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	MeasurementDate time.Time `json:"measurement_date"`
	MeasurementTime time.Time `json:"measurement_time"`
	Pressure1       float64 `json:"pressure_1"`
	Pressure2       float64 `json:"pressure_2"`
	Temperature1    float64 `json:"temperature_1"`
	Temperature2    float64 `json:"temperature_2"`
}




func NewClient()(*Client, error){	
	//Load env file and set env variables
	godotenv.Load()
	loadedDataFolder := os.Getenv("DATA_FOLDER")
	if loadedDataFolder == ""{
		return nil, fmt.Errorf("Failed to load Data folder path")
	}
	loadedServerURL := os.Getenv("SERVER_URL")
	if loadedServerURL == ""{
		return nil, fmt.Errorf("Failed to load env server URL")
	}
	loadedApiKey := os.Getenv("TEST_API_KEY")
		if loadedApiKey == ""{
			return nil, fmt.Errorf("Failed to load env api key")
	}
	// Create a client struct with the env variables
	c :=&Client{
		BaseURL: loadedServerURL,
		BaseClient: &http.Client{
		Timeout: time.Second * 30,
		},
		ApiKey: loadedApiKey,
		DataFolder: loadedDataFolder,
	}
	return c,nil
}

func (c *Client)PostMeasurementList(measurementList []CreateMeasurementJSON)([]MeasurementJSON,error){
	// Marshal measurementList to JSON
	jsonData, err := json.Marshal(measurementList)
		if err != nil{
			return []MeasurementJSON{}, err			
	}
	// Create a post request
	req, err := http.NewRequest("POST",c.BaseURL + MeasurementsEndpoint, bytes.NewBuffer(jsonData))
	if err != nil{
		return []MeasurementJSON{}, err
	}
	// Set Request headers
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Authorization","ApiKey "+ c.ApiKey)	
	//Send the request
	res, err := c.BaseClient.Do(req)
	if err != nil{
		fmt.Println(err)
		return []MeasurementJSON{},err
	}
	defer res.Body.Close()

	// Decode the response
	var createdMeasurements []MeasurementJSON
	decoder := json.NewDecoder(res.Body)
	fmt.Printf("res.body : %v",res.Body)
	err = decoder.Decode(&createdMeasurements)
	if err != nil{
		return []MeasurementJSON{},err
	}
	return createdMeasurements, nil
}


