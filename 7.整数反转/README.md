### 7. 整数反转
 
##### 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
#### 示例 1:
```
输入: 123
输出: 321
示例 2:
        
输入: -123
输出: -321
示例 3:
 
输入: 120
输出: 21
```

#### 解：数学运算模拟栈

```go
func reverse(x int) int {
	ans := 0
	for x != 0 {
		if ans > math.MaxInt32 /10 || ans < math.MinInt32 / 10 {
			return 0
		}
		ans = ans * 10
		ans += x % 10
		x /= 10
	}
	return ans
}
```
1. 在防止溢出的前提下：利用数学原理进行模拟栈的 Pop 与 Push 操作
2. 其中：假设 x 为需要反转的数，Pop 操作则是 x % 10，对其取余数
3. Push 操作为：``temp = x * 10 + pop; x = temp;``
4. 反复执行操作
5. 最终获取反转后的数
```
Pop（取个位）: 
x = 123
用 10 取余
则拿到了个位 3，相当于获取栈的顶层

Push
同样的：
x = 123
先将 x 乘以 10，再加上个位数（即需要 Pop 的数），存储到 temp 
此时 temp = 123 * 10 + 3 = 1233
互换变量
x = 1233（相当于 Push 了 3）
```