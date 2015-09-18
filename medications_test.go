package ptgen

import (
	"time"

	"github.com/intervention-engine/fhir/models"
	. "gopkg.in/check.v1"
)

type MedicationSuite struct {
}

var _ = Suite(&MedicationSuite{})

func (m *MedicationSuite) TestGenerateMedication(c *C) {
	mmd := LoadMedications()
	t := time.Now()
	med := GenerateMedication(3, &models.FHIRDateTime{Time: t, Precision: models.Timestamp}, mmd)
	c.Assert(med.MedicationCodeableConcept.Text, Equals, "Lisinopril 5mg Oral Tablet")
	c.Assert(med.EffectivePeriod.Start.Time, Equals, t)
}
