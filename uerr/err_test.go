package uerr

import (
	"errors"
	"strings"
	"testing"
)

func TestUErrBasics(t *testing.T) {
	t.Run("default detail", func(t *testing.T) {
		e := New(CodeInternalError)
		expected := "CodeInternalError"
		if got := e.Detail(); got != expected {
			t.Fatalf("want detail %q, got %q", expected, got)
		}
		if errStr := e.Error(); !strings.Contains(errStr, "[500]") {
			t.Fatalf("unexpected error string %q", errStr)
		}
	})

	t.Run("custom detail", func(t *testing.T) {
		detail := "boom"
		e := NewWithDetail(CodeBadRequest, detail)
		if e.Detail() != detail {
			t.Fatalf("expected %q, got %q", detail, e.Detail())
		}
	})

	t.Run("wrap fmt err", func(t *testing.T) {
		err := errors.New("disk full")
		e := NewWithError(CodeForbidden, err, "write failed")
		if !strings.Contains(e.Detail(), "disk full") {
			t.Fatalf("detail %q missing wrapped error", e.Detail())
		}
	})
}

func TestFormatStackContainsTestFile(t *testing.T) {
	stackErr := buildStack()
	stack := stackErr.FormatStack()
	if !strings.Contains(stack, "err_test.go") {
		t.Fatalf("stack missing err_test.go: %q", stack)
	}
}

func TestToUErr(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		e := ToUErr(nil)
		if e.Code() != CodeSuccess {
			t.Fatalf("want success, got %d", e.Code())
		}
	})

	t.Run("already UErr", func(t *testing.T) {
		e0 := New(CodeMySQLError)
		e := ToUErr(e0)
		if e != e0 {
			t.Fatalf("expected same pointer, got new one")
		}
	})

	t.Run("generic error", func(t *testing.T) {
		err := errors.New("boom")
		e := ToUErr(err)
		if e.Code() != CodeInternalError {
			t.Fatalf("expected internal error, got %d", e.Code())
		}
		if !strings.Contains(e.Detail(), "boom") {
			t.Fatalf("detail %q missing message", e.Detail())
		}
	})
}

func TestRegisterCodeDetail(t *testing.T) {
	const customCode = Code(9999)
	RegisterCodeDetail(map[Code]string{customCode: "custom detail"})
	if detail := customCode.Detail(); detail != "custom detail" {
		t.Fatalf("want custom detail, got %q", detail)
	}
}

func buildStack() *UErr {
	return level1()
}

func level1() *UErr { return level2() }
func level2() *UErr { return level3() }
func level3() *UErr { return New(CodeInternalError) }
