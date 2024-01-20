package collections

import "testing"

func assertEmptyStack[T any](t *testing.T, s Stack[T]) {
	if got := s.IsEmpty(); got != true {
		t.Errorf("IsEmpty() = %v, want %v", got, true)
	}
	if got := s.HasElements(); got != false {
		t.Errorf("HasElements() = %v, want %v", got, false)
	}
}

func assertNotEmptyStack[T any](t *testing.T, s Stack[T]) {
	if got := s.IsEmpty(); got != false {
		t.Errorf("IsEmpty() = %v, want %v", got, true)
	}
	if got := s.HasElements(); got != true {
		t.Errorf("HasElements() = %v, want %v", got, true)
	}
}

func Test_uint8Stack_IsEmpty(t *testing.T) {
	t.Parallel()

	t.Run(``, func(t *testing.T) {

		s := NewStack[uint8](10)

		assertEmptyStack(t, s)

		s.Push(1)
		s.Push(2)
		s.Push(3)

		assertNotEmptyStack(t, s)
		want := uint8(3)
		if got := s.Pop(); got != want {
			t.Errorf("pop() = %v, want %v", got, want)
		}

		assertNotEmptyStack(t, s)
		want = uint8(2)
		if got := s.Pop(); got != want {
			t.Errorf("pop() = %v, want %v", got, want)
		}

		assertNotEmptyStack(t, s)
		want = uint8(1)
		if got := s.Pop(); got != want {
			t.Errorf("pop() = %v, want %v", got, want)
		}

		assertEmptyStack(t, s)

		defer func() {
			if err := recover(); err != nil {
				if err.(string) != `stack is empty` {
					t.Error(`pop() on empty stack must throw error with message "stack is empty"`)
				}
			} else {
				t.Error(`pop() on empty stack must throw error`)
			}
		}()

		// must call panic
		s.Pop()
	})
}
