package main_test

import (
	"github.com/Microsoft/hcsshim/hcn"
	"github.com/Microsoft/windows-container-networking/test/utilities"
	"testing"
)

func CreateNatTestNetwork() *hcn.HostComputeNetwork {
	ipams := util.GetDefaultIpams()
	return util.CreateTestNetwork("natNet", "Nat", ipams, false)
}

func TestNatCmdAdd(t *testing.T) {
	testNetwork := CreateNatTestNetwork()
	pt := util.MakeTestStruct(t, testNetwork, "nat", false, false, "")
	pt.RunAll(t)
}
