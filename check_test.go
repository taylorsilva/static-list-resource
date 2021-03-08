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

func (c *CheckTestSuite) TestCheck() {
	tests := []struct {
		Description     string
		Request         resource.CheckRequest
		ExpectedVersion resource.Version
		ExpctedErr      string
	}{
		{
			Description: "given no version, it should return the first item",
			Request: resource.CheckRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
			},
			ExpectedVersion: resource.Version{Item: "item1"},
		},
		{
			Description: "should return the next item: given item3 it should return item4",
			Request: resource.CheckRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
				Version: resource.Version{Item: "item3"},
			},
			ExpectedVersion: resource.Version{Item: "item4"},
		},
		{
			Description: "given the last item in the list, it should return the first item",
			Request: resource.CheckRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
				Version: resource.Version{Item: "item5"},
			},
			ExpectedVersion: resource.Version{Item: "item1"},
		},
		{
			Description: "first item in list should be returned if given version is not found",
			Request: resource.CheckRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
				Version: resource.Version{Item: "item6"},
			},
			ExpectedVersion: resource.Version{Item: "item1"},
		},
		{
			Description: "first item in list should be returned if given version is not found",
			Request: resource.CheckRequest{
				Source: resource.Source{
					List: []string{},
				},
				Version: resource.Version{Item: "item2"},
			},
			ExpctedErr: "empty list provided in resouce's Source. At least one item required",
		},
	}

	for _, test := range tests {
		check := resource.NewCheck()
		response, err := check.Run(test.Request)
		if response != nil {
			c.Equal(test.ExpectedVersion.Item, response[0].Item, test.Description)
			c.NotEqual(time.Time{}, response[0].Date, "time is not nil/default time.Time")
		}
		if test.ExpctedErr != "" {
			c.EqualError(err, test.ExpctedErr)
		} else {
			c.NoError(err)
		}
	}
}

func TestCheckSuite(t *testing.T) {
	suite.Run(t, &CheckTestSuite{Assertions: require.New(t)})
}
