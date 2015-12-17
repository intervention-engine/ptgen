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
	start := models.FHIRDateTime{Time: t, Precision: models.Timestamp}
	end := models.FHIRDateTime{Time: t.AddDate(0, 3, 0), Precision: models.Date}
	med := GenerateMedication(3, &start, &end, mmd)
	c.Assert(med.MedicationCodeableConcept.Text, Equals, "Lisinopril 5mg Oral Tablet")
	c.Assert(med.EffectivePeriod.Start.Time, Equals, t)
	c.Assert(med.EffectivePeriod.End.Time, Equals, t.AddDate(0, 3, 0))
	c.Assert(med.Status, Equals, "completed")
}
