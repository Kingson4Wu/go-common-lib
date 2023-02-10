package network_test

import (
	"fmt"
	"github.com/kingson4wu/go-common-lib/network"
	"testing"
)

func TestGetExtranetIp(t *testing.T) {

	fmt.Println(network.GetExtranetIp())
}
