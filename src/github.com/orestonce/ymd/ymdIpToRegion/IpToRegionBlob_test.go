package ymdIpToRegion

import (
	"testing"
	"fmt"
	"github.com/orestonce/ymd/ymdRand"
	"github.com/orestonce/ymd/ymdError"
	"github.com/orestonce/ymd/ymdAssert"
	"strings"
)

// memory: 5MB
func TestIp2Region_MemorySearch(t *testing.T) {
	info, err := GetIpInfo("23.89.124.35")
	fmt.Println(info, err)
}

func TestGetIpInfoMemorySearch(t *testing.T) {
	_, err := GetIpInfo(`12`)
	ymdAssert.True(strings.Contains(err.Error(), `3ujds20o`))
	_, err = GetIpInfo(`q.w.q.w`)
	ymdAssert.True(strings.Contains(err.Error(), `jcynfhuj`))
	_, err = GetIpInfo(`1234.1.2.2`)
	ymdAssert.True(strings.Contains(err.Error(), `ikw84xa7`))
}

//2000000	       763 ns/op
func BenchmarkIp2Region_MemorySearch(b *testing.B) {
	ips := make([]string, 0, b.N)
	b.StopTimer()
	r := ymdRand.NewMathRandom(123)
	for i := 0; i < b.N; i++ {
		ips = append(ips, fmt.Sprintf("%v.%v.%v.%v", r.Intn(256), r.Intn(256), r.Intn(256), r.Intn(256)))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := GetIpInfo(ips[i])
		ymdError.PanicIfError(err)
	}
}

func TestIpInfo_String(t *testing.T) {
	err := ymdError.PanicToError(func() {
		getIpInfo([]byte(`123`))
	})
	ymdAssert.True(strings.Contains(err.Error(), `kxrbiwiv`))
}

func TestGetIpInfoMemorySearch2(t *testing.T) {
	_, err := GetIpInfo(`8.8.8.8`)
	ymdAssert.True(err == nil, err)

	_, err = getIpInfoFromDataPtr(0)
	ymdAssert.True(strings.Contains(err.Error(), `l5swcknq`))
}
