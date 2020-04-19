// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  dy_pro
 * @Version: 1.0.0
 * @Date: 2020/4/10 21:13
 */

package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	f1 := 1
	f2 := 2
	f3 := 0
	for i := 3; i <= n; i++ {
		f3 = f1 + f2

		f1 = f2
		f2 = f3
	}

	return f3
	//return climb_s(0, n)
}

func climb_s(i, n int) int {
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	return climb_s(i+1, n) + climb_s(i+2, n)
}
