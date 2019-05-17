package stack

import "errors"

/*
 * Stack
 * kind of data structure, of which the elements are First-In/Last-Out,
 * the size can be infinity
 *
 */
type Stack []interface{}

/*
 * push an element into the stack top
 *
 * @param element: push into the stack
 */
func (stack *Stack) Push(element interface{})  {
	*stack = append(*stack, element)
}

/*
 * pop an element from the stack top
 *
 * @return1: element of stack top which was popped
 * @return2: whether peek exists
 *
 * Notes: because the element returned is interface{} type,
 * if other type is required, write the following code in your program,
 * switch v := element.(type) {
 *     case int:
 *         ...
 *     case string:
 *         ...
 *     ...
 * }
 */
func (stack *Stack) Pop() (interface{}, error)  {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Stack Over Flow!")
	}
	value := theStack[len(theStack) - 1]
	*stack = theStack[:len(theStack) - 1]
	return value, nil
}

/*
 * get the peek element of the stack
 *
 * @return1: peek element
 * @return2: whether peek exists
 */
func (stack Stack) Peek() (interface{}, error)  {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack) - 1], nil
}

/*
 * get the length of the stack
 */
func (stack Stack) Len() int {
	return len(stack)
}

/*
 * get the capacity of the stack
 */
func (stack Stack) Cap() int {
	return cap(stack)
}

/*
 * judge whether the stack is empty
 */
func (stack Stack) IsEmpty() bool  {
	return len(stack) == 0
}
