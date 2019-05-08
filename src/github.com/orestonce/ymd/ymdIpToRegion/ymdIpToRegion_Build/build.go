package ymdIpToRegion_Build

import (
	"github.com/orestonce/ymd/ymdBindata"
)

func YmdBuild() {
	ymdBindata.MustBuildResource(ymdBindata.MustBuildResourceRequest{
		Source:       `src/github.com/orestonce/ymd/ymdIpToRegion/ymdIpToRegion_Build/ip2region.db`,
		Output:       `src/github.com/orestonce/ymd/ymdIpToRegion/zzzig_ip2region.go`,
		UseByteSlice: true,
	})
	// Origin:      `https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.db?raw=true`,
	//	UpdateDate:  2018-05-17`,
}
