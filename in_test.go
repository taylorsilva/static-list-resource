package resource_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	resource "github.com/taylorsilva/static-list-resource"
)

type InTestSuite struct {
	suite.Suite
	*require.Assertions
}

func (c *InTestSuite) TestReturnVersion() {
	request := resource.InRequest{
		Source: resource.Source{
			List: []interface{}{"item1", "item2", "item3", "item4", "item5"},
		},
		Version: resource.Version{Item: "item4"},
	}

	in := resource.NewIn()
	response, _ := in.Run(request)
	expected := resource.InResponse{
		Version: resource.Version{Item: "item4"},
	}
	c.Equal(expected, response, "given item4 it should return item4")
}

func (c *InTestSuite) TestVersionNotFound() {
}

func TestInSuite(t *testing.T) {
	suite.Run(t, &InTestSuite{Assertions: require.New(t)})
}
