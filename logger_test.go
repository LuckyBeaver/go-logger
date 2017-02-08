package log

import (
	"testing"
	"bytes"
	"bufio"
	"github.com/nbio/st"
)

func TestDebug(t *testing.T) {
	tmp := bytes.NewBufferString(``)

	logger := New(tmp, DEBUG)
	SetLogger(logger)
	SetFlags(0)
	Trace(`test trace`)
	Debug(`test debug`)
	Errorf(`%s qwe %d`, `test error`, 123)


	scanner := bufio.NewScanner(tmp)

	scanner.Scan()
	st.Expect(t, scanner.Text(), `[DEBUG] test debug`)

	scanner.Scan()
	st.Expect(t, scanner.Text(), `[ERROR] test error qwe 123`)

	SetLogLevel(TRACE)
	Tracef(`test trace 2`)

	scanner.Scan()
	st.Expect(t, scanner.Text(), `[TRACE] test trace 2`)

	st.Expect(t, scanner.Scan(), false)
}
