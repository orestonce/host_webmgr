package ymdLyear

import (
	"github.com/orestonce/ymd/ymdView"
	"io"
	"fmt"
	"github.com/orestonce/ymd/ymdView/ymdXss"
)

func WriteHtmlPage(w io.Writer, warp WarpHtmlPage, name string, main ymdView.HtmlRenderer) {
	left := warp.LeftSide
	var title string
	for idx, nodeL1 := range left {
		if nodeL1.GetShortName() == name {
			title = nodeL1.ShowName
			nodeL1.IsActive = true
		}
		for idx, nodeL2 := range nodeL1.ChildNode {
			if nodeL2.GetShortName() == name {
				title = nodeL2.ShowName
				nodeL2.IsActive = true
				nodeL1.IsActive = true
			}
			nodeL1.ChildNode[idx] = nodeL2
		}
		left[idx] = nodeL1
	}
	writeHeader(w, title)
	w.Write([]byte(`<!--左侧导航-->
    <aside class="lyear-layout-sidebar">
      
      <!-- logo -->
      <div id="logo" class="sidebar-header">
        <a ><img src="/lyear/images/logo-sidebar.png" title="LightYear" alt="LightYear"/></a>
      </div>
      <div class="lyear-layout-sidebar-scroll"> 
        
        <nav class="sidebar-main">
          <ul class="nav nav-drawer">`))
	for _, oneL1 := range left {
		w.Write([]byte(`<li class="nav-item `))
		if len(oneL1.ChildNode) > 0 {
			w.Write([]byte(`nav-item-has-subnav `))
		}
		if oneL1.IsActive {
			w.Write([]byte(`active `))
			if len(oneL1.ChildNode) > 0 {
				w.Write([]byte(`open `))
			}
		}
		href := oneL1.UrlPath
		if len(oneL1.ChildNode) > 0 {
			href = `javascript:void(0)`
		}
		fmt.Fprint(w, `"> <a href="`, href, `">`, ymdXss.Urlv(oneL1.ShowName)+`</a>`)
		if len(oneL1.ChildNode) > 0 {
			fmt.Fprint(w, `<ul class="nav nav-subnav">
              `)
			//   <li class="active"> <a href="lyear_pages_rabc.html">设置权限</a> </li>
			for _, oneL2 := range oneL1.ChildNode {
				if oneL2.IsActive {
					fmt.Fprint(w, `<li class="active"> `)
				} else {
					fmt.Fprint(w, `<li> `)
				}
				fmt.Fprint(w, `<a href="`, oneL2.UrlPath, `">`, ymdXss.Urlv(oneL2.ShowName), `</a></li>`)
			}
			fmt.Fprint(w, `</ul>`)
		}
		fmt.Fprint(w, `</li>`)
	}
	w.Write([]byte(`
          </ul>
        </nav>
      </div>
    </aside>
    <!--End 左侧导航-->
    
    <!--头部信息-->
    <header class="lyear-layout-header">
      
      <nav class="navbar navbar-default">
        <div class="topbar">
          
          <div class="topbar-left">
            <div class="lyear-aside-toggler">
              <span class="lyear-toggler-bar"></span>
              <span class="lyear-toggler-bar"></span>
              <span class="lyear-toggler-bar"></span>
            </div>
            <span class="navbar-page-title"> `))
	w.Write([]byte(ymdXss.Urlv(title)))
	fmt.Fprintln(w, ` </span>
          </div>
			<ul class="topbar-right">
            <li class="dropdown dropdown-profile">
              <a href="javascript:void(0)" data-toggle="dropdown">
                <img class="img-avatar img-avatar-48 m-r-10" src="/lyear/images/users/avatar.jpg" alt="`,
		ymdXss.Urlv(warp.Username),
		`" />
                <span>`,
		ymdXss.Urlv(warp.Username),
		`<span class="caret"></span></span>
              </a>
              <ul class="dropdown-menu dropdown-menu-right">
                <li> <a href="`,
		ymdXss.Urlv(warp.LogoutUrl),
		`"><i class="mdi mdi-logout-variant"></i> 退出登录</a> </li>
              </ul>
            </li>
      		</ul>
        </div>
      </nav>
    </header>
    <!--End 头部信息-->
	<!--页面主要内容-->
    <main class="lyear-layout-content">
      <div class="container-fluid">`)
	main.HtmlRender(w)
	fmt.Fprint(w, `</div>
</main>`)
	writeFooter(w)
}
