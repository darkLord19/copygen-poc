package models

type KeyResultOutcome struct {
	StartValue int32
	EndValue   int32
}

type OKR struct {
	XId         string
	Name        string
	HasChildren bool
	Type        string

	// objective only fields
	KeyResultIds         []string
	ContributeToProgress bool

	// key result only fields
	Outcome KeyResultOutcome
}

type Objective struct {
	XId                  string
	Name                 string
	HasChildren          bool
	Type                 string
	KeyResultIds         []string
	ContributeToProgress bool
}

type KeyResult struct {
	XId         string
	Name        string
	HasChildren bool
	Type        string
	Outcome     KeyResultOutcome
}
