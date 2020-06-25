package view

import (
	"fuckdb/bases/model"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	Name = "test"
)

type testView struct{}

func (v *testView) Do(items []model.DB, out io.Writer) error {
	return nil
}

func clearPool() {
	mux.Lock()
	defer mux.Unlock()
	for k := range pool {
		delete(pool, k)
	}
}

func TestRegister(t *testing.T) {
	Convey("视图注册", t, func() {
		defer clearPool()

		Register(Name, new(testView))
		So(SelectViewer(Name), ShouldNotBeNil)

		So(func() { Register(Name, nil) }, ShouldPanicWith, "Register viewer is nil")
		So(func() { Register(Name, new(testView)) }, ShouldPanicWith, "Register called twice for viewer "+Name)
	})
}

func TestRegistered(t *testing.T) {
	Convey("返回已注册视图名称", t, func() {
		defer clearPool()

		So(Registered(), ShouldBeEmpty)

		Register(Name, new(testView))
		So(len(Registered()), ShouldEqual, 1)
		So(Registered()[0], ShouldEqual, Name)
	})
}

func TestSelectViewer(t *testing.T) {
	Convey("根据名称查找视图", t, func() {
		defer clearPool()

		So(SelectViewer(""), ShouldBeNil)
		Register(Name, new(testView))
		So(SelectViewer(Name), ShouldNotBeNil)
	})
}
