package openai

import (
	"fmt"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestImageGen(t *testing.T)  {
	c.Convey("test image gen", t, func() {
		url, err := ImageClient.GetImage("A beautiful girl")
		c.So(err, c.ShouldBeNil)
		fmt.Println(url)
	})
}
