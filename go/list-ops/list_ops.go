package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	accumulator := initial
	for _, v := range s {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	accumulator := initial
	s = s.Reverse()
	for _, v := range s {
		accumulator = fn(v, accumulator)
	}
	return accumulator
}

func (s IntList) Filter(fn func(int) bool) IntList {
	listSize := 0
	for _, v := range s {
		if fn(v) {
			listSize++
		}
	}
	newList := make([]int, listSize)
	innerIter := 0
	for _, v := range s {
		if fn(v) {
			newList[innerIter] = v
			innerIter++
		}
	}
	return newList
}

func (s IntList) Length() int {
	length := 0
	for range s {
		length++
	}
	return length
}

func (s IntList) Map(fn func(int) int) IntList {
	newList := make([]int, s.Length())
	for i, v := range s {
		newList[i] = fn(v)
	}
	return newList
}

func (s IntList) Reverse() IntList {
	newList := make([]int, s.Length())
	for i, v := range s {
		newList[s.Length()-1-i] = v
	}
	return newList
}

func (s IntList) Append(lst IntList) IntList {
	newList := make([]int, s.Length()+lst.Length())
	for i, v := range s {
		newList[i] = v
	}
	for i, v := range lst {
		newList[i+s.Length()] = v
	}
	return newList
}

func (s IntList) Concat(lists []IntList) IntList {
	finalListSize := s.Length()
	for _, lst := range lists {
		finalListSize += lst.Length()
	}
	newList := make([]int, finalListSize)
	listIter := 0
	for _, v := range s {
		newList[listIter] = v
		listIter++
	}
	for _, lst := range lists {
		for _, v := range lst {
			newList[listIter] = v
			listIter++
		}
	}
	return newList
}
