package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://raw.githubusercontent.com/whiteKnightz/learning-go/main/main/data/testFile.json"

func main() {
	resp, err := http.Get(url)
	checkErr(err)
	fmt.Printf("Response type: %T \n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	content := string(bytes)
	fmt.Printf("Response type: %v \n", content)

	analysisData := analysisDataFromJson(content)

	fmt.Println(analysisData)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func analysisDataFromJson(content string) AnalysisData {
	//decoder := json.NewDecoder(strings.NewReader(content))
	//_, err := decoder.Token()
	//checkErr(err)
	//var data AnalysisData
	//for decoder.More() {
	//	err := decoder.Decode(&data)
	//	checkErr(err)
	//}
	//return data
	var data AnalysisData
	err := json.Unmarshal([]byte(content), &data)
	checkErr(err)
	return data
}

type AnalysisData struct {
	Name, LabelCol, RequestType string
	InferSchema, HasHeaders     bool
	Features                    []string
}
