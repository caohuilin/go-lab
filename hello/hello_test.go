package hello

import (
	"math"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// https://leetcode.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	for i, a := range nums {
		for j, b := range nums {
			if i == j {
				continue
			}
			if a+b == target {
				return []int{i, j}
			}
		}
	}
	panic("no ans")
}

// https://leetcode.com/problems/reverse-integer/description/
func reverse(x int) int {
	sum := 0
	for {
		if x == 0 {
			break
		}
		a := x % 10
		sum = sum*10 + a
		x = x / 10
	}
	if sum > math.MaxInt32 {
		return 0
	}
	if sum < math.MinInt32 {
		return 0
	}
	return sum
}

// https://leetcode.com/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func TestSpec(t *testing.T) {
	Convey("twoSum", t, func() {
		So(twoSum([]int{2, 7, 11, 15}, 9), ShouldResemble, []int{0, 1})
		So(twoSum([]int{2, 7, 11, 15}, 26), ShouldResemble, []int{2, 3})
		So(twoSum([]int{1, 2}, 3), ShouldResemble, []int{0, 1})
		So(twoSum([]int{3, 2, 4}, 6), ShouldResemble, []int{1, 2})
	})

	Convey("reverse", t, func() {
		So(reverse(123), ShouldResemble, 321)
		So(reverse(-123), ShouldResemble, -321)
		So(reverse(-120), ShouldResemble, -21)
		So(reverse(120), ShouldResemble, 21)
		So(reverse(0), ShouldResemble, 0)
		So(reverse(1), ShouldResemble, 1)
		So(reverse(-1), ShouldResemble, -1)
		So(reverse(1534236469), ShouldResemble, 0)
		So(reverse(-2147483648), ShouldResemble, 0)
	})

	Convey("Palindrome Number", t, func() {
		So(isPalindrome(123), ShouldResemble, false)
		So(isPalindrome(121), ShouldResemble, true)
		So(isPalindrome(1), ShouldResemble, true)
		So(isPalindrome(11), ShouldResemble, true)
	})
}
