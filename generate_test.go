package ptgen

import (
	"testing"

	"github.com/intervention-engine/fhir/models"
	. "gopkg.in/check.v1"
)

type GenerateSuite struct {
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&GenerateSuite{})

func (g *GenerateSuite) TestGeneratePatient(c *C) {
	stuff := GeneratePatient()
	pt := stuff[0].(*models.Patient)

	c.Assert(pt.BirthDate, NotNil)
}

func (g *GenerateSuite) TestRandomBirthDate(c *C) {
	bd := RandomBirthDate()
	c.Assert(bd, NotNil)
}
