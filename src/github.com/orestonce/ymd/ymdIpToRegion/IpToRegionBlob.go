package ymdIpToRegion

import (
	"strings"
	"errors"
	"encoding/binary"
	"sync"
	"fmt"
	"strconv"
)

const (
	cINDEX_BLOCK_LENGTH = 12
)

var gIp2Region struct {
	firstIndexPtr uint32 // super block index info
	lastIndexPtr  uint32
	totalBlocks   uint32
	dbBinStr      []byte // the original db binary string
	initOnce      sync.Once
}

type IpInfo struct {
	Country  string
	Region   string
	Province string
	City     string
	ISP      string
}

func (ip IpInfo) String() string {
	return ip.Country + "|" + ip.Region + "|" + ip.Province + "|" + ip.City + "|" + ip.ISP
}

func getIpInfo(line []byte) IpInfo {
	lineSlice := strings.Split(string(line), "|")
	ipInfo := IpInfo{}
	length := len(lineSlice)
	if length < 5 {
		panic(`kxrbiwiv ` + string(line))
	}
	getInfo := func(s string) string {
		if s == `0` {
			return ``
		}
		return s
	}
	ipInfo.Country = getInfo(lineSlice[0])
	ipInfo.Region = getInfo(lineSlice[1])
	ipInfo.Province = getInfo(lineSlice[2])
	ipInfo.City = getInfo(lineSlice[3])
	ipInfo.ISP = getInfo(lineSlice[4])
	return ipInfo
}

func GetIpInfo(ipStr string) (ipInfo IpInfo, err error) {
	ipInfo = IpInfo{}
	getUint32 := func(b []byte, offset uint32) uint32 {
		return binary.LittleEndian.Uint32(b[offset:])
	}
	gIp2Region.initOnce.Do(func() {
		data, _ := GetBinData(`/ip2region.db`)
		gIp2Region.dbBinStr = data.Content
		gIp2Region.firstIndexPtr = getUint32(gIp2Region.dbBinStr, 0)
		gIp2Region.lastIndexPtr = getUint32(gIp2Region.dbBinStr, 4)
		gIp2Region.totalBlocks = (gIp2Region.lastIndexPtr-gIp2Region.firstIndexPtr)/cINDEX_BLOCK_LENGTH + 1
	})
	ip, err := ip2uint32(ipStr)
	if err != nil {
		return ipInfo, err
	}
	h := gIp2Region.totalBlocks
	var dataPtr, l uint32;
	for l <= h {
		m := (l + h) / 2
		p := gIp2Region.firstIndexPtr + m*cINDEX_BLOCK_LENGTH
		sip := getUint32(gIp2Region.dbBinStr, p)
		eip := getUint32(gIp2Region.dbBinStr, p+4)
		if ip < sip {
			h = m - 1
		} else if ip <= eip {
			dataPtr = getUint32(gIp2Region.dbBinStr, p+8)
			break
		} else {
			l = m + 1
		}
	}
	return getIpInfoFromDataPtr(dataPtr)
}

func getIpInfoFromDataPtr(dataPtr uint32) (ipInfo IpInfo, err error) {
	if dataPtr <= 0 {
		return ipInfo, fmt.Errorf(`l5swcknq %v`, dataPtr)
	}
	dataLen := ((dataPtr >> 24) & 0xFF)
	dataPtr = (dataPtr & 0x00FFFFFF);
	ipInfo = getIpInfo(gIp2Region.dbBinStr[(dataPtr)+4 : dataPtr+dataLen])
	return ipInfo, nil
}

func ip2uint32(IpStr string) (uint32, error) {
	bits := strings.Split(IpStr, ".")
	if len(bits) != 4 {
		return 0, errors.New("3ujds20o ip format error")
	}
	var sum uint32
	for i, n := range bits {
		bit, err := strconv.Atoi(n)
		if err != nil {
			return 0, errors.New(`jcynfhuj ` + err.Error())
		}
		if bit < 0 || bit > 255 {
			return 0, errors.New(`ikw84xa7 ` + strconv.Itoa(bit))
		}
		sum += uint32(bit) << uint32(24-8*i)
	}
	return sum, nil
}
