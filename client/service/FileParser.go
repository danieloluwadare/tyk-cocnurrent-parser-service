package service

import (
	"bufio"
	"encoding/json"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/utils"
	"github.com/kpango/glg"
	"os"
)

type FileParser struct {
}

func NewFileParser() *FileParser {
	return &FileParser{}
}

//Parse returns an array of TykTaskConfig
func (f FileParser) Parse() ([]model.TykTaskConfig, error) {
	path := utils.GetDataPath()
	file, err := os.Open(path)

	if err != nil {
		glg.Error(err)
	}
	// remember to close the file at the end of the program
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	var tykTaskConfig []model.TykTaskConfig

	for scanner.Scan() {
		// do something with a line
		var tykConfig model.TykTaskConfig
		glg.Log(scanner.Text())
		if err := json.Unmarshal(scanner.Bytes(), &tykConfig); err != nil {
			glg.Error(err)
		}
		tykTaskConfig = append(tykTaskConfig, tykConfig)
	}

	if err := scanner.Err(); err != nil {
		glg.Error(err)
	}

	if len(tykTaskConfig) == 0 {
		glg.Error("FILE IS EMPTY")
	}

	return tykTaskConfig, nil
}
