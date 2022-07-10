package render

import (
	"bytes"
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3/helpers"
	"path"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

var engine *jet.Set

var initOnce *sync.Once

var WebHost string

func init() {
	WebHost = ""
}

type JetEngine struct {
	Set     *jet.Set
	Host    string
	TplRoot string
}

func New(host string, theme string) *JetEngine {
	tplRoot := path.Join(boot.HomeDir(), "themes", theme)

	s := jet.NewHTMLSet(tplRoot)

	addCustomFunc(s)

	e := &JetEngine{
		Set:     s,
		Host:    host,
		TplRoot: tplRoot,
	}

	boot.GinCtx().Engine.Static("themes", tplRoot)

	return e
}
func ArgToString(value reflect.Value) string {
	buffer := bytes.NewBuffer(nil)
	_, _ = fmt.Fprint(buffer, value)
	return string(buffer.Bytes())
}

func ArgToInt(value reflect.Value) (int, error) {
	return strconv.Atoi(ArgToString(value))
}

func addCustomFunc(set *jet.Set) {
	set.AddGlobalFunc("url", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("url", 1, 1)

		buffer := bytes.NewBuffer(nil)
		_, _ = fmt.Fprint(buffer, a.Get(0))

		str := string(buffer.Bytes())

		if strings.HasPrefix(str, "http") || strings.HasPrefix(str, "//") || strings.HasPrefix(str, "javascript") {
			return reflect.ValueOf(str)
		}

		return reflect.ValueOf(WebHost + str)
	})

	set.AddGlobalFunc("formatDate", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("formatDate", 1, 1)

		date, ok := a.Get(0).Interface().(time.Time)
		if ok {
			str := helpers.FormatDefaultDate(date)
			return reflect.ValueOf(str)
		}
		return reflect.ValueOf("")
	})

	set.AddGlobalFunc("forList", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("forList", 2, 2)
		bStart := bytes.NewBuffer(nil)
		_, _ = fmt.Fprint(bStart, a.Get(0))
		start := string(bStart.Bytes())
		bEnd := bytes.NewBuffer(nil)
		_, _ = fmt.Fprint(bEnd, a.Get(1))
		end := string(bEnd.Bytes())

		iStart, err := strconv.Atoi(start)
		if err != nil {
			return reflect.ValueOf(make([]int, 0))
		}
		iEnd, err := strconv.Atoi(end)
		if err != nil {
			return reflect.ValueOf(make([]int, 0))
		}
		if iEnd <= iStart {
			return reflect.ValueOf(make([]int, 0))
		}
		result := make([]int, iEnd-iStart+1)
		for i := iStart; i <= iEnd; i++ {
			result[i-iStart] = i
		}
		return reflect.ValueOf(result)
	})
	set.AddGlobalFunc("forPages", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("forPages", 2, 3)

		iLen := 10
		if a.NumOfArguments() == 3 {
			if _iLen, err := ArgToInt(a.Get(2)); err == nil {
				iLen = _iLen
			}
		}
		iTotal, err := ArgToInt(a.Get(0))
		if err != nil {
			return reflect.ValueOf(make([]int, 0))
		}
		iCur, err := ArgToInt(a.Get(1))
		if err != nil {
			return reflect.ValueOf(make([]int, 0))
		}
		if iTotal < 1 {
			return reflect.ValueOf(make([]int, 0))
		}
		if iLen > iTotal {
			iLen = iTotal
		}
		if iCur > iTotal {
			iCur = 1
		}
		iStart := iCur - (iLen/2 - 1)
		if iStart < 1 {
			iStart = 1
		}
		iEnd := iStart + iLen - 1
		if iEnd > iTotal {
			iEnd = iTotal
		}
		if iEnd-iStart+1 != iLen {
			iStart = iEnd - iLen + 1
		}
		if iStart < 1 {
			iStart = 1
		}
		result := make([]int, iLen)
		for i := 0; i < iLen; i++ {
			result[i] = iStart + i
		}
		return reflect.ValueOf(result)
	})

}
