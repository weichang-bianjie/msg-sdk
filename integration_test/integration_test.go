package integration

import (
	"github.com/stretchr/testify/suite"
	msg_sdk "github.com/weichang-bianjie/msg-sdk"
	"testing"
)

type IntegrationTestSuite struct {
	msg_sdk.MsgClient
	suite.Suite
}

type SubTest struct {
	testName string
	testCase func(s IntegrationTestSuite)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.MsgClient = msg_sdk.NewMsgClient()
}
