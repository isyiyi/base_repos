package practice

import "fmt"

func MainProcess() {
	var students [10]string
	showUi(&students)
}

func showUi(students *[10]string) {
	var option int
	for {
		fmt.Println("**********欢迎使用学生管理系统************\n" +
			"1. 添加学生姓名\n" +
			"请输入你要选择的功能：")
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("有点问题哈")
			return
		}
		if option == 1 {
			if addStudent(students) {
				fmt.Println("添加成功")
				printAll(*students)
			}else{
				fmt.Println("添加失败")
			}
		} else {
			fmt.Println("退出系统，再见")
			break
		}
	}
}

func addStudent(students *[10]string) bool {
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return false
	}
	for k, v := range students {
		if v == "" {
			students[k] = name
			return true
		}
	}
	return false
}

func printAll(student [10]string) {
	for _, v := range student {
		fmt.Println(v)
	}
}
