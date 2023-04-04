package main

func main() {
	findCircleNum([][]int{})
}

// 初始时，有N个朋友圈，即每个人一个圈子。
// 如果两个人是朋友，则进行一次merge，如果merge成功，则朋友圈减一。

var fa map[int]int  // 父节点
func init_fa(n int) {
	fa = make(map[int]int,0)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
}
// find+路径压缩
func f(x int) int {
	r := x
	for fa[r] != r { // 第一个for：找到x所在树的根节点r
		r = fa[r]
	}
	for fa[x] != x { // 第二个for：将x-r路径上的所有节点的fa更新为r
		t := fa[x]
		fa[x] = r // 路径上的节点的fa更新为r
		x = t
	}
	return x
}
// 合起来
func merge(u, v int) bool {
	ru := f(u)
	rv := f(v)
	if ru==rv {
		return false
	}
	fa[ru] = rv
	return true
}

func findCircleNum(M [][]int) int {
	n := len(M)
	init_fa(n)  // 初始化
	ans := n
	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			if (M[i][j] == 1) && (merge(i,j)==true) {
				ans--
			}
		}
	}
	return ans
}