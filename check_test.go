package resource_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	resource "github.com/taylorsilva/static-list-resource"
)

type CheckTestSuite struct {
	suite.Suite
	*require.Assertions
}

func (c *CheckTestSuite) TestInitialCheck() {
	request := resource.Request{
		Source: resource.Source{
			List: []interface{}{"item1", "item2", "item3", "item4", "item5"},
		},
	}
	check := resource.NewCheck()
	response := check.Run(request)
	c.Equal(request.Source.List, response)
}

func (c *CheckTestSuite) TestReturnNextItem() {
	request := resource.Request{
		Source: resource.Source{
			List: []interface{}{"item1", "item2", "item3", "item4", "item5"},
		},
		Version: "item3",
	}
	check := resource.NewCheck()
	response := check.Run(request)
	c.Equal([]interface{}{"item4"}, response)
}

func (c *CheckTestSuite) TestReturnFirstItemWhenEndIsReached() {
	request := resource.Request{
		Source: resource.Source{
			List: []interface{}{"item1", "item2", "item3", "item4", "item5"},
		},
		Version: "item5",
	}
	check := resource.NewCheck()
	response := check.Run(request)
	c.Equal([]interface{}{"item1"}, response)
}

func TestCheckSuite(t *testing.T) {
	suite.Run(t, &CheckTestSuite{Assertions: require.New(t)})
}
