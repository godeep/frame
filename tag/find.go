package tag

import (
	"bytes"
	"fmt"
)

var (
	Lefts    = [...]byte{'(', '{', '[', '<', '\''}
	Rights   = [...]byte{')', '}', ']', '>', '\''}
	Free     = [...]byte{'"', '\'', '`', '\n'}
	AlphaNum = []byte("*&!%-_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func any(b byte, s []byte) int64{
	for i, v := range s {
		if b == v {
			return int64(i)
		}
	}
	return -1	
}

func isany(b byte, s []byte) bool {
	for _, v := range s {
		if b == v {
			return true
		}
	}
	return false
}

func acceptback(p []byte, i int64, sep []byte) int64 {
	q0 := i
	for ; q0-1 >= 0 && isany(p[q0-1], sep); q0-- {
	}
	return q0
}
func accept(p []byte, j int64, sep []byte) int64 {
	q1 := j
	for ; q1 != int64(len(p)) && isany(p[q1], sep); q1++ {
	}
	return q1
}

func findback(p []byte, i int64, sep []byte) int64 {
	q0 := i
	for ; q0-1 >= 0 && !isany(p[q0-1], sep); q0-- {
	}
	if q0 < 0{
		return i
	}
	return q0
}
func find(p []byte, j int64, sep []byte) int64 {
	q1 := j
	for ; q1 != int64(len(p)) && !isany(p[q1], sep); q1++ {
	}
	if q1 == int64(len(p)){
		return j
	}
	return q1
}

func FindAlpha(p []byte, i int64) (int64, int64) {
	j := accept(p, i, AlphaNum)
	i = acceptback(p, i, AlphaNum)
	return i, j
}

func FindParity(f File) (q0, q1 int64, ok bool) {
	q0, q1 = f.Dot()
	for i := range Lefts {
		q0, q1 = findParity(f, Lefts[i], Rights[i], false)
		if q0 != -1 {
			return q0, q1, true
		}
	}
	return -1, -1, false
}

func findParity(f File, l byte, r byte, back bool) (int64, int64) {
	if back {
		panic("unimplemented")
	}
	/*
	b := t.ReadByte()
	if b != l {
		return -1, -1
	}
	*/
	push := 1
	//j := -1
	q0, _ := f.Dot()
	for i, v := range f.Bytes()[q0:] {
		if v == l {
			push++
		}
		if v == r {
			push--
			if push == 0 {
				return q0, q0+int64(i)
			}
		}
	}
	return -1, -1
}
func FindNext(f File, text []byte) (q0, q1 int64){
	i, j := f.Dot()
	p := f.Bytes()
	x := text
	q0 = int64(bytes.Index(p[j:], x))
	if q0 == -1 {
		println("a")
		q0 = int64(bytes.Index(p[:i], x))
		if q0 < 0 {
			println("b")
			return i, j
		}
	} else {
		println("c")
		q0 += j
	}
	q1 = q0 + int64(len(x))
	println("d")
	return q0, q1
}

func Next(p []byte, i, j int64) (q0 int64, q1 int64) {
	defer func(r0, r1 int64) {
		fmt.Printf("Next: [%d:%d]->[%d:%d]\n", r0, r1, q0, q1)
	}(i, j)
	x := p[i:j]
	q0 = int64(bytes.Index(p[j:], x))
	if q0 == -1 {
		println("a")
		q0 = int64(bytes.Index(p[:i], x))
		if q0 < 0 {
			println("b")
			return i, j
		}
	} else {
		println("c")
		q0 += j
	}
	q1 = q0 + int64(len(x))
	println("d")
	return q0, q1
}

/*
func (t *Tick) Find(p []byte, back bool) int {
	return t.accept(p, t.P1, len(t.Fr.s), back)
}

func (t *Tick) accept(p []byte, i, j int, back bool) int {
	if back {
		panic("unimplemented")
	}
	//fmt.Printf("debug: accept: %q check frame[%d:]\n", p, t.P1)
	x := bytes.Index(t.Fr.s[i:j], p)
	if x == -1 {
		return -1
	}
	println("found at index", i, ":", x+i)
	return x + i

}

func (t *Tick) FindSpecial(i int) (int, int) {
	fmt.Println("NUMBER", i)
	if i == 0 {
		return i, t.FindOrEOF([]byte{'\n'})
	}
	t.Open(i - 1)
	t.Sweep(i)
	t.Commit()
	if t.ReadByte() == '\n' {
		return i, t.FindOrEOF([]byte{'\n'})
	}
	if x := t.FindQuote(); x != -1 {
		return i, x
	}
	if x := t.FindParity(); x != -1 {
		return i, x
	}
	if isany(t.ReadByte(), AlphaNum) {
		return t.FindAlpha(i)
	}
	return i, -1
}

func (t *Tick) FindOrEOF(p []byte) int {
	i := t.Find(p, false)
	if i == -1 {
		return t.Fr.nbytes
	}
	return i
}

func (t *Tick) FindQuote() int {
	b := t.ReadByte()
	for _, v := range Free {
		if b != v {
			continue
		}
		return t.Find([]byte{v}, false)
	}
	return -1
}

*/
