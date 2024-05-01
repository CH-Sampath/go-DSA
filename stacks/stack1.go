package stack

type Stack struct {
	Capacity int
	Array    []string
	Top      int
}

var Max_stack int = 10

func CreateStack(s **Stack) {
	*s = new(Stack)
	(*s).Capacity = Max_stack
	(*s).Array = make([]string, (*s).Capacity)
	(*s).Top = -1
}

func IsEmpty(s **Stack) bool {
	if (*s).Top == -1 {
		return true
	}
	return false
}

func IsFull(s **Stack) bool {
	if (*s).Top == (*s).Capacity-1 {
		return true
	} else {
		return false
	}
}

func PopStack(s **Stack) string {
	if IsEmpty(s) {
		return "-1"
	} else {
		ele := (*s).Array[(*s).Top]
		(*s).Array[(*s).Top] = ""
		(*s).Top -= 1
		return ele
	}
}

func PeekStack(s **Stack) string {
	if IsEmpty(s) {
		return "-1"
	} else {
		return (*s).Array[(*s).Top]
	}
}

func DoubleSize(s **Stack) {
	new_cap := (*s).Capacity * 2
	newArray := make([]string, new_cap)
	copy(newArray, (*s).Array)
	(*s).Array = newArray
	(*s).Capacity = new_cap
}

func PushStack(s **Stack, ele string) {
	if IsFull(s) {
		DoubleSize(s)
	}
	(*s).Array[(*s).Top+1] = ele
	(*s).Top += 1
}
