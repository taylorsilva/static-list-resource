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
	request := resource.CheckRequest{
		Source: resource.Source{
			List: []interface{}{"item1", "item2", "item3", "item4", "item5"},
		},
	}
	check := resource.NewCheck()
	response, _ := check.Run(request)
	expected := resource.CheckResponse{
		resource.Version{Item: "item1"},
		resource.Version{Item: "item2"},
		resource.Version{Item: "item3"},
		resource.Version{Item: "item4"},
		resource.Version{Item: "item5"},
	}
	c.Equal(expected, response, "given no version, it should return the entire list")
}

func (c *CheckTestSuite) TestReturnNextItem() {
	request := resource.CheckRequest{
		Source: resource.Source{
			List: []interface{}{"item1", "item2", "item3", "item4", "item5"},
		},
		Version: resource.Version{Item: "item3"},
	}
	check := resource.NewCheck()
	response, _ := check.Run(request)
	expected := resource.CheckResponse{
		resource.Version{Item: "item4"},
	}
	c.Equal(expected, response, "given item3 it should return item4")
}

func (c *CheckTestSuite) TestReturnFirstItemWhenEndIsReached() {
	request := resource.CheckRequest{
		Source: resource.Source{
			List: []interface{}{"item1", "item2", "item3", "item4", "item5"},
		},
		Version: resource.Version{Item: "item5"},
	}
	check := resource.NewCheck()
	response, _ := check.Run(request)
	expected := resource.CheckResponse{
		resource.Version{Item: "item1"},
	}
	c.Equal(expected, response, "given the last item in the list, it should return the first item")
}

func (c *CheckTestSuite) TestLastVersionRemovedFromList() {
	request := resource.CheckRequest{
		Source: resource.Source{
			List: []interface{}{"item1", "item2", "item3", "item4", "item5"},
		},
		Version: resource.Version{Item: "item6"},
	}
	check := resource.NewCheck()
	response, _ := check.Run(request)
	expected := resource.CheckResponse{
		resource.Version{Item: "item1"},
	}
	c.Equal(expected, response, "first item in list should be returned if given version not found")
}

func (c *CheckTestSuite) TestErrorIfListIsEmpty() {
	request := resource.CheckRequest{
		Source: resource.Source{
			List: []interface{}{},
		},
		Version: resource.Version{Item: "item2"},
	}
	check := resource.NewCheck()
	_, err := check.Run(request)
	c.EqualError(err, "empty list provided in resouce's Source. At least one item required")
}

func TestCheckSuite(t *testing.T) {
	suite.Run(t, &CheckTestSuite{Assertions: require.New(t)})
}
