package main

import "fmt"

/*
 * FiniteStack
 * kind of data structure, of which the elements are First-In/Last-Out
 *
 * @size: the size of the stack
 * @top: the index(from 0) of the peek element, top==-1 means empty
 * @data: the base slice which implements the stack
 */
type FiniteStack struct {
	size int
	top  int
	data []interface{}
}

/*
 * generate a stack with a given size
 *
 * @param size
 * @return FiniteStack
 */
func MakeStack(size int) FiniteStack {
	q := FiniteStack{}
	q.size = size
	q.top = -1
	q.data = make([]interface{}, size)
	return q
}

/*
 * push an element into the stack top
 *
 * @param element: push into the stack
 * @return bool
 */
func (t *FiniteStack) Push(element interface{}) bool {
	if t.IsFull() {
		return false
	}
	t.top++
	t.data[t.top] = element
	return true
}

/*
 * pop an element from the stack top
 *
 * @return r: element of stack top which was popped
 * @return err: whether peek exists
 *
 * Notes: because the r returned is interface{} type,
 * if other type is required, write the following codes in your program,
 * switch v := r.(type) {
 *     case int:
 *         ...
 *     case string:
 *         ...
 *     ...
 * }
 */
func (t *FiniteStack) Pop() (r interface{}, err error) {
	if t.IsEmpty() {
		err = fmt.Errorf("FiniteStack Over Flow!")
		return
	}
	r = t.data[t.top]
	t.top--
	return
}

/*
 * get the peek element of the stack
 *
 * @return r: peek element
 * @return err: whether peek exists
 */
func (t *FiniteStack) Peek() (r interface{}, err error) {
	if t.IsEmpty() {
		err = fmt.Errorf("FiniteStack Over Flow!")
		return
	}
	r = t.data[t.top]
	return
}


/*
 * get the length of the stack
 */
func (t *FiniteStack) Length() int {
	return t.top+1
}

/*
 * clear the stack
 * it is not necessary to clear, just let top=-1, and cover elements
 */
func (t *FiniteStack) Clear() {
	t.top = -1
}

/*
 * judge whether the stack is empty
 */
func (t *FiniteStack) IsEmpty() bool {
	return t.top == -1
}

/*
 * judge whether the stack is full
 */
func (t *FiniteStack) IsFull() bool {
	return t.top == t.size-1
}

/*
 * traverse the stack from top to bottom
 *
 * @return fn: function that operates stack elements
 * @return isTop2Bottom: whether the traverse is from top to bottom
 */
func (t *FiniteStack) Traverse(fn func(node interface{}), isTop2Bottom bool) {
	if isTop2Bottom {
		for i := 0; i <= t.top; i++ {
			fn(t.data[i])
		}
	} else {
		for i := t.top; i >= 0; i-- {
			fn(t.data[i])
		}
	}

}
