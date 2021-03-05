package resource_test

import (
	"testing"
	"time"

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
			List: []string{"item1", "item2", "item3", "item4", "item5"},
		},
	}
	check := resource.NewCheck()
	response, _ := check.Run(request)
	expected := resource.CheckResponse{
		resource.Version{Item: "item1"},
	}
	c.Equal(expected[0].Item, response[0].Item, "given no version, it should return the first item")
	c.NotEqual(time.Time{}, response[0].Date, "time is not nil/default time.Time")
}

func (c *CheckTestSuite) TestReturnNextItem() {
	request := resource.CheckRequest{
		Source: resource.Source{
			List: []string{"item1", "item2", "item3", "item4", "item5"},
		},
		Version: resource.Version{Item: "item3"},
	}
	check := resource.NewCheck()
	response, _ := check.Run(request)
	expected := resource.CheckResponse{
		resource.Version{Item: "item4"},
	}
	c.Equal(expected[0].Item, response[0].Item, "given item3 it should return item4")
	c.NotEqual(time.Time{}, response[0].Date, "time is not nil/default time.Time")
}

func (c *CheckTestSuite) TestReturnFirstItemWhenEndIsReached() {
	request := resource.CheckRequest{
		Source: resource.Source{
			List: []string{"item1", "item2", "item3", "item4", "item5"},
		},
		Version: resource.Version{Item: "item5"},
	}
	check := resource.NewCheck()
	response, _ := check.Run(request)
	expected := resource.CheckResponse{
		resource.Version{Item: "item1"},
	}
	c.Equal(expected[0].Item, response[0].Item, "given the last item in the list, it should return the first item")
	c.NotEqual(time.Time{}, response[0].Date, "time is not nil/default time.Time")
}

func (c *CheckTestSuite) TestLastVersionRemovedFromList() {
	request := resource.CheckRequest{
		Source: resource.Source{
			List: []string{"item1", "item2", "item3", "item4", "item5"},
		},
		Version: resource.Version{Item: "item6"},
	}
	check := resource.NewCheck()
	response, _ := check.Run(request)
	expected := resource.CheckResponse{
		resource.Version{Item: "item1"},
	}
	c.Equal(expected[0].Item, response[0].Item, "first item in list should be returned if given version not found")
	c.NotEqual(time.Time{}, response[0].Date, "time is not nil/default time.Time")
}

func (c *CheckTestSuite) TestErrorIfListIsEmpty() {
	request := resource.CheckRequest{
		Source: resource.Source{
			List: []string{},
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
