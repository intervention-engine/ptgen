package ptgen

import (
	"bytes"
	"encoding/json"

	"github.com/intervention-engine/fhir/models"
)

type MedicationMetadata struct {
	ID         int    `json:"medication_id"`
	RxNormCode string `json:"rxNormCode"`
	BrandName  string `json:"brandName"`
}

func GenerateMedication(medicationID int, onset *models.FHIRDateTime, mmd []MedicationMetadata) *models.MedicationStatement {
	if medicationID == 0 {
		return nil
	} else {
		ms := &models.MedicationStatement{}
		mmd := medicationByID(medicationID, mmd)
		ms.EffectivePeriod = &models.Period{Start: onset}
		ms.MedicationCodeableConcept = &models.CodeableConcept{Coding: []models.Coding{{Code: mmd.RxNormCode, System: "http://www.nlm.nih.gov/research/umls/rxnorm"}}, Text: mmd.BrandName}
		return ms
	}
}

func LoadMedications() []MedicationMetadata {
	j, err := Asset("data/medications.json")
	if err != nil {
		panic("Can't get the medications data")
	}
	decoder := json.NewDecoder(bytes.NewReader(j))
	md := []MedicationMetadata{}
	decoder.Decode(&md)
	return md
}

func medicationByID(id int, md []MedicationMetadata) *MedicationMetadata {
	for _, m := range md {
		if m.ID == id {
			return &m
		}
	}
	return nil
}
