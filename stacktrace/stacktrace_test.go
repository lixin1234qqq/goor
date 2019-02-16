package stacktrace

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStacktrace(t *testing.T) {
	expect := []string{
		"github.com/baidu/openrasp/stacktrace.TestStacktrace.func1",
		"runtime.call32",
		"runtime.gopanic",
		"github.com/baidu/openrasp/stacktrace.(*panicker).panic",
		"github.com/baidu/openrasp/stacktrace.TestStacktrace",
	}
	defer func() {
		err := recover()
		if err == nil {
			t.FailNow()
		}
		allFrames := AppendStacktrace(nil, 1, 5)
		functions := make([]string, len(allFrames))
		for i, frame := range allFrames {
			functions[i] = frame.Function
		}
		if diff := cmp.Diff(functions, expect); diff != "" {
			t.Fatalf("%s", diff)
		}
	}()
	(&panicker{}).panic()
}

type panicker struct{}

func (*panicker) panic() {
	panic("oh noes")
}
