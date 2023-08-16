package test_linkedList

import (
	"testing"

	. "github.com/wdlea/GOGenericLinkedList"
)

func TestLinkedList(t *testing.T) {
	l := new(LinkedList[int])
	for i := 0; i < 10; i++ {
		//add value to the end of the linked list
		node := l.AddLast(i)
		t.Log(node.Value)

	}

	if l.First.Value != 0 {
		t.Fatalf("First list element(%d) was not 0", l.First.Value)
	}
	if l.Last.Value != 9 {
		t.Fatalf("Last list element(%d) was not 9", l.Last.Value)
	}

	l.First.Prepend(-1, l) //add -1 element

	popped := l.Last.Pop(l)  //delete last element
	l.Last.Append(popped, l) //add last element back

	i := -1
	//iterate over list from first to last
	for value := l.First; value != nil; value = value.Next {
		if i != value.Value {
			t.Fatalf("Linked list order wrong")
		}
		i++
	}

	i--

	//iterate over list from last to first
	for value := l.Last; value != nil; value = value.Previous {
		if i != value.Value {
			t.Fatalf("Linked list order wrong")
		}
		i--
	}

}
