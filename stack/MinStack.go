//155. Min Stack

package stack

type MinStack struct {
	min   int
	stack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)
}

func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	min := m.stack[0]
	for i := 1; i < len(m.stack); i++ {
		if min > m.stack[i] {
			min = m.stack[i]
		}
	}
	m.min = min
	return m.min
}

//runtime 37ms Beats 6.05%
//Memory 8.94MB Beats 80.38%
