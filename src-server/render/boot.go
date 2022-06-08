package render

import (
	"bytes"
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/zhouhp1295/g3-cms/boot"
	"path"
	"reflect"
	"strings"
	"sync"
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
}
