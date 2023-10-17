package e2e

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/dungnh3/mfv-codingchallenge/config"
	"github.com/dungnh3/mfv-codingchallenge/internal/repositories"
	"github.com/dungnh3/mfv-codingchallenge/internal/repositories/mysql"
	"github.com/dungnh3/mfv-codingchallenge/internal/services"
	pkgdb "github.com/dungnh3/mfv-codingchallenge/pkg/db"
	l "github.com/dungnh3/mfv-codingchallenge/pkg/log"
	"github.com/stretchr/testify/suite"
)

type E2ETestSuite struct {
	suite.Suite
	svr *services.Server
	r   *repositories.Repository
}

func TestE2E(t *testing.T) {
	suite.Run(t, &E2ETestSuite{})
}

func (s *E2ETestSuite) SetupSuite() {
	var err error
	conf := config.Load()
	logger := l.New()

	conf.MySQL.Database = "mfv_test"
	db := pkgdb.ConnectMySQL(conf.MySQL)
	repo := mysql.New(db, conf)

	s.svr = services.New(conf, repo)
	go func() {
		if err = s.svr.Run(); err != nil {
			logger.Error("running application error", l.Error(err))
		}
	}()
}

func (s *E2ETestSuite) TearDownSuite() {
	var err error
	err = s.svr.Close(context.Background())
	s.Require().NoError(err)
}

//go:embed register_user_payload.json
var registerUserPayload []byte

func (s *E2ETestSuite) Test_01_RegisterUser_Success() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/users/register", bytes.NewBuffer(registerUserPayload))
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var user services.RegisterUserResponse
	err = json.Unmarshal(body, &user)
	s.Require().NoError(err)
	s.NotEqualValues(int64(0), user.ID)
	s.EqualValues("user01", user.Name)
}

func (s *E2ETestSuite) Test_01_RegisterUser_Failed() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/users/register", nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	s.EqualValues(http.StatusBadRequest, resp.StatusCode)
}

//go:embed create_account_payload.json
var createAccountPayload []byte

func (s *E2ETestSuite) Test_02_CreateAccount_Success() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/users/1/accounts", bytes.NewBuffer(createAccountPayload))
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var account services.CreateAccountResponse
	err = json.Unmarshal(body, &account)
	s.Require().NoError(err)
	s.NotEqualValues(int64(0), account.ID)
	s.EqualValues("account01", account.Name)
	s.EqualValues(int64(1), account.UserID)
}

func (s *E2ETestSuite) Test_02_CreateAccount_Failed() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/users/1/accounts", nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	s.EqualValues(http.StatusBadRequest, resp.StatusCode)
}

func (s *E2ETestSuite) Test_03_GetUser_Success() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/users/1", nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var user services.GetUserResponse
	err = json.Unmarshal(body, &user)
	s.Require().NoError(err)

	s.EqualValues(1, user.ID)
	s.EqualValues("user01", user.Name)
	s.EqualValues([]int64{1}, user.AccountIDs)
}

func (s *E2ETestSuite) Test_03_GetUser_Failed() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/users/abc", nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	s.EqualValues(http.StatusBadRequest, resp.StatusCode)
}

func (s *E2ETestSuite) Test_04_ListUserAccount_Success() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/users/1/accounts", nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var arr []*services.GetAccountResponse
	err = json.Unmarshal(body, &arr)
	s.Require().NoError(err)

	s.NotEqualValues(0, len(arr))

	firstElement := arr[0]
	s.EqualValues(int64(1), firstElement.UserID)
	s.EqualValues(int64(1), firstElement.ID)
	s.EqualValues("account01", firstElement.Name)
}

func (s *E2ETestSuite) Test_04_ListUserAccount_Failed() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/users/abc/accounts", nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	s.EqualValues(http.StatusBadRequest, resp.StatusCode)
}

func (s *E2ETestSuite) Test_05_GetAccount_Success() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/accounts/1", nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var acc services.GetAccountResponse
	err = json.Unmarshal(body, &acc)
	s.Require().NoError(err)

	s.EqualValues(int64(1), acc.UserID)
	s.EqualValues(int64(1), acc.ID)
	s.EqualValues("account01", acc.Name)
}

func (s *E2ETestSuite) Test_05_GetAccount_Failed() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/accounts/abc", nil)
	s.Require().NoError(err)
	resp, err := client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()
	s.EqualValues(http.StatusBadRequest, resp.StatusCode)
}
