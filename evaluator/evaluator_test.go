package evaluator

import "testing"

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range test {
		evaluated := testEval(tt.input)
		teestIntegerObject(t, evaluated, tt.expected)
	}
}
