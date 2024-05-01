package stack

import "strconv"

type AdvStack struct {
	nrml *Stack
	min  *Stack
}

func CreateAdvStack(a **AdvStack) {
	(*a) = new(AdvStack)
	CreateStack(&((*a).nrml))
	CreateStack(&((*a).min))
}

func PushAdvStack(a **AdvStack, ele string) {
	if (IsEmpty(&(*a).nrml) && IsEmpty(&(*a).nrml)) {
		PushStack(&(*a).nrml, ele)
		PushStack(&(*a).min, ele)
	} else {
		PushStack(&(*a).nrml, ele)
		temp1, _ := strconv.Atoi(ele)
		temp2, _ := strconv.Atoi(PeekStack(&(*a).min))
		if (temp1 <= temp2) {
			PushStack(&(*a).min, ele)
		}
	}
}

func GetMini(a **AdvStack) string{
	if (IsEmpty(&(*a).min)) {
		return "-1"
	} else {
		return PeekStack(&(*a).min)
	}
}

func PopAdvStack(a **AdvStack) string{
	if PeekStack(&(*a).nrml) == PeekStack(&(*a).nrml) {
		temp := PeekStack(&(*a).nrml)
		PopStack(&(*a).nrml)
		PopStack(&(*a).min)
		return temp
	} else {
		return PopStack(&(*a).nrml)
	}
}