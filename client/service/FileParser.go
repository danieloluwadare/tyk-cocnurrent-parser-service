package service

import (
	"bufio"
	"encoding/json"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/utils"
	"github.com/kpango/glg"
	"log"
	"os"
)

type FileParser struct {
}

func NewFileParser() *FileParser {
	return &FileParser{}
}

func (f FileParser) Parse() ([]model.TykTaskConfig, error) {
	path := utils.GetDataPath()
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	//var sliceOfByte [][]byte
	var tykTaskConfig []model.TykTaskConfig

	for scanner.Scan() {
		// do something with a line
		var tykConfig model.TykTaskConfig
		//sliceOfByte = append(sliceOfByte, scanner.Bytes())
		glg.Log(scanner.Text())
		if err := json.Unmarshal(scanner.Bytes(), &tykConfig); err != nil {
			log.Fatal(err)
		}
		tykTaskConfig = append(tykTaskConfig, tykConfig)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	glg.Log(tykTaskConfig)
	glg.Log(tykTaskConfig[0].Name)

	return tykTaskConfig, nil
}
