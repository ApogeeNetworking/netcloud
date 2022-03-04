package main

import (
	"fmt"
	"os"

	"github.com/ApogeeNetworking/netcloud"
	"github.com/subosito/gotenv"
)

var (
	cpApiID   string
	cpApiKey  string
	ecmApiID  string
	ecmApiKey string
)

func init() {
	gotenv.Load()
	cpApiID = os.Getenv("XCP_APIID")
	cpApiKey = os.Getenv("XCP_APIKEY")
	ecmApiID = os.Getenv("XECM_APIID")
	ecmApiKey = os.Getenv("XECM_APIKEY")
}

func main() {
	cp := netcloud.NewClient(netcloud.AuthParams{
		CpApiID:   cpApiID,
		CpApiKey:  cpApiKey,
		EcmApiID:  ecmApiID,
		EcmApiKey: ecmApiKey,
	})
	rtrs, _ := cp.GetRouter(netcloud.RouterReqParams{})
	for _, rtr := range rtrs {
		fmt.Println(rtr.Name)
	}
}
