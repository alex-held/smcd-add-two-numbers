package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
	"testing/quick"
)

func TestEmptyLinkedListDoesNotContainNext(t *testing.T) {
	sut := ListNode{}
	assert.Nil(t, sut.Next)
}

func TestAddTwoNumbersWithTableDataStructure(t *testing.T) {
	tableData := []struct {
		desc     string
		input1   int
		input2   int
		expected *ListNode
	}{
		{desc: "2 + 3 = 5 -> adding two one digit numbers to one one digit number", input1: 2, input2: 3, expected: FromInt(5)},
		{desc: "7 + 7 = 14 -> adding two one digit numbers to one two digit number", input1: 7, input2: 7, expected: FromInt(14)},
		{desc: "20 + 50 = 70 -> equal amount of digits in numbers adding up correctly", input1: 20, input2: 50, expected: FromInt(70)},
		{desc: "5 + 12 = 17 -> different amount of digits add up correctly", input1: 5, input2: 12, expected: FromInt(17)},
		{desc: "1060 + 1050 = 2110 -> external and internal carryover", input1: 1060, input2: 1050, expected: FromInt(2110)},
		{desc: "9999999 + 9999 = 10009998 -> Testing Example 3 of given expectations", input1: 9999999, input2: 9999, expected: FromInt(10009998)},
		{desc: "99 + 9 = 108 -> external carryover only one time", input1: 99, input2: 9, expected: FromInt(108)},
		{desc: "12 + 8 = 20 -> internal carryover only one time", input1: 12, input2: 8, expected: FromInt(20)},
		{desc: "999 + 2 = 1001 -> external carryover more than 1 time", input1: 999, input2: 2, expected: FromInt(1001)},
		{desc: "0 + 0 = 0 -> Testing Example 2 of given expectation", input1: 0, input2: 0, expected: FromInt(0)},
		{desc: "342 + 465 = 807 -> Testing Example 1 of given expectation", input1: 342, input2: 465, expected: FromInt(807)},
		{desc: "1234 + 1 = 1235 -> Second Argument shorter than first one", input1: 1234, input2: 1, expected: FromInt(1235)},
	}
	for _, v := range tableData {
		t.Run(v.desc, func(t *testing.T) {
			actual := addTwoNumbers(FromInt(v.input1), FromInt(v.input2))
			assert.Equal(t, *v.expected, *actual)
		})
	}
}

func TestFuzzFails(t *testing.T) {
	tableData := []struct {
		desc     string
		input1   string
		input2   string
		expected *ListNode
	}{
		{desc: "3390679507942867972 + 736882823585834720 = ?", input1: "3390679507942867972", input2: "736882823585834720", expected: FromInt(4127562331528702692)},
		{
			desc: "Two 100 digit numbers add together correctly -> 100 digit result",
			input1: strings.Repeat("1", 100),
			input2: strings.Repeat("1", 100),
			expected: FromString(strings.Repeat("2",100)),
		},{
			desc: "100 digit number and 1 adds to 101 digit result -> Carry Over Large Number",
			input1: strings.Repeat("9", 100),
			input2: "1",
			expected: FromString("1" + strings.Repeat("0",100)),
		},
	}
	for _, v := range tableData {
		t.Run(v.desc, func(t *testing.T) {
			actual := addTwoNumbers(FromString(v.input1), FromString(v.input2))
			assert.Equal(t, *v.expected, *actual)
		})
	}
}

func TestConstructorForListNodesWithInteger(t *testing.T) {
	tcs := []struct {
		expected *ListNode
		input    int
		desc     string
	}{
		{
			desc:     "-1",
			input:    -1,
			expected: nil,
		},
		{
			desc:     "21",
			input:    21,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			desc:     "22",
			input:    22,
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 2}},
		},
		{
			desc:     "1",
			input:    1,
			expected: &ListNode{Val: 1},
		},
		{
			desc:     "2",
			input:    2,
			expected: &ListNode{Val: 2},
		}, {
			desc:     "54321",
			input:    54321,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		}, {
			desc:     "0",
			input:    0,
			expected: &ListNode{Val: 0},
		}, {
			desc:     "504",
			input:    504,
			expected: &ListNode{Val: 4, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5}}},
		}, {
			desc:     "500",
			input:    500,
			expected: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5}}},
		},
	}

	for _, v := range tcs {
		t.Run(v.desc, func(t *testing.T) {
			actual := FromInt(v.input)
			assert.Equal(t, v.expected, actual)
		})
	}
}

func TestStringConstructorForListNodesWithInteger(t *testing.T) {
	tcs := []struct {
		expected *ListNode
		input    string
		desc     string
	}{
		{
			desc:     "-1",
			input:    "-1",
			expected: nil,
		},
		{
			desc:     "21",
			input:    "21",
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			desc:     "22",
			input:    "22",
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 2}},
		},
		{
			desc:     "1",
			input:    "1",
			expected: &ListNode{Val: 1},
		},
		{
			desc:     "2",
			input:    "2",
			expected: &ListNode{Val: 2},
		}, {
			desc:     "54321",
			input:    "54321",
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		}, {
			desc:     "0",
			input:    "0",
			expected: &ListNode{Val: 0},
		}, {
			desc:     "504",
			input:    "504",
			expected: &ListNode{Val: 4, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5}}},
		}, {
			desc:     "500",
			input:    "500",
			expected: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5}}},
		},{
			desc: strings.Repeat("1", 100),
			input: strings.Repeat("1", 100),
			expected: func() *ListNode {
				result := &ListNode{Val: 1}
				current := result
				for i := 0; i < 99; i++ {
					current.Next = &ListNode{Val: 1}
					current = current.Next
				}
				return result
			}(),
		},
	}

	for _, v := range tcs {
		t.Run(v.desc, func(t *testing.T) {
			actual := FromString(v.input)
			assert.Equal(t, v.expected, actual)
		})
	}
}

// TODO: Use FromString ListNode Constructor, to get large Numbers working correctly
func TestPropertiesOfAddTwoNumbers(t *testing.T) {
	assertion := func(l1, l2 int) bool {
		if l1 < 0 || l2 < 0 {
			log.Printf("l1: %d,l2: %d\n", l1, l2)
			return true
		}

		var listNode1 = FromInt(l1)
		var listNode2 = FromInt(l2)
		var result = addTwoNumbers(listNode1, listNode2)
		var expected = FromInt(l1 + l2)
		return result == expected
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}