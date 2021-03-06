package frame

import (
	"fmt"
	"github.com/as/etch"
	"image"
	"testing"
)

const (
	modeSaveResult = iota
	modeCheckResult
)

func check(t *testing.T, have image.Image, name string, mode int) {
	wantfile := fmt.Sprintf("testdata/%s.expected.png", name)
	if mode == modeSaveResult {
		etch.WriteFile(t, wantfile, have)
	}
	etch.AssertFile(t, have, wantfile, fmt.Sprintf("%s.png", name))
}

func TestSelect0to0(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	h.Select(0, 0)
	check(t, have, "TestSelect0to0", modeCheckResult)
}

func TestSelect0to1(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	h.Select(0, 1)
	check(t, have, "TestSelect0to1", modeCheckResult)
}
func TestSelectLine(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	h.Select(0, 12)
	check(t, have, "TestSelectLine", modeCheckResult)
}

func TestSelectLinePlus(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	h.Select(0, 13)
	check(t, have, "TestSelectLinePlus", modeCheckResult)
}

func TestSelectLinePlus1(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	h.Select(0, 13+1)
	check(t, have, "TestSelectLinePlus1", modeCheckResult)
}

func TestSelectAll(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	h.Select(0, 9999)
	check(t, have, "TestSelectAll", modeCheckResult)
}

func TestSelectAllSub1(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	h.Select(0, h.Len())
	p0, p1 := h.Dot()
	p1--
	h.Select(p0, p1)
	check(t, have, "TestSelectAllSub1", modeCheckResult)
}

func TestSelectAllSubAll(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	h.Select(0, h.Len())
	h.Select(0, 0)
	check(t, have, "TestSelectAllSubAll", modeCheckResult)
}

func TestMidToEnd(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	h.Select(h.Len()/2, h.Len())
	check(t, have, "TestMidToEnd", modeCheckResult)
}
func TestMidToEndThenStartToMid(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	h.Select(h.Len()/2, h.Len())
	h.Select(0, h.Len()/2)
	check(t, have, "TestMidToEndThenStartToMid", modeCheckResult)
}

func TestSelectTabSpaceNewline(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	for j := 0; j < 5; j++ {
		h.Insert([]byte("abc\t \n\n\t $\n"), int64(j))
	}
	h.Select(h.Len()/2, h.Len()-5)
	check(t, have, "TestSelectTabSpaceNewline", modeCheckResult)
}
func TestSelectTabSpaceNewlineSub1(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	for j := 0; j < 5; j++ {
		h.Insert([]byte("abc\t \n\n\t $\n"), int64(j))
	}
	h.Select(h.Len()/2, h.Len()-5-1)
	check(t, have, "TestSelectTabSpaceNewlineSub1", modeCheckResult)
}
func TestSelectEndLineAndDec(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	h.Select(167+9, 168+9)
	check(t, have, "TestSelectEndLineAndDec", modeCheckResult)
}

var testSelectData = []byte(`Hello world.
Your editor doesn't always know best.
	Your empty file directory has been deleted.
func main(){
	for i := 0; i < 100; i++{
		// comment
	}
}
$ Editor (vi or emacs)?
Usenet is like letters to the editor, only without an editor.  - Larry Wall
Type C-h for help; C-x u to undo changes.  ('C-' means use CTRL key.) GNU Emacs comes with ABSOLUTELY NO WARRANTY; type C-h C-w for full details.You may give out copies of Emacs; type C-h C-c to see the conditions.Type C-h t for a tutorial on using Emacs.





`)

// Broken tests that need work

// TODO(as): regenerate without trailing broken tickmark and rerun this test
func TestSelectNone(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	check(t, have, "TestSelectNone", modeCheckResult)
}

// TODO(as): regenerate without trailing broken tickmark and rerun this test
func TestSelectNoneUntick(t *testing.T) {
	h, _, have, _ := abtestbg(R)
	h.Insert(testSelectData, 0)
	check(t, have, "TestSelectNoneUntick", modeCheckResult)
}
