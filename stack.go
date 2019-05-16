package main

import "fmt"

/*
 * Stack
 * kind of data structure, of which the elements are First-In/Last-Out
 *
 * @size: the size of the stack
 * @top: the index of the peek element
 * @data: the base slice which implements the stack
 */
type Stack struct {
	size int
	top  int
	data []interface{}
}

/*
 * generate a stack with a given size
 *
 * @param size
 * @return Stack
 */
func MakeStack(size int) Stack {
	q := Stack{}
	q.size = size
	q.data = make([]interface{}, size)
	return q
}

/*
 * push an element into the stack top
 *
 * @param element: push into the stack
 * @return bool
 */
func (t *Stack) Push(element interface{}) bool {
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
func (t *Stack) Pop() (r interface{}, err error) {
	if t.IsEmpty() {
		err = fmt.Errorf("Stack Over Flow!")
		return
	}
	t.top--
	r = t.data[t.top]
	return
}

/*
 * get the peek element of the stack
 *
 * @return r: peek element
 * @return err: whether peek exists
 */
func (t *Stack) Peek() (r interface{}, err error) {
	if t.IsEmpty() {
		err = fmt.Errorf("Stack Over Flow!")
		return
	}
	r = t.data[t.top]
	return
}


/*
 * get the length of the stack
 */
func (t *Stack) Length() int {
	return t.size
}

/*
 * clear the stack
 * it is not necessary to clear, just let top=0, and cover elements
 */
func (t *Stack) Clear() {
	t.top = 0
}

/*
 * judge whether the stack is empty
 */
func (t *Stack) IsEmpty() bool {
	return t.top == 0
}

/*
 * judge whether the stack is full
 */
func (t *Stack) IsFull() bool {
	return t.top == t.size-1
}

/*
 * traverse the stack from top to bottom
 *
 * @return fn: function that operates stack elements
 * @return isTop2Bottom: whether the traverse is from top to bottom
 */
func (t *Stack) Traverse(fn func(node interface{}), isTop2Bottom bool) {
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
