package main

import "fmt"

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	//初始化课程表
	var courses = make(map[int][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		courses[i] = make([]int, 0)
	}

	//遍历所有的courses得到当前所有课程的先修课程列表
	for i := 0; i < len(prerequisites); i++ {
		courses[prerequisites[i][0]] = append(courses[prerequisites[i][0]], prerequisites[i][1])
	}
	//记录是否无法修课
	var n int

	var flag = false
	var f = func() {
		var temp int

		//从课程中删除所有当前可以修的课程
		for k, v := range courses {
			if len(v) == 0 {
				temp++
				continue
			}

			var preCourses []int

			for i := 0; i < len(v); i++ {
				//如果不在目标修的课程中则保留
				if len(courses[v[i]]) > 0 {
					preCourses = append(preCourses, v[i])
				}

			}

			courses[k] = preCourses

		}
		if temp == n {
			flag = true
		} else {
			n = temp
		}
	}

	for n < numCourses {
		f()
		if flag {
			return false
		}
	}

	return true
}
