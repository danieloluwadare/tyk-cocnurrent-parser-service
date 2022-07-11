package service

import (
	"sync"
)

type Initializer struct {
	parser           Parser
	afterExtractions []AfterExtraction
}

func NewInitializer(parser Parser, afterExtractions ...AfterExtraction) *Initializer {
	initializedAfterExtractions := make([]AfterExtraction, 0)
	for _, afterExtraction := range afterExtractions {
		initializedAfterExtractions = append(initializedAfterExtractions, afterExtraction)
	}
	return &Initializer{parser: parser, afterExtractions: initializedAfterExtractions}
}

func (i Initializer) Execute() {
	var wg sync.WaitGroup
	tykTaskConfigs, _ := i.parser.Parse()

	for _, afterExtraction := range i.afterExtractions {
		afterExtraction.ExecuteAsync(tykTaskConfigs, &wg)
	}
	wg.Wait()
}