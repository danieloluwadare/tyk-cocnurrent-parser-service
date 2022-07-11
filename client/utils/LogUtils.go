package utils

import (
	"github.com/kpango/glg"
	"log"
	"os"
	"path/filepath"
)

func InitializeLog() {
	workingDirectoryPath, err := os.Getwd()
	if err != nil {
	}
	logPath := os.Getenv("LOG.NAME")
	fullPathToLogFile := filepath.Join(workingDirectoryPath, "logs", logPath)

	_, err = os.Stat(fullPathToLogFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("File does not exist.")
			_ = os.MkdirAll(filepath.Join(workingDirectoryPath, "logs"), 0700)
			newFile, err := os.Create(fullPathToLogFile)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(newFile)
			newFile.Close()
		}
	}
	log.Println("Log path:", fullPathToLogFile)

	log.Println("File does exist. File information:")
	//log.Println(fileInfo.Name())

	infoLog := glg.FileWriter(fullPathToLogFile, 0666)

	//defer infoLog.Close()
	glg.Get().
		SetMode(glg.BOTH). // default is STD
		DisableColor().
		AddLevelWriter(glg.INFO, infoLog).  // add info log file destination
		AddLevelWriter(glg.DEBG, infoLog).  // add info log file destination
		AddLevelWriter(glg.FAIL, infoLog).  // add info log file destination
		AddLevelWriter(glg.FATAL, infoLog). // add info log file destination
		AddLevelWriter(glg.WARN, infoLog).  // add info log file destination
		AddLevelWriter(glg.LOG, infoLog).   // add info log file destination
		AddLevelWriter(glg.OK, infoLog).    // add info log file destination
		AddLevelWriter(glg.ERR, infoLog)    // add error log file destination
}
