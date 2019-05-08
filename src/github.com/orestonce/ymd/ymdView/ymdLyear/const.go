package ymdLyear

import (
	"github.com/orestonce/ymd/ymdView/ymdXss"
	"io"
)

const faviconAndCss = `<link rel="icon" href="/lyear/favicon.ico" type="image/ico">
<link href="/lyear/css/bootstrap.min.css" rel="stylesheet">
<link href="/lyear/css/materialdesignicons.min.css" rel="stylesheet">
<link href="/lyear/css/style.min.css" rel="stylesheet">
`

func writeHeader(buf io.Writer, title string) {
	buf.Write([]byte(`<!DOCTYPE html>
<html lang="zh">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
<title>`))
	buf.Write([]byte(ymdXss.Urlv(title)))
	buf.Write([]byte(`</title>`))
	buf.Write([]byte(faviconAndCss))
	buf.Write([]byte(`
</head>
<body>
`))
}

func writeFooter(buf io.Writer) {
	buf.Write([]byte(`
<script type="text/javascript" src="/lyear/js/jquery.min.js"></script>
<script type="text/javascript" src="/lyear/js/bootstrap.min.js"></script>
<script type="text/javascript" src="/lyear/js/perfect-scrollbar.min.js"></script>
<!--消息提示-->
<script src="/lyear/js/bootstrap-notify.min.js"></script>
<script type="text/javascript" src="/lyear/js/lightyear.js"></script>
<script type="text/javascript" src="/lyear/js/main.min.js"></script>
<script type="text/javascript" src="/lyear/js/ymd.js"></script>
</body>
</html>`))
}
