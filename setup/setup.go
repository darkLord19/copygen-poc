// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/darkLord19/copygen-poc/models"
)

// Copygen defines the functions that will be generated.
type Copygen interface {
	OKRToObjective(models.OKR) *models.Objective
	OKRToKeyResult(models.OKR) *models.KeyResult
}
