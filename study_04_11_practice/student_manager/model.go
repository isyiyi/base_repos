package student_manager

// model contains three struct
// 1. commonInfo
// 2. studentInfo
// 3. teacherInfo

type CommonInfo struct {
	id, name string
	gender bool
	age int
}


type StudentInfo struct {
	CommonInfo
	school string
	score float64
	teacherInfo TeacherInfo
}

type TeacherInfo struct {
	commonInfo CommonInfo
	papers []string
	moneyResearch float64
}