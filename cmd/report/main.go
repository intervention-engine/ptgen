package main

import (
	"flag"
	"html/template"
	"os"
	"sort"
	"strings"

	"github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/ptgen"
)

func main() {
	num := flag.Int("n", 100, "Number of patients to generate")
	flag.Parse()
	count := *num
	resources := make([][]interface{}, count)
	for i := 0; i < count; i++ {
		resources[i] = ptgen.GeneratePatient()
	}
	generateReport(resources)
}

func generateReport(resources [][]interface{}) {
	reports := make([]*PatientReportData, len(resources))
	for i := range resources {
		reports[i] = NewPatientReportData(resources[i])
	}
	sort.Sort(PatientReportDataByFamilyName(reports))

	t, err := template.ParseFiles("report.template.html")
	if err != nil {
		panic(err)
	}

	data := struct {
		PatientInfos []*PatientReportData
	}{
		reports,
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

type PatientReportData struct {
	Patient         models.Patient
	Conditions      []models.Condition
	Medications     []models.MedicationStatement
	Encounters      []models.Encounter
	EncObservations map[string][]models.Observation
}

func NewPatientReportData(resources []interface{}) *PatientReportData {
	d := new(PatientReportData)
	d.EncObservations = make(map[string][]models.Observation)
	for _, r := range resources {
		switch r := r.(type) {
		case *models.Patient:
			d.Patient = *r
		case *models.Condition:
			d.Conditions = append(d.Conditions, *r)
		case *models.MedicationStatement:
			d.Medications = append(d.Medications, *r)
		case *models.Encounter:
			d.Encounters = append(d.Encounters, *r)
		case *models.Observation:
			enc := strings.TrimPrefix(r.Encounter.Reference, "cid:")
			d.EncObservations[enc] = append(d.EncObservations[enc], *r)
		}
	}
	sort.Sort(ConditionsByOnset(d.Conditions))
	sort.Sort(MedicationsByStart(d.Medications))
	sort.Sort(EncountersByStart(d.Encounters))
	return d
}

type PatientReportDataByFamilyName []*PatientReportData

func (a PatientReportDataByFamilyName) Len() int      { return len(a) }
func (a PatientReportDataByFamilyName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PatientReportDataByFamilyName) Less(i, j int) bool {
	return a[i].Patient.Name[0].Family[0] < a[j].Patient.Name[0].Family[0]
}

type ConditionsByOnset []models.Condition

func (a ConditionsByOnset) Len() int      { return len(a) }
func (a ConditionsByOnset) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ConditionsByOnset) Less(i, j int) bool {
	return a[i].OnsetDateTime.Time.Before(a[j].OnsetDateTime.Time)
}

type MedicationsByStart []models.MedicationStatement

func (a MedicationsByStart) Len() int      { return len(a) }
func (a MedicationsByStart) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a MedicationsByStart) Less(i, j int) bool {
	return a[i].EffectivePeriod.Start.Time.Before(a[j].EffectivePeriod.Start.Time)
}

type EncountersByStart []models.Encounter

func (a EncountersByStart) Len() int      { return len(a) }
func (a EncountersByStart) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a EncountersByStart) Less(i, j int) bool {
	return a[i].Period.Start.Time.Before(a[j].Period.Start.Time)
}
