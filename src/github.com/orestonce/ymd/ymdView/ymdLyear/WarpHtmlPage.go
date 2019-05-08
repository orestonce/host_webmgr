package ymdLyear

import "github.com/orestonce/ymd/ymdStrings"

type WarpHtmlPage struct {
	Username     string
	LogoutUrl    string
	LeftSide     []SideNode
}

type SideNode struct {
	UrlPath   string
	ShowName  string
	IsActive  bool
	ChildNode []SideNode
}

func (obj SideNode) GetShortName() (short string) {
	return ymdStrings.StringAfterLastSubString(obj.UrlPath, `/`)
}
