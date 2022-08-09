package student_manager

const(
	_ = iota
	addStu
	delStu
	updateStu
	selStu
)


func MainProcess() {
	var students []StudentInfo
	for ;; {
		option := showUi(len(students))
		if option == -1 || option == 5 {
			break
		}
		if option == addStu {
			students = addStudent(students)
		}else if option == delStu {
			students = delStudent(students)
		}else if option == updateStu {
			updateStudent(students)
		}else if option == selStu {
			selStudent(students)
		}
	}

}