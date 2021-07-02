package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyLinkedListDoesNotContainNext(t *testing.T) {
	sut := ListNode{}
	assert.Nil(t, sut.Next)
}

// func TestAdditionWithInternalCarryOver(t *testing.T) {
// 	assert.Failf(t, "Not implemented", "")
// }

// func TestCarryBiggerThanOne(t *testing.T) {
// 	// TBD next time
// }

func TestAddTwoNumbersWithTableDataStructure(t *testing.T) {
	tableData := []struct {
		desc     string
		input1   int
		input2   int
		expected *ListNode
	}{
		{desc: "TestAddTwoListsWithOneEntryReturningListWithOneEntry", input1: 2, input2: 3, expected: New(5)},
		{desc: "TestAddTwoListsWithOneEntryReturningListWith2Entries", input1: 7, input2: 7, expected: New(14)},
		{desc: "TestAddingTwentyAndFiftyResultSeventy", input1: 20, input2: 50, expected: New(70)},
		{desc: "Twelve plus five results in twelve", input1: 5, input2: 12, expected: New(17)},
		{desc: "TestAdditionWithInternalCarryOver", input1: 1060, input2: 1050, expected: New(2110)},
		{desc: "TestAdditionWithExternalCarryOver", input1: 60, input2: 50, expected: New(110)},
	}
	for _, v := range tableData {
		t.Run(v.desc, func(t *testing.T) {
			actual := addTwoNumbers(New(v.input1), New(v.input2))
			assert.Equal(t, *v.expected, *actual)
		})
	}
}

func TestConstructorForListNodesWithInteger(t *testing.T) {
	tcs := []struct {
		expected ListNode
		input    int
		desc     string
	}{
		{
			desc:     "21",
			input:    21,
			expected: ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			desc:     "22",
			input:    22,
			expected: ListNode{Val: 2, Next: &ListNode{Val: 2}},
		},
		{
			desc:     "1",
			input:    1,
			expected: ListNode{Val: 1},
		},
		{
			desc:     "2",
			input:    2,
			expected: ListNode{Val: 2},
		}, {
			desc:     "54321",
			input:    54321,
			expected: ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		}, {
			desc:     "0",
			input:    0,
			expected: ListNode{Val: 0},
		}, {
			desc:     "504",
			input:    504,
			expected: ListNode{Val: 4, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5}}},
		}, {
			desc:     "500",
			input:    500,
			expected: ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5}}},
		},
	}

	for _, v := range tcs {
		t.Run(v.desc, func(t *testing.T) {
			actual := New(v.input)
			assert.Equal(t, v.expected, *actual)
		})
	}
}
