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
	logger.SetFlags(0)
	logger.Trace(`test trace`)
	logger.Debug(`test debug`)
	logger.Errorf(`%s qwe %d`, `test error`, 123)


	scanner := bufio.NewScanner(tmp)

	scanner.Scan()
	st.Expect(t, scanner.Text(), `[DEBUG] test debug`)

	scanner.Scan()
	st.Expect(t, scanner.Text(), `[ERROR] test error qwe 123`)

	logger.SetlogLevel(TRACE)
	logger.Tracef(`test trace 2`)

	scanner.Scan()
	st.Expect(t, scanner.Text(), `[TRACE] test trace 2`)

	st.Expect(t, scanner.Scan(), false)
}
