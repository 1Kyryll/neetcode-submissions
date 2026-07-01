func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := NewStack()
	m := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c == ")" || c == "}" || c == "]" {
			last := stack.Pop()

			if last == "" || m[last] != c {
				return false
			}
		} else {
			stack.Push(c)
		}
	}

	return stack.IsEmpty()
}

type Stack struct {
	elements []string
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]string, 0),
	}
}

func (s *Stack) Push(data string) {
	s.elements = append(s.elements, data)
}

func (s *Stack) Pop() string {
	if len(s.elements) == 0 {
		return ""
	}

	n := len(s.elements) - 1
	data := s.elements[n]
	s.elements = s.elements[:n]
	return data
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
