package ymdGin

import (
	"reflect"
	"log"
	"fmt"
	"strings"
)

func RegisterObjToEngine(engine *Engine, prefix string, obj interface{}) {
	objType := reflect.TypeOf(obj)
	var ctxTemp *Context
	for i := 0; i < objType.NumMethod(); i ++ {
		method := objType.Method(i)
		handlePath := prefix + method.Name
		if method.Type.NumIn() != 2 || method.Type.In(1) != reflect.TypeOf(ctxTemp) {
			log.Println("Skip ", handlePath)
			continue
		}
		engine.Any(handlePath, func(ctx *Context) {
			method.Func.Call([]reflect.Value{
				reflect.ValueOf(obj),
				reflect.ValueOf(ctx),
			})
		})
		log.Println("Register ", handlePath)
	}
}

func (ctx *Context) InStr(name string) (value string) {
	value = ctx.Query(name)
	if value != `` {
		return value
	}
	value = ctx.PostForm(name)
	return
}

func (ctx *Context) InBool(name string) (value bool) {
	if strings.ToLower(ctx.InStr(name)) == `true` {
		return true
	}
	return false
}

func (ctx *Context) RedirectTemp(location string) {
	fmt.Fprint(ctx.Writer, `<html><script language="javascript"type="text/javascript"> 
　　window.location.href="`, location, `"; 
</script> </html>`)
}
