package service_test

import (
	_ "errors"
	"fmt"
	model2 "github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/service"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/mocks"
	_ "github.com/jinzhu/gorm"
	_ "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//Todo: modify CreateUser test
func TestService_CreateUser(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipped test: The test fails cos it hasn't been implemented")
	}

	task1 := model2.TykTaskConfig{"Task1", "Description"}

	testCases := []struct {
		name           string
		tykTaskConfigs []model2.TykTaskConfig
		task           []model2.Task
	}{
		{
			"Test with valid a slice of tykTaskConfigs input",
			[]model2.TykTaskConfig{task1},
			[]model2.Task{{TykTaskConfig: task1}},
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			fmt.Println(time.Now())
			// Create dependency userRepo with mock implementation
			iTaskRepository := mocks.ITaskRepository{}
			iTaskRepository.On("Save", testCase.task[0]).Return(nil, nil)

			// Create userService and inject mock repo
			databaseAfterExtractionService := service.NewDatabaseAfterExtractionService(&iTaskRepository)

			databaseAfterExtractionService.Execute(testCase.tykTaskConfigs)

			iTaskRepository.AssertCalled(t, "Save", testCase.task[0])
		})
	}
}
