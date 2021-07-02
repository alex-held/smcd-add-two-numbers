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
		{desc: "2 + 3 = 5 -> adding two one digit numbers to one one digit number", input1: 2, input2: 3, expected: New(5)},
		{desc: "7 + 7 = 14 -> adding two one digit numbers to one two digit number", input1: 7, input2: 7, expected: New(14)},
		{desc: "20 + 50 = 70 -> equal amount of digits in numbers adding up correctly", input1: 20, input2: 50, expected: New(70)},
		{desc: "5 + 12 = 17 -> different amount of digits add up correctly", input1: 5, input2: 12, expected: New(17)},
		{desc: "1060 + 1050 = 2110 -> external and internal carryover", input1: 1060, input2: 1050, expected: New(2110)},
		{desc: "9999999 + 9999 = 10009998 -> Testing Example 3 of given expectations", input1: 9999999, input2: 9999, expected: New(10009998)},
		{desc: "99 + 9 = 108 -> external carryover only one time", input1: 99, input2: 9, expected: New(108)},
		{desc: "12 + 8 = 20 -> internal carryover only one time", input1: 12, input2: 8, expected: New(20)},
		{desc: "999 + 2 = 1001 -> external carryover more than 1 time", input1: 999, input2: 2, expected: New(1001)},
		{desc: "0 + 0 = 0 -> Testing Example 2 of given expectation", input1: 0, input2: 0, expected: New(0)},
		{desc: "342 + 465 = 807 -> Testing Example 1 of given expectation", input1: 342, input2: 465, expected: New(807)},
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
