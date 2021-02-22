package resource_test

import (
	"os"
	"path/filepath"
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
			List: []string{"item1", "item2", "item3", "item4", "item5"},
		},
		Version: resource.Version{Item: "item4"},
	}

	tmpdir := c.T().TempDir()

	in := resource.NewIn()
	response, err := in.Run(request, tmpdir)
	expected := resource.InResponse{
		Version: resource.Version{Item: "item4"},
	}

	c.NoError(err)
	c.Equal(expected, response, "given item4 it should return item4")

	file := filepath.Join(tmpdir, "item")
	c.FileExists(file)
	contents, _ := os.ReadFile(file)
	c.Equal([]byte(`item4`), contents)
}

func (c *InTestSuite) TestVersionNotFound() {
	request := resource.InRequest{
		Source: resource.Source{
			List: []string{"item1", "item2", "item3", "item4", "item5"},
		},
		Version: resource.Version{Item: "other"},
	}

	tmpdir := c.T().TempDir()

	in := resource.NewIn()
	response, err := in.Run(request, tmpdir)
	c.Equal(response, resource.InResponse{})
	c.EqualError(err, "selected item not found in source.list")
}

func TestInSuite(t *testing.T) {
	suite.Run(t, &InTestSuite{Assertions: require.New(t)})
}
