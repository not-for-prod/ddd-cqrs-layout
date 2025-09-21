package addreview

import (
	"context"
	"testing"

	"yelp/internal/domain/entity/review"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	cmd *Handler
}

func (suite *TestSuite) SetupSuite() {
	repo := &reviewRepositoryMock{
		AddFunc: func(ctx context.Context, reviewMoqParam review.Review) error {
			return nil
		},
	}
	suite.cmd = &Handler{
		reviewRepository: repo,
	}
}

func (suite *TestSuite) TearDownSuite() {}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
