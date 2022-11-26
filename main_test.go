package main

import (
	"testing"

	"github.com/darkLord19/copygen-poc/copygen"
	"github.com/darkLord19/copygen-poc/models"
	"github.com/jinzhu/copier"
)

func BenchmarkObjectiveConversionCopyegen(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		var objective models.Objective
		copygen.OKRToObjective(&objective, o)
	}
}

func BenchmarkKeyresultConversionCopygen(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		var keyResult models.KeyResult
		copygen.OKRToKeyResult(&keyResult, kr)
	}
}

func BenchmarkObjectiveConversionCopier(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		var objective models.Objective
		copier.Copy(&o, &objective)
	}
}

func BenchmarkKeyresultConversionCopier(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		var keyResult models.KeyResult
		copier.Copy(&kr, &keyResult)
	}
}
