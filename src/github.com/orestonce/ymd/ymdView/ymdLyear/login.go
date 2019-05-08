package ymdLyear

import (
	"io"
	"fmt"
	"github.com/orestonce/ymd/ymdView/ymdXss"
	"github.com/orestonce/ymd/ymdGin"
)

type TplLoginPageRequest struct {
	ActionUrl string
	ErrMsg    string
}

type LoginActionRequest struct {
	Username string
	Password string
}

func LoginActionGetRequest(ctx *ymdGin.Context) (req LoginActionRequest) {
	req.Username = ctx.InStr(`username`)
	req.Password = ctx.InStr(`password`)
	return
}

func TplLoginPage(buf io.Writer, req TplLoginPageRequest) {
	writeHeader(buf, `登陆`)
	fmt.Fprintln(buf, `
<style>
.lyear-wrapper {
    position: relative;
}
.lyear-login {
    display: flex !important;
    min-height: 100vh;
    align-items: center !important;
    justify-content: center !important;
}
.login-center {
    background: #fff;
    min-width: 38.25rem;
    padding: 2.14286em 3.57143em;
    border-radius: 5px;
    margin: 2.85714em 0;
}
.login-header {
    margin-bottom: 1.5rem !important;
}
.login-center .has-feedback.feedback-left .form-control {
    padding-left: 38px;
    padding-right: 12px;
}
.login-center .has-feedback.feedback-left .form-control-feedback {
    left: 0;
    right: auto;
    width: 38px;
    height: 38px;
    line-height: 38px;
    z-index: 4;
    color: #dcdcdc;
}
.login-center .has-feedback.feedback-left.row .form-control-feedback {
    left: 15px;
}
</style>
<div class="row lyear-wrapper">
  <div class="lyear-login">
    <div class="login-center">
      <div class="login-header text-center">
        <img alt="light year admin" src="/lyear/images/logo-sidebar.png">
      </div>`)
	if req.ErrMsg != `` {
		fmt.Fprintln(buf, `<div class="alert alert-danger" role="alert">`, ymdXss.Urlv(req.ErrMsg), `</div>`)
	}
	fmt.Fprintln(buf, `<form action="`+req.ActionUrl)
	fmt.Fprintln(buf, `" method="post">
        <div class="form-group has-feedback feedback-left">
          <input type="text" placeholder="用户名" class="form-control" name="username"/>
        </div>
        <div class="form-group has-feedback feedback-left">
          <input type="password" placeholder="密码" class="form-control" name="password" />
        </div>
        <div class="form-group">
          <button class="btn btn-block btn-primary" type="submit">立即登录</button>
        </div>
      </form>
      <hr>`)
	fmt.Fprintln(buf, `
    </div>
  </div>
</div>
<script type="text/javascript" src="/lyear/js/jquery.min.js"></script>
<script type="text/javascript" src="/lyear/js/bootstrap.min.js"></script>
</body>
</html>`)
}
