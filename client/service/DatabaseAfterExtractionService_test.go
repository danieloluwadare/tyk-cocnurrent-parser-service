package service_test

import (
	"bou.ke/monkey"
	_ "errors"
	model2 "github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/model"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/client/service"
	"github.com/danieloluwadare/tyk-cocnurrent-parser-service/mocks"
	"github.com/jinzhu/gorm"
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
	timeTestInputForRepo := time.Now()

	//timeTest2 := time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return timeTestInputForRepo
	})
	task1 := model2.TykTaskConfig{"Task1", "Description"}
	//task2 := model2.TykTaskConfig{"Task2", "Description"}

	testCases := []struct {
		name           string
		tykTaskConfigs []model2.TykTaskConfig
		task           []model2.Task
	}{
		{
			"Test with valid a alice of tykTaskConfigs input",
			[]model2.TykTaskConfig{task1},
			[]model2.Task{{Model: gorm.Model{CreatedAt: timeTestInputForRepo, UpdatedAt: timeTestInputForRepo}, TykTaskConfig: task1}},
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			// Create dependency userRepo with mock implementation
			iTaskRepository := mocks.ITaskRepository{}
			iTaskRepository.On("Save", testCase.task[0]).Return(nil, nil)
			//iTaskRepository.On("Save", testCase.task[1]).Return(nil, nil)

			// Create userService and inject mock repo
			databaseAfterExtractionService := service.NewDatabaseAfterExtractionService(&iTaskRepository)

			// Actual method call
			//output, _ := userService.CreateUser(testCase.repoInputData)
			databaseAfterExtractionService.Execute(testCase.tykTaskConfigs)

			monkey.Unpatch(time.Now)

			// Expected output
			//expected := testCase.expectedVal
			//
			//assert.Equal(t, expected, output)
			iTaskRepository.AssertCalled(t, "Save")
		})
	}
}
