package strX

import (
	"github.com/xzf/testX"
	"testing"
)

func Test_SubBefore(t *testing.T) {
	testX.Equal(SubBefore(`abc`, "b"), "a")
	testX.Equal(SubBefore(`111你好222`, "你好"), "111")
}

func Test_SubAfter(t *testing.T) {
	testX.Equal(SubAfter(`abc`, "b"), "c")
	testX.Equal(SubAfter(`111你好222`, "你好"), "222")
}

func Test_SubBeforeLast(t *testing.T) {
	testX.Equal(SubBeforeLast(`abbc`, "b"), "ab")
	testX.Equal(SubBeforeLast(`111你好222`, "你好"), "111")

}

func Test_SubAfterLast(t *testing.T) {
	testX.Equal(SubAfterLast(`abbc`, "b"), "c")
	testX.Equal(SubAfterLast(`111你好222`, "你好"), "222")
}
