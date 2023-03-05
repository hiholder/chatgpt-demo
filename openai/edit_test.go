package openai

import (
	"fmt"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEditText(t *testing.T)  {
	c.Convey("test edit", t, func() {
		edit, err := EditClient.GetEdit("What day of the wek is it?", "Fix the spelling mistakes")
		c.So(err, c.ShouldBeNil)
		fmt.Println(edit)
	})
}
