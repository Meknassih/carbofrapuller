package main

import (
	"fmt"
	"os"
	"time"
)

var data_URL = os.Getenv("CARBOFRA_DATA_URL")
var mongo_URL = os.Getenv("CARBOFRA_MONGO_URL")
var wait_time = os.Getenv("CARBOFRA_WAIT_TIME")

func main() {
	start_time := time.Now()
	fmt.Println("Starting data import at", start_time.Format("15:04:05"))
	defer func() {
		fmt.Println("Done at", time.Now().Format("15:04:05"), "in", time.Since(start_time).String())
	}()

	read_closer, err := Get_all_data()
	if err != nil {
		fmt.Println("Error while retrieving data", err)
		return
	}
	defer read_closer.Close()
	
	var unzipped []byte
	err = Unzip(read_closer, &unzipped)
	if err != nil {
		fmt.Println("Error while unzipping data", err)
		return
	}
	
	var xml_data ListePDV
	err = Parse_XML(unzipped, &xml_data)
	if err != nil {
		fmt.Println("Error while parsing XML", err)
		return
	}
	fmt.Println("Number of stations: ", len(xml_data.PDVs))
	
	err = Insert_PDVs(xml_data.PDVs)
	if err != nil {
		fmt.Println("Error while inserting into MongoDB", err)
		return
	}
	
	if wait_time != "" {
		duration, err := time.ParseDuration(wait_time)
		if err != nil {
			fmt.Println("Error while parsing wait time", err)
			return
		}
		fmt.Println("Waiting for", duration.String())
		time.Sleep(duration)
	}
}