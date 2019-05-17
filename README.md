# Stack-in-Go
This repository gives three implementations of stack, kind of data structure, of which the elements are First-In/Last-Out in Go.
1. Stack: the general stack, implemented by a base slice, of which the size can be infinity, with functions: Push(), Pop(), Peek(), Len(), Cap(), IsEmpty();
2. Stack with Top Pointer: stack based 1 but a top pointer "top" and a slice pointer "pData" were added, with functions: MakeStack(), Push(), Pop(), Peek(), Length(), Clear(), IsEmpty(), Traverse().
3. Finite Stack: a finite-length stack, including a size "size", a top pointer "top" and a base slice "data", with functions: MakeStack(), Push(), Pop(), Peek(), Length(), Clear(), IsEmpty(), IsFull(), Traverse().

Notes:
As for functions such as Pop() and Peek(), the element returned is interface{} type. If you want to transform it into type int, type string and other types, write the following code in your program:
```go
switch v := element.(type) {
    case int:
        ...
    case string:
        ...
    ...
}
```
