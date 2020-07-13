package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
/**
445. 两数相加 II
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nums := changeToNums(l1)
	nums1 := changeToNums(l2)
	fmt.Println(nums, nums1)
	r := []int{}
	lg := len(nums)
	lg1 := len(nums1)
	jw := 0
	if lg >= lg1 {
		for i := 0; i < lg; i++ {
			if i < lg1 {
				temp := nums[i] + nums1[i] + jw
				if jw > 0 {
					jw--
				}
				if temp >= 10 {
					temp = temp - 10
					jw++
				}
				r = append(r, temp)
			} else {
				temp := nums[i] + jw
				if jw > 0 {
					jw--
				}
				if temp >= 10 {
					temp = temp - 10
					jw++
				}
				r = append(r, temp)

			}

		}
	} else {
		for i := 0; i < lg1; i++ {
			if i < lg {
				temp := nums[i] + nums1[i] + jw
				if jw > 0 {
					jw--
				}
				if temp >= 10 {
					temp = temp - 10
					jw++
				}
				r = append(r, temp)
			} else {
				temp := nums1[i] + jw
				if jw > 0 {
					jw--
				}
				if temp >= 10 {
					temp = temp - 10
					jw++
				}
				r = append(r, temp)
			}

		}
	}
	if jw > 0 {
		r = append(r, jw)
		jw--
	}
	return SetRListNode(r)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func SetRListNode(nums []int) (n *ListNode) {
	n = &ListNode{
		Val:  nums[len(nums)-1],
		Next: nil,
	}
	for i := len(nums) - 2; i >= 0; i-- {
		insertEndNode(n, nums[i])
	}
	return
}
func SetListNode(nums []int) (n *ListNode) {
	n = &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	for i := 1; i < len(nums); i++ {
		insertEndNode(n, nums[i])
	}
	return
}
func insertEndNode(l *ListNode, val int) {
	if l.Next == nil {
		l.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		return
	}
	insertEndNode(l.Next, val)
}
func changeToNums(n *ListNode) (nums []int) {
	if n.Next == nil {
		nums = append(nums, n.Val)
		return
	}
	for {
		nums = append(nums, findIntFromListNode(n))
		if n.Next == nil {
			nums = append(nums, n.Val)
			break
		}
	}

	return
}

func findIntFromListNode(n *ListNode) int {
	if n.Next == nil {
		return n.Val
	}
	if n.Next.Next == nil {
		a := n.Next.Val
		n.Next = nil
		return a
	}
	return findIntFromListNode(n.Next)
}

/**
剑指 Offer 61. 扑克牌中的顺子
*/
func isStraight(nums []int) bool {
	l := len(nums)
	count := 0
	if l != 5 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i < l-1; i++ {
		if nums[i] == 0 {
			count++
		} else {
			if nums[i+1]-nums[i] > 1 {
				count = count - (nums[i+1] - nums[i] - 1)
			}
			if nums[i+1]-nums[i] == 0 {
				return false
			}
		}
	}
	if count < 0 {
		return false
	}
	return true
}

/**
1343. 大小为 K 且平均值大于等于阈值的子数组数目
*/
func numOfSubarrays(arr []int, k int, threshold int) int {
	count := 0
	sum := 0
	num := k * threshold
	for i := 0; i < len(arr); i++ {
		if i < k {
			sum += arr[i]
		} else {
			if i == k {
				if sum >= num {
					count++
				}
			}
			sum = sum - arr[i-k] + arr[i]
			if sum >= num {
				count++
			}
		}
	}
	if k == len(arr) {
		if sum >= num {
			count++
		}
	}
	return count
}

/**
1330. 翻转子数组得到最大的数组值
使用数轴表示
a-----b
   c------d
差值为 |c-b|*2  所以需要 b 最小  c 最大 才能获得最大的值
*/
func maxValueAfterReverse(nums []int) int {
	sum := 0
	length := len(nums)
	a := -100000 //区间小值
	b := 100000  //区间大值
	for i := 0; i < length-1; i++ {
		sum += IntAbs(nums[i] - nums[i+1])
		a = IntMax(a, IntMin(nums[i], nums[i+1]))
		b = IntMin(b, IntMax(nums[i], nums[i+1]))
	}
	ans := sum
	ans = IntMax(ans, 2*(a-b)+sum)
	for i := 0; i < length-1; i++ {
		if i > 0 {
			minus := IntAbs(nums[0]-nums[i+1]) - IntAbs(nums[i]-nums[i+1])
			ans = IntMax(ans, sum+minus)
			minus = IntAbs(nums[i-1]-nums[length-1]) - IntAbs(nums[i-1]-nums[i])
			ans = IntMax(ans, sum+minus)
		}
		//for j:=i+1;j<length-1;j++  {
		//	minus:= IntAbs(nums[i]-nums[j])+IntAbs(nums[i+1]-nums[j+1])-(IntAbs(nums[i]-nums[i+1])+IntAbs(nums[j]-nums[j+1]))
		//	ans = IntMax(ans,sum+minus)
		//}
	}

	return ans
}
func IntAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func IntMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

/**
1258 查找双位数
*/
func findNumbers(nums []int) int {
	count := 0
	for _, value := range nums {
		valueStr := strconv.Itoa(value)
		if len(valueStr)%2 == 0 {
			count++
		}
	}
	return count
}

/*
998. 最大二叉树 II
*/
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {

	//right 为空返回 新增又树
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	//节点大于树，原树入左侧 ，节点做根
	if root.Val < val {
		return &TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
654. 最大二叉树  递归解法
*/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return construct(nums, 0, len(nums))
}

func construct(nums []int, l, r int) *TreeNode {

	if l == r {
		return nil
	}
	max_index := max(nums, l, r)
	root := &TreeNode{
		Val:   nums[max_index],
		Left:  construct(nums, l, max_index),
		Right: construct(nums, max_index+1, r),
	}
	return root
}
func max(nums []int, l, r int) int {
	max_index := l
	for i := l; i < r; i++ {
		if nums[i] > nums[max_index] {
			max_index = i
		}
	}
	return max_index
}

/*
1300. 转变数组后最接近目标值的数组和
*/
func findBestValue(arr []int, target int) int {
	//默认排序
	sort.Ints(arr)
	length := len(arr)
	presum := 0
	endLen := length
	for index, value := range arr {
		k := endLen - index
		// 条件  未改变的和 + 当前 value*剩余的项 与 target 的比较
		d := presum + value*k - target
		if d >= 0 {
			//fmt.Println(d,value,endLen,index)
			// c小于等于0.5那么取小 大于0.5 去取上值
			c := value - (d+k/2)/k
			return c
		}
		presum += value
	}

	return arr[length-1]
}

/**
1052. 爱生气的书店老板
*/
func maxSatisfied(customers []int, grumpy []int, X int) int {

	count := 0
	//默认的值
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			count += customers[i]
			//设为0
			customers[i] = 0
		}
	}

	max := 0
	temp := 0
	for i := 0; i < len(customers); i++ {

		//for j:=0 ;j<X;j++{
		//	if grumpy[i+j]==1{
		//		temp+=customers[i+j]
		//	}
		//}
		if i < X {
			max += customers[i]
			if temp < max {
				temp = max
			}
		} else {
			temp = temp + customers[i] - customers[i-X]
			if temp > max {
				max = temp
			}
		}
	}

	return count + max
}

/**
747. 至少是其他数字两倍的最大数
*/
func dominantIndex(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}
	big := nums[0]
	secdbig := nums[1]
	if nums[1] >= nums[0] {
		big = nums[0]
		secdbig = nums[1]
	}

	count := 0
	for i := 1; i < len(nums); i++ {
		if big < nums[i] {
			secdbig = big
			big = nums[i]
			count = i

		}
		if big > nums[i] && nums[i] > secdbig {
			secdbig = nums[i]
		}
	}
	if secdbig*2 <= big {
		return count
	}
	return -1
}

//221. 最大正方形
func maximalSquare(matrix [][]byte) int {

	side := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = byte(int(matrix[i][j]) % 48)
			if i < 1 || j < 1 {
				if matrix[i][j] == 1 && side < 1 {
					side = 1
				}
				continue
			} else {
				if matrix[i][j] == 1 {
					temp := min(min(int(matrix[i-1][j]), int(matrix[i][j-1])), int(matrix[i-1][j-1])) + 1
					matrix[i][j] = byte(temp)
					if temp > side {
						side = temp
					}
				}
			}
		}
	}

	return side * side
}

//1201. 丑数
//请你帮忙设计一个程序，用来找出第 n 个丑数
//丑数是可以被 a 或 b 或 c 整除的 正整数。
//  x/a +x/b+x/c-x/ab-x/ac-x/bc+x/abc
func nthUglyNumber3(n int, a int, b int, c int) int {

	ab := int64(lcm(a, b))
	ac := int64(lcm(a, c))
	bc := int64(lcm(b, c))
	abc := int64(lcm(lcm(a, b), c))
	l := int64(min(a, min(b, c)))
	r := int64(2 * 10e9)
	//while (l < r)
	for {
		if l >= r {
			break
		}
		// 从中间开始查找，每次的偏移量是了l/2
		m := l + (r-l)/2
		count := m/int64(a) + m/int64(b) + m/int64(c) - m/ab - m/ac - m/bc + m/abc
		//计算的数量如果小于 n 则 l= m+1
		if count < int64(n) {
			l = m + 1
		} else {
			//如果大于 n ，则继续二分
			r = m
		}

	}
	return int(l)
}
func min(i int, j int) int {
	if i <= j {
		return i
	}
	return j

}

/*
*公式解法：最小公倍数=两数之积/最大公约数
 */
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

/*
*辗转相除法：最大公约数
*递归写法，进入运算是x和y都不为0
 */
func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}
func canCompleteCircuit(gas []int, cost []int) int {

	//贪心算法
	start := 0
	total := 0
	last := 0
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		if last < 0 {
			last = gas[i] - cost[i]
			start = i
		} else {
			last += gas[i] - cost[i]
		}
	}
	if total >= 0 {
		return start
	}
	return -1
	//	gas := []int{5,1,2,3,4}
	//	cost :=[]int{4,4,1,5,1}
	// for i:=0;i< len(gas);i++{
	//   g :=gas[i]
	//   j :=i
	//   if g<cost[j]{
	//		 continue
	//   }else {
	//	   for {
	//	   	    if j+1>=len(cost){
	//				g = g - cost[j] +gas[0]
	//				if( g<cost[0]){
	//					break
	//				}
	//
	//			}else{
	//				g = g - cost[j] +gas[j+1]
	//				if( g<cost[j+1]){
	//					break
	//				}
	//			}
	//	   	   j++
	//	   	   if(j >=len(cost)){
	//	   	   	j=0
	//		   }
	//		   if(j==i ){
	//			   return i
	//		   }
	//	   }
	//   }
	// }

	//return -1
}

//263. 丑数
//动态规划  1*2  1*3 1*5 加入到切片中 然后 下标+1  继续加入
func nthUglyNumber(n int) int {

	s := []int{}
	s = append(s, 1)

	j := 0
	k := 0
	h := 0
	for i := 0; i < n; i++ {

		min := min(min(s[j]*2, s[k]*3), s[h]*5)
		s = append(s, min)
		if min == s[j]*2 {
			j++
		}
		if min == s[k]*3 {
			k++
		}
		if min == s[h]*5 {
			h++
		}
		if i == (n - 1) {
			return s[i]
		}
	}
	return s[n-1]
}

//263. 丑数
func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	if num == 2 || num == 3 || num == 5 || num == 1 {
		return true
	} else if num%2 == 0 {
		return isUgly(num / 2)
	} else if num%3 == 0 {
		return isUgly(num / 3)
	} else if num%5 == 0 {
		return isUgly(num / 5)
	}
	return false
}

//20. 有效的括号
func isValid(s string) bool {

	var m = make(map[byte]byte)
	m['}'] = '{'
	m[']'] = '['
	m[')'] = '('
	var slice []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			slice = append(slice, s[i])
		} else {
			if len(slice) > 0 {
				if (slice[len(slice)-1]) != m[s[i]] {
					return false
				} else {
					slice = slice[:len(slice)-1]
				}
			} else {
				return false
			}

		}
	}

	return len(slice) == 0
}

//367. 有效的完全平方数
func isPerfectSquare(num int) bool {
	//更优解
	//sumnum := 1
	//for
	//{
	//	num -= sumnum;
	//	sumnum += 2;
	//	if(num<=0){
	//		break
	//	}
	//}
	//return num==0
	n := getnum(num, num)
	for i := n; i < 2*n; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
func getnum(n int, num int) int {
	if n*n > num {
		n = getnum(n/2, num)
	}
	return n

}

//953. 验证外星语词典
func isAlienSorted(words []string, order string) bool {

	mapStr := make(map[byte]int)
	count := len(order)
	//设置大小
	for i := count; i > 0; i-- {
		mapStr[order[count-i]] = i
	}

	mapWords := []map[int]int{}
	for i := 0; i < len(words); i++ {
		temp := make(map[int]int)
		for j := 0; j < len(words[i]); j++ {
			temp[j] = mapStr[words[i][j]]
		}
		mapWords = append(mapWords, temp)
	}
	//[map[0:97 1:112 2:112 3:108 4:101] map[0:97 1:112 2:112]]
	fmt.Println(mapWords)
	for i := 0; i < len(mapWords)-1; i++ {

		for j := 0; j < len(mapWords[i]); j++ {

			m1, ok1 := mapWords[i][j]
			if !ok1 {
				m1 = 27
			}
			m2, ok2 := mapWords[i+1][j]
			if !ok2 {
				m2 = 27
			}
			if m1 < m2 {
				return false
			} else if m1 == m2 {
				continue
			} else {
				break
			}
		}

	}
	return true
}

//874模拟行走机器人
func robotSim(commands []int, obstacles [][]int) int {
	x := 0
	y := 0
	ons := 0

	//var mapStr =""
	//for i := 0; i < len(obstacles); i++ {
	//	mapStr =mapStr+"["+strconv.Itoa(obstacles[i][0])+","+strconv.Itoa(obstacles[i][1])+"]"
	//}
	mapStr := make(map[int][]int)
	for i := 0; i < len(obstacles); i++ {
		mapStr[obstacles[i][0]] = append(mapStr[obstacles[i][0]], obstacles[i][1])
	}
	//op 0 y+ -2or2 y- 1 x+ -1 x-
	var op = 0
	for i := 0; i < len(commands); i++ {
		if commands[i] == -2 {
			if op == 2 || op == -2 {
				op = 1
			} else {
				op--
			}

		} else if commands[i] == -1 {
			if op == 2 || op == -2 {
				op = -1
			} else {
				op++
			}
		} else if commands[i] <= 9 && commands[i] >= 1 {
			switch op {
			case 0:
				for j := 0; j < commands[i]; j++ {
					item, ok := mapStr[x]
					ok1 := false
					for _, v := range item {
						if v == (y + 1) {
							ok1 = true
							break
						}
					}
					if ok && ok1 {
						break
					} else {
						y++
					}
				}

			case 1:
				for j := 0; j < commands[i]; j++ {
					item, ok := mapStr[x+1]
					ok1 := false
					for _, v := range item {
						if v == (y) {
							ok1 = true
							break
						}
					}
					if ok && ok1 {
						break
					} else {
						x++
					}
				}
			case -1:
				for j := 0; j < commands[i]; j++ {
					item, ok := mapStr[x-1]
					ok1 := false
					for _, v := range item {
						if v == (y) {
							ok1 = true
							break
						}
					}
					if ok && ok1 {
						break
					} else {
						x--
					}
				}
			case 2, -2:
				for j := 0; j < commands[i]; j++ {
					item, ok := mapStr[x]
					ok1 := false
					for _, v := range item {
						if v == (y - 1) {
							ok1 = true
							break
						}
					}
					if ok && ok1 {
						break
					} else {
						y--
					}
				}
			}
		}
		if (x*x + y*y) > ons {
			ons = x*x + y*y
		}
	}
	return ons
}

//方阵中战斗力最弱的 K 行
func kWeakestRows(mat [][]int, k int) []int {
	type array struct {
		value int
		index int
	}
	slice1 := make([]array, len(mat))
	slice2 := make([]int, k)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				slice1[i] = array{
					value: j,
					index: i,
				}
				break
			}
			if j == len(mat[i])-1 {
				slice1[i] = array{
					value: j + 1,
					index: i,
				}
			}

		}

	}
	//[2 4 1 2 5]
	for i := 0; i < len(slice1); i++ {
		for j := i + 1; j < len(slice1); j++ {
			if slice1[i].value > slice1[j].value {
				temp := slice1[i]
				slice1[i] = slice1[j]
				slice1[j] = temp
			} else if slice1[i].value == slice1[j].value && slice1[i].index > slice1[j].index {
				temp := slice1[i]
				slice1[i] = slice1[j]
				slice1[j] = temp
			}
		}
	}
	fmt.Println(slice1)
	for i := 0; i < k; i++ {
		slice2[i] = slice1[i].index
	}
	return slice2
}
