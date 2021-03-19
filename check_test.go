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
	tests := map[string]struct {
		Request          resource.CheckRequest
		ExpectedResponse resource.CheckResponse
		ExpctedErr       string
	}{
		"given no version, it should return the first item": {
			Request: resource.CheckRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
			},
			ExpectedResponse: resource.CheckResponse{resource.Version{Item: "item1"}},
		},
		"when given a previouos version it should return nothing": {
			Request: resource.CheckRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
				Version: resource.Version{Item: "item3"},
			},
			ExpectedResponse: resource.CheckResponse{},
		},
		"first item in list should be returned if given version is not found": {
			Request: resource.CheckRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
				Version: resource.Version{Item: "item6"},
			},
			ExpectedResponse: resource.CheckResponse{resource.Version{Item: "item1"}},
		},
		"return error when source list is empty": {
			Request: resource.CheckRequest{
				Source: resource.Source{
					List: []string{},
				},
				Version: resource.Version{Item: "item2"},
			},
			ExpctedErr: "empty list provided in resouce's Source. At least one item required",
		},
	}

	for name, tc := range tests {
		c.Run(name, func() {
			check := resource.NewCheck()
			response, err := check.Run(tc.Request)
			if response != nil {
				if len(tc.ExpectedResponse) == 0 {
					c.Equal(tc.ExpectedResponse, response, name)
				} else {
					c.Equal(tc.ExpectedResponse[0].Item, response[0].Item, name)
					c.NotEqual(time.Time{}, response[0].Date, "time is not nil/default time.Time")
				}
			}
			if tc.ExpctedErr != "" {
				c.EqualError(err, tc.ExpctedErr)
			} else {
				c.NoError(err)
			}
		})
	}
}

func TestCheckSuite(t *testing.T) {
	suite.Run(t, &CheckTestSuite{Assertions: require.New(t)})
}
