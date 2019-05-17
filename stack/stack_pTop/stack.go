package stack_pTop

import "fmt"

/*
 * Stack
 * kind of data structure, of which the elements are First-In/Last-Out,
 * the size can be infinity
 *
 * @top: the index(from 0) of the peek element, top==-1 means empty
 * @pData: the pointer to the base slice which implements the stack
 */
type Stack struct {
	top   int
	pData *[]interface{}
}

/*
 * generate a stack
 *
 * @return Stack
 */
func MakeStack() Stack {
	q := Stack{}
	q.top = -1
	data := []interface{}{}
	q.pData = &data
	return q
}

/*
 * push an element into the stack top
 *
 * @param element: push into the stack
 */
func (t *Stack) Push(element interface{}) {
	t.top++
	*t.pData = append(*t.pData, element)
	return
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
	r = (*t.pData)[t.top]
	*t.pData = (*t.pData)[:t.top]
	t.top--
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
	r = (*t.pData)[t.top]
	return
}


/*
 * get the size of the stack
 */
func (t *Stack) Size() int {
	return t.top+1
}

/*
 * clear the stack
 * it is not necessary to clear, just let top=-1, and cover elements
 */
func (t *Stack) Clear() {
	t.top = -1
}

/*
 * judge whether the stack is empty
 */
func (t *Stack) IsEmpty() bool {
	return t.top == -1
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
			fn((*t.pData)[i])
		}
	} else {
		for i := t.top; i >= 0; i-- {
			fn((*t.pData)[i])
		}
	}

}
