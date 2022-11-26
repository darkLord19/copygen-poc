package main

import (
	"log"

	"github.com/darkLord19/copygen-poc/copygen"
	"github.com/darkLord19/copygen-poc/models"
)

func main() {
	o := models.OKR{
		XId:                  "objectiveId",
		Name:                 "TestCopygenO",
		HasChildren:          false,
		Type:                 "objective",
		KeyResultIds:         []string{},
		ContributeToProgress: true,
		Outcome: models.KeyResultOutcome{
			StartValue: 1,
			EndValue:   100,
		},
	}

	kr := models.OKR{
		XId:                  "keyResultId",
		Name:                 "TestCopygenKR",
		HasChildren:          false,
		Type:                 "keyResult",
		KeyResultIds:         []string{},
		ContributeToProgress: true,
		Outcome: models.KeyResultOutcome{
			StartValue: 1,
			EndValue:   100,
		},
	}
	var objective models.Objective
	var keyResult models.KeyResult
	copygen.OKRToObjective(&objective, o)
	copygen.OKRToKeyResult(&keyResult, kr)
	log.Println("objective is", objective)
	log.Println("key result is", keyResult)
}
