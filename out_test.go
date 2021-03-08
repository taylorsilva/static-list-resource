package resource_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	resource "github.com/taylorsilva/static-list-resource"
)

type OutTestSuite struct {
	suite.Suite
	*require.Assertions
}

func (o *OutTestSuite) TestOut() {
	tests := []struct {
		Description     string
		Request         resource.OutRequest
		ArtifactPath    string
		ExpectedVersion resource.Version
		PreviousVersion string
		ExpctedErr      string
	}{
		{
			Description: "returns first item if no previous version is found",
			Request: resource.OutRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
				Params: resource.OutParams{
					Previous: "item",
				},
			},
			ArtifactPath:    o.T().TempDir(),
			PreviousVersion: "",
			ExpectedVersion: resource.Version{Item: "item1"},
		},
		{
			Description: "returns next item when previous version provided",
			Request: resource.OutRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
				Params: resource.OutParams{
					Previous: "item",
				},
			},
			ArtifactPath:    o.T().TempDir(),
			PreviousVersion: "item2",
			ExpectedVersion: resource.Version{Item: "item3"},
		},
		{
			Description: "returns first item if end of list is reached",
			Request: resource.OutRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
				Params: resource.OutParams{
					Previous: "item",
				},
			},
			ArtifactPath:    o.T().TempDir(),
			PreviousVersion: "item5",
			ExpectedVersion: resource.Version{Item: "item1"},
		},
		{
			Description: "returns error when source list is empty",
			ExpctedErr:  "empty list provided in resouce's Source. At least one item required",
			Request: resource.OutRequest{
				Source: resource.Source{
					List: []string{},
				},
				Params: resource.OutParams{
					Previous: "item",
				},
			},
			ArtifactPath: o.T().TempDir(),
		},
		{
			Description: "returns error when no file is provided",
			ExpctedErr:  "no file provided with previous version",
			Request: resource.OutRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
			},
		},
	}

	for _, t := range tests {
		if t.Request.Params.Previous != "" {
			pf := filepath.Join(t.ArtifactPath, t.Request.Params.Previous)
			os.WriteFile(pf, []byte(t.PreviousVersion), 0666)
		}

		out := resource.NewOut()
		response, err := out.Run(t.Request, t.ArtifactPath)

		if t.ExpectedVersion.Item != "" {
			o.Equal(t.ExpectedVersion.Item, response.Version.Item, t.Description)
			o.NotEqual(time.Time{}, response.Version.Date, "time is not nil/default time.Time")
		}

		if t.ExpctedErr != "" {
			o.EqualError(err, t.ExpctedErr, t.Description)
		} else {
			o.NoError(err)
		}
	}
}

func TestOutSuite(t *testing.T) {
	suite.Run(t, &OutTestSuite{Assertions: require.New(t)})
}
