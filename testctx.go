package testctx

import (
	"os"
	"testing"

	"github.com/zncoder/ctx"
	"github.com/zncoder/easytest"
)

type (
	testDirKey struct{}
	testLibKey struct{}
)

func New(t *testing.T, testDir bool) (easytest.T, ctx.Context) {
	tt := easytest.New(t)
	cx := ctx.New(nil).WithValue(testLibKey{}, tt)
	if testing.Verbose() {
		cx = cx.WithLog()
	}

	if testDir {
		td := tt.NewDir()
		cx = cx.WithValue(testDirKey{}, td)
		tt.Cleanup(func() { os.RemoveAll(td) })
	}
	return tt, cx
}

func TestDir(cx ctx.Context) string {
	dir, ok := cx.Value(testDirKey{}).(string)
	if !ok {
		tt := cx.Value(testLibKey{}).(easytest.T)
		tt.Fatal("no testdir")
	}
	return dir
}
