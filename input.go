package main
	import(
	"pressureClient/internal/api"
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"	
)	
func loadMeasurements(folderpath string) ([]api.CreateMeasurementJSON, error){
	sliceMeasurementJSON := []api.CreateMeasurementJSON{}
	//List the files in given folder path	
	files, err := os.ReadDir(folderpath)	
	if err != nil{
		return []api.CreateMeasurementJSON{}, err
	}

	//open the first file, logic to only open files not used by
	//the measurement equipment will go here
	
	file, err := os.Open(folderpath +"/"+ files[0].Name())	
	if err != nil{
		return []api.CreateMeasurementJSON{}, err
	}
	defer file.Close()
	//Read the file line by line
	scanner := bufio.NewScanner(file)	
	//While the scanner is retrieving lines give me the current line as a string
	// with scanner.Text()
	for scanner.Scan(){
		splitLine := strings.Split(scanner.Text()," ")
			
		newMeasurement := api.CreateMeasurementJSON{}
			//Conv the strings to numbers 
			conPressure1, err := strconv.ParseFloat(splitLine[3],64)
				if err != nil{	
					log.Printf("Error converting pressure 1: %v",err)
					continue		
				
			}
			
			conPressure2, err := strconv.ParseFloat(splitLine[4],64)
				if err != nil{
					log.Printf("Error converting pressure 2: %v", err)
					continue	
			}
			
			conTemperature1, err := strconv.ParseFloat(splitLine[5],64)
				if err != nil{
					log.Printf("Error converting Temperature 1: %v",err)
				continue	
			}
			
			conTemperature2, err := strconv.ParseFloat(splitLine[6],64)
				if err != nil{
					log.Printf("Error converting Temperature 2: %v", err)
				continue	
			}
			


			//Assign our various strings to the struct

			newMeasurement.MeasurementDate = splitLine[1]
			newMeasurement.MeasurementTime = splitLine[2]
			newMeasurement.Pressure1 = float64(conPressure1)
			newMeasurement.Pressure2 = float64(conPressure2)
			newMeasurement.Temperature1 = float64(conTemperature1)
			newMeasurement.Temperature2 = float64(conTemperature2)
		sliceMeasurementJSON = append(sliceMeasurementJSON,newMeasurement)
	}
	//Checks for errors during the scanning process
	if err := scanner.Err(); err != nil {
        	return []api.CreateMeasurementJSON{}, err
    }
	return sliceMeasurementJSON, nil
}
