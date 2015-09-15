package ptgen

type ConditionMetadata struct {
	ID                   int    `json:"condition_id"`
	ICD9                 string `json:"icd9code"`
	Display              string `json:"display"`
	MedicationID         int    `json:"medication_id"`
	Overnights           string `json:"overnights"`
	AbatementChance      int    `json:"abatementChance"`
	Fatal                bool   `json:"healOrDeath"`
	MortalityChance      int    `json:"mortalityChance"`
	MortalityTime        string `json:"mortalityTime"`
	RecoveryEstimate     string `json:"recoveryEstimate"`
	ProcedureChance      int    `json:"procedureChance"`
	ProcedureSuccess     int    `json:"procedureSuccess"`
	CheckUp              string `json:"checkUp"`
	ProcedureDescription string `json:"procedureDescription"`
	ProcedureCode        string `json:"procedureCode"`
	ProcedureName        string `json:"procedureCodeName"`
}
