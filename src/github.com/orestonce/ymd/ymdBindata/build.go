package ymdBindata

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"github.com/orestonce/ymd/ymdError"
	"github.com/orestonce/ymd/ymdFile"
	"fmt"
	"bytes"
	"strings"
	"path"
)

type MustBuildResourceRequest struct {
	Source       string
	Output       string
	UseByteSlice bool
	SkipFileCb   func(filename string) bool
}

func MustBuildResource(req MustBuildResourceRequest) {
	err := os.Chdir(os.Getenv(`GOPATH`))
	ymdError.PanicIfError(err)
	req.Source = filepath.ToSlash(req.Source)
	req.Output = filepath.ToSlash(req.Output)
	if strings.HasSuffix(req.Source, string(`/`)) {
		req.Source = strings.TrimSuffix(req.Source, `/`)
	}
	pkg := path.Base(path.Dir(req.Output))
	wFileList := bytes.NewBuffer(nil)
	w := bytes.NewBuffer(nil)
	w.WriteString(`package `)
	w.WriteString(pkg)
	isFirst := true
	finalType := `string`
	if req.UseByteSlice {
		finalType = `[]byte`
	}
	filepath.Walk(req.Source, func(path string, info os.FileInfo, err error) error {
		var relativePath = filepath.ToSlash(path)
		relativePath = strings.TrimPrefix(relativePath, req.Source)
		if relativePath == `` {
			// req.Source is file
			relativePath = `/` + filepath.Base(req.Source)
		}
		if info == nil || info.IsDir() {
			return nil
		}
		if req.SkipFileCb != nil && req.SkipFileCb(relativePath) {
			return nil
		}
		content, err := ioutil.ReadFile(path)
		ymdError.PanicIfError(err)
		if isFirst {
			fmt.Fprint(w, `

import "time"

func init() {
	var content   `, finalType, "\n")
		}
		isFirst = false

		state, err := os.Stat(path)
		ymdError.PanicIfError(err)
		if !req.UseByteSlice {
			w.WriteString("\t")
			w.WriteString(`content = "`)
			inWriter := &StringWriter{
				Writer: w,
			}
			inWriter.Write(content)
			fmt.Fprint(w, `"`)
		} else {
			w.WriteString("\t")
			w.WriteString(`content = []byte{`)
			inWriter := &ByteWriter{
				Writer: w,
			}
			inWriter.Write(content)
			w.WriteString("}")
		}
		w.WriteString("\n\t")
		fmt.Fprint(w, "_addBinData(`", relativePath, "`, content, time.Unix(", state.ModTime().Unix(), `,0))
`)
		wFileList.WriteString("\t\t\t" + relativePath + "\n")
		return nil
	})

	if !isFirst {
		w.WriteString("}\n")
		ymdFile.MustWriteFile(req.Output, w.Bytes())
	}

	ymdFile.MustWriteFile(strings.TrimSuffix(req.Output, `.go`)+"_getter.go", []byte(fmt.Sprint(`package `, pkg, `

import "time"

type YmdBinData struct {
	DataPath			string
	Content				`, finalType, `
	ModTime      		time.Time
}
var gYmdBinDataAll = map[string] YmdBinData{}

func GetBinData(path string) (data YmdBinData, exists bool) {
	data, exists = gYmdBinDataAll[path]
	return 
}

func _addBinData(path string, content `, finalType, `, modTime time.Time) {
	_, ok := GetBinData(path)
	if ok {
		panic("r5c8hln3 " + path)
	}
	gYmdBinDataAll[path] = YmdBinData {
		DataPath:	path,
		Content:	content,
		ModTime:	modTime,
	}
}

/* filelist:
`+ wFileList.String()+ "*/")),
	)
}
