package e2e

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	errorz "github.com/dungnh3/mfv-codingchallenge/internal/errors"

	"gorm.io/gorm"

	"github.com/dungnh3/mfv-codingchallenge/config"
	mocks "github.com/dungnh3/mfv-codingchallenge/internal/mocks/repositories"
	"github.com/dungnh3/mfv-codingchallenge/internal/models"
	"github.com/dungnh3/mfv-codingchallenge/internal/models/store"
	"github.com/dungnh3/mfv-codingchallenge/internal/services"
	l "github.com/dungnh3/mfv-codingchallenge/pkg/log"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const host = "http://localhost:9090"

func GenURI(path string) string {
	return host + path
}

type MockTestSuite struct {
	suite.Suite
	svr *services.Server
	r   *mocks.Repository
}

func TestMock(t *testing.T) {
	suite.Run(t, &MockTestSuite{})
}

func (s *MockTestSuite) SetupSuite() {
	var err error
	conf := config.Load()
	logger := l.New()

	repo := mocks.NewRepository(s.T())
	s.r = repo

	s.svr = services.New(conf, repo)
	go func() {
		if err = s.svr.Run(); err != nil {
			logger.Error("running application error", l.Error(err))
		}
	}()
}

func (s *MockTestSuite) TearDownSuite() {
	var err error
	err = s.svr.Close(context.Background())
	s.Require().NoError(err)
}

func (s *MockTestSuite) Test_GetUser_Success() {
	s.r.On("GetUser", mock.Anything, mock.Anything).
		Return(&models.User{
			User: &store.User{
				ID:     1,
				Name:   "user01",
				Status: "active",
			},
			AccountIDs: []int64{1, 2, 3},
		}, nil).Once()

	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, GenURI("/users/1"), nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var user models.User
	err = json.Unmarshal(body, &user)
	s.Require().NoError(err)

	s.EqualValues(1, user.ID)
	s.EqualValues("user01", user.Name)
	s.EqualValues([]int64{1, 2, 3}, user.AccountIDs)
}

func (s *MockTestSuite) Test_GetUser_Failed() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, GenURI("/users/abc"), nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.NotEqualValues(http.StatusOK, resp.StatusCode)
}

func (s *MockTestSuite) Test_GetAccount_Success() {
	s.r.On("GetAccount", mock.Anything, mock.Anything).
		Return(&models.UserAccount{
			UserAccount: &store.UserAccount{
				ID:      1,
				UserID:  1,
				Name:    "account01",
				Status:  "active",
				Balance: 2000,
			},
		}, nil).Once()

	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, GenURI("/accounts/1"), nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var acc models.UserAccount
	err = json.Unmarshal(body, &acc)
	s.Require().NoError(err)

	s.EqualValues(1, acc.ID)
	s.EqualValues(1, acc.UserID)
	s.EqualValues("account01", acc.Name)
	s.EqualValues(2000, acc.Balance)
}

func (s *MockTestSuite) Test_GetAccount_Failed() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, GenURI("/accounts/a"), nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.NotEqualValues(http.StatusOK, resp.StatusCode)
}

func (s *MockTestSuite) Test_ListUserAccount_Success() {
	arr := []*models.UserAccount{
		{
			UserAccount: &store.UserAccount{
				ID:      1,
				UserID:  1,
				Name:    "account01",
				Balance: 100,
			},
		}, {
			UserAccount: &store.UserAccount{
				ID:      2,
				UserID:  1,
				Name:    "account01",
				Balance: 200,
			},
		},
	}
	s.r.On("ListAccounts", mock.Anything, mock.Anything).Return(arr, nil).Once()

	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, GenURI("/users/1/accounts"), nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var arrResponse []*models.UserAccount
	err = json.Unmarshal(body, &arrResponse)
	s.Require().NoError(err)

	s.EqualValues(len(arrResponse), 2)
}

func (s *MockTestSuite) Test_ListUserAccount_Failed() {
	s.r.On("ListAccounts", mock.Anything, mock.Anything).
		Return(nil, gorm.ErrRecordNotFound).Once()

	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, GenURI("/users/1/accounts"), nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var errResponse errorz.ErrResponse
	err = json.Unmarshal(body, &errResponse)
	s.Require().NoError(err)

	s.EqualValues(10001, errResponse.ErrorCode)
	s.EqualValues("record not found", errResponse.Description)
}
