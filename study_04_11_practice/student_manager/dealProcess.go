package student_manager

import "fmt"

func showUi(num int) int {
	fmt.Printf("***********欢迎使用学生信息管理系统************\n"+
		"\t\t当前共有%d个学生\n"+
		"\t\t1. 添加学生信息\n"+
		"\t\t2. 删除学生信息\n"+
		"\t\t3. 修改学生信息\n"+
		"\t\t4. 显示所有学生信息\n"+
		"\t\t5. 退出系统\n"+
		"\t\t请输入你的选择：", num)
	var option int
	_, ok := fmt.Scan(&option)
	if ok != nil {
		fmt.Println("系统出现问题，将强制关闭！！！")
		return -1
	}
	return option
}

func addStudent(students []StudentInfo) []StudentInfo {
	var studentInfo StudentInfo
	fmt.Println("请输入学号：")
	fmt.Scan(&studentInfo.id)
	fmt.Println("请输入姓名：")
	fmt.Scan(&studentInfo.name)
	fmt.Println("请输入性别：")
	var tmpGender string
	fmt.Scan(&tmpGender)
	if tmpGender == "男" {
		studentInfo.gender = true
	} else {
		studentInfo.gender = false
	}
	fmt.Println("请输入年龄：")
	fmt.Scan(&studentInfo.age)
	fmt.Println("请输入学校：")
	fmt.Scan(&studentInfo.school)
	fmt.Println("请输入分数：")
	fmt.Scan(&studentInfo.score)

	fmt.Println("请输入教师信息:")
	fmt.Println("请输入教师姓名：")
	fmt.Scan(&studentInfo.teacherInfo.commonInfo.name)
	fmt.Println("请输入教师的论文名称：（输入-1终止）")
	for {
		var tmpPaper string
		fmt.Scan(&tmpPaper)
		if tmpPaper == "-1" {
			fmt.Println("输入已终止")
			break
		}
		studentInfo.teacherInfo.papers = append(studentInfo.teacherInfo.papers, tmpPaper)
	}
	res := append(students, studentInfo)
	return res
}

func checkStuExist(students []StudentInfo, id string) int {
	for k, _ := range students {
		if students[k].id == id {
			return k
		}
	}
	return -1
}

func delStudent(students []StudentInfo) []StudentInfo {
	var id string
	fmt.Println("请输入你要删除的学生id：")
	fmt.Scan(&id)
	var index int
	if index = checkStuExist(students, id); index == -1 {
		fmt.Println("该学生不存在")
		return students
	}
	students[index] = students[len(students)-1]
	return students[:len(students)-1]
}

func updateStudent(students []StudentInfo) {
	var id string
	fmt.Println("请输入你要修改的学生id：")
	fmt.Scan(&id)

	var index int
	if index = checkStuExist(students, id); index == -1 {
		fmt.Println("该学生不存在")
		return
	}

	fmt.Println("请输入姓名：")
	fmt.Scan(&students[index].name)
	fmt.Println("请输入性别：")
	var tmpGender string
	fmt.Scan(&tmpGender)
	if tmpGender == "男" {
		students[index].gender = true
	} else {
		students[index].gender = false
	}
	fmt.Println("请输入年龄：")
	fmt.Scan(&students[index].age)
	fmt.Println("请输入学校：")
	fmt.Scan(&students[index].school)
	fmt.Println("请输入分数：")
	fmt.Scan(&students[index].score)

	fmt.Println("请输入教师信息:")
	fmt.Println("请输入教师姓名：")
	fmt.Scan(&students[index].teacherInfo.commonInfo.name)
	fmt.Println("请输入教师的论文名称：（输入-1终止）")
	for {
		var tmpPaper string
		fmt.Scan(&tmpPaper)
		if tmpPaper == "-1" {
			fmt.Println("输入已终止")
			break
		}
		students[index].teacherInfo.papers = append(students[index].teacherInfo.papers, tmpPaper)
	}
}

func selStudent(students []StudentInfo) {
	for _, v := range students {
		fmt.Printf("学生id：%s\t学生姓名：%s\t学生年龄：%d\t学校：%s\t成绩：%f\n", v.id, v.name, v.age, v.school, v.score)
		fmt.Printf("教师姓名：%s\t\n他所发过的论文：\n", v.teacherInfo.commonInfo.name)
		for k, v2 := range v.teacherInfo.papers {
			fmt.Printf("%d. %s", k+1, v2)
		}
		fmt.Println()
	}
}
