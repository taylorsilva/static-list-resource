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
	tests := map[string]struct {
		Description     string
		Request         resource.OutRequest
		ArtifactPath    string
		ExpectedVersion resource.Version
		PreviousVersion string
		ExpctedErr      string
	}{
		"returns first item if no previous version is found": {
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
		"returns next item when previous version provided": {
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
		"returns first item if end of list is reached": {
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
		"returns error when source list is empty": {
			ExpctedErr: "empty list provided in resouce's Source. At least one item required",
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
		"returns error when no file is provided": {
			ExpctedErr: "no file provided with previous version",
			Request: resource.OutRequest{
				Source: resource.Source{
					List: []string{"item1", "item2", "item3", "item4", "item5"},
				},
			},
		},
	}

	for name, tc := range tests {
		o.Run(name, func() {
			if tc.Request.Params.Previous != "" {
				pf := filepath.Join(tc.ArtifactPath, tc.Request.Params.Previous)
				os.WriteFile(pf, []byte(tc.PreviousVersion), 0666)
			}

			out := resource.NewOut()
			response, err := out.Run(tc.Request, tc.ArtifactPath)

			if tc.ExpectedVersion.Item != "" {
				o.Equal(tc.ExpectedVersion.Item, response.Version.Item)
				o.NotEqual(time.Time{}, response.Version.Date, "time should not be default time.Time")
			}

			if tc.ExpctedErr != "" {
				o.EqualError(err, tc.ExpctedErr, tc.Description)
			} else {
				o.NoError(err)
			}
		})
	}
}

func TestOutSuite(t *testing.T) {
	suite.Run(t, &OutTestSuite{Assertions: require.New(t)})
}
