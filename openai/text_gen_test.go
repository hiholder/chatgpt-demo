package openai

import (
	"fmt"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTextGen(t *testing.T)  {
	c.Convey("test text gen", t, func() {
		text, err := TextClient.GetText("你能背诵多少位圆周率")
		c.So(err,c.ShouldBeNil)
		fmt.Println(text)
	})
}
