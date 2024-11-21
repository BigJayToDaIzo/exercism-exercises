package sublist

import (
	"reflect"
)

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if reflect.DeepEqual(l1, l2) {
		return RelationEqual
	}
	if len(l1) == 0 {
		return RelationSublist
	}
	if len(l2) == 0 {
		return RelationSuperlist
	}
	// check for superlist
	if len(l1) > len(l2) {
		for i := 0; i < len(l1)-len(l2)+1; i++ {
			if reflect.DeepEqual(l1[i:i+len(l2)], l2) {
				return RelationSuperlist
			}
		}
		//check for sublist
	} else {
		for i := 0; i < len(l2)-len(l1)+1; i++ {
			if reflect.DeepEqual(l1, l2[i:i+len(l1)]) {
				return RelationSublist
			}
		}
	}
	return RelationUnequal
}
