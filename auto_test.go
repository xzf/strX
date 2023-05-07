package strX

import (
	"github.com/xzf/testX"
	"testing"
)

func Test_SubBefore(t *testing.T) {
	testX.Equal(SubBefore(`abc`, "b"), "a")
}

func Test_SubAfter(t *testing.T) {
	testX.Equal(SubAfter(`abc`, "b"), "c")
}

func Test_SubBeforeLast(t *testing.T) {
	testX.Equal(SubBeforeLast(`abbc`, "b"), "ab")

}

func Test_SubAfterLast(t *testing.T) {
	testX.Equal(SubAfterLast(`abbc`, "b"), "c")
}
