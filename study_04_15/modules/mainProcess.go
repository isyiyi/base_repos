package modules

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

var Users []User
var Patients []Patient

const (
	UserPath    = "data/user.txt"
	PatientPath = "data/patient.txt"
	TryCount    = 3
)

func init() {
	// 将之前存储的用户名和密码读取进来
	if readUser() && readPatient() {
		log.Println("数据初始化成功")
	}
}

func MainProcess() {
	defer func() {
		if storeUser() && storePatient() {
			log.Println("数据存储完成")
		} else {
			log.Println("数据存储失败了")
		}
	}()
	showWelcome()

	var isLogin bool = false
	if len(Users) == 0 {
		isLogin = register()
	} else {
		isLogin = login()
	}
	if isLogin {
		showUi()
	}
}

func showWelcome() {
	fmt.Printf("##############################\n" +
		"##    欢迎使用新冠管理系统    ###\n" +
		"##############################\n")
}

func showUi() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("------------------------\n" +
			"       新冠管理系统       \n" +
			"------------------------\n" +
			"1. 新增患者信息\n" +
			"2. 统计患者信息\n" +
			"4. 新增疫苗接种记录\n" +
			"5. 退出系统\n" +
			"请选择对应的功能：\n")
		if input.Scan() {
			option, err := strconv.Atoi(input.Text())
			if err != nil {
				log.Println("请输入合法的选项")
			}
			switch option {
			case 1:
				addInfo()
			case 2:
				staticInfo()
			case 4:
				addRecord()
			case 5:
				log.Println("你已退出系统，再见")
				return
			default:
				log.Println("输入无效，请重新选择")
			}
		}
	}

}

func readUser() bool {
	data, err := ioutil.ReadFile(UserPath)
	if err != nil {
		log.Println("读取用户数据失败")
		return false
	}
	err = json.Unmarshal(data, &Users)
	if err != nil {
		log.Println("用户数据反序列化失败")
	}
	return true
}

func readPatient() bool {
	data, err := ioutil.ReadFile(PatientPath)
	if err != nil {
		log.Println("患者数据读取失败")
		return false
	}
	err = json.Unmarshal(data, &Patients)
	if err != nil {
		log.Println("患者数据反序列化失败")
	}
	return true
}

func storeUser() bool {
	data, err := json.Marshal(Users)
	if err != nil {
		log.Println("用户数据序列化失败")
		return false
	}
	err = ioutil.WriteFile(UserPath, data, fs.FileMode(os.O_RDWR))
	if err != nil {
		log.Println("用户数据写入失败")
		return false
	}
	return true
}

func storePatient() bool {
	data, err := json.Marshal(Patients)
	if err != nil {
		log.Println("患者数据序列化失败")
		return false
	}
	err = ioutil.WriteFile(PatientPath, data, fs.FileMode(os.O_RDWR))
	if err != nil {
		log.Printf("病人数据写入失败%v\n", err)
		return false
	}
	return true
}

func login() bool {
	var user User
	var input = bufio.NewScanner(os.Stdin)
	for i := 0; i < TryCount; i++ {
		fmt.Println("输入用户名及密码：")
		fmt.Println("用户名：")
		if input.Scan() {
			user.Username = input.Text()
		}
		fmt.Println("密码：")
		if input.Scan() {
			user.Password = input.Text()
		}
		for _, v := range Users {
			if v.Username == user.Username && v.Password == user.Password {
				fmt.Println("登陆成功")
				return true
			}
		}
		log.Println("登录失败")
	}
	return false
}

func register() bool {
	var user User
	var input = bufio.NewScanner(os.Stdin)
	for i := 0; i < TryCount; i++ {
		fmt.Println("注册》输入用户名及密码：")
		fmt.Println("用户名：")
		if input.Scan() {
			user.Username = input.Text()
		}
		fmt.Println("密码：")
		if input.Scan() {
			user.Password = input.Text()
		}
		if user.Username != "" && user.Password != "" {
			Users = append(Users, user)
			log.Println("注册成功")
			return true
		} else {
			log.Println("用户名或密码为空，请重新输入")
		}
	}
	return false
}

func addInfo() {
	var patient Patient
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("!!!!!!!!!!新增一条患者记录!!!!!!!!!!")
	for {
		fmt.Println("姓名：")
		if input.Scan() {
			patient.Name = input.Text()
			if len([]rune(patient.Name)) >= 2 && len([]rune(patient.Name)) <= 6 {
				break
			}
			log.Println("姓名长度应该在2-6个字符之间")
		}
	}
	for {
		fmt.Println("性别：")
		if input.Scan() {
			patient.Gender = input.Text()
			if !(patient.Gender != "男" && patient.Gender != "女") {
				break
			}
			log.Println("性别只能是男或女")
		}
	}
	for {
		fmt.Println("年龄：")
		if input.Scan() {
			tmpAge, err := strconv.Atoi(input.Text())
			if err != nil {
				log.Println("年龄的格式不对")
				continue
			}
			if tmpAge >= 1 && tmpAge <= 120 {
				patient.Age = tmpAge
				break
			}
			log.Println("年龄必须在1-120之间")
		}
	}
	fmt.Println("住址：")
	if input.Scan() {
		patient.Address = input.Text()
	}

	for {
		fmt.Println("手机：")
		var flag bool = true
		if input.Scan() {
			patient.Phone = input.Text()
		}
		if len(patient.Phone) != 11 {
			log.Println("手机号码必须由11位数字组成")
			continue
		}
		for _, v := range patient.Phone {
			if string(v) < "0" || string(v) > "9" {
				flag = false
				break
			}
		}
		if flag {
			break
		} else {
			log.Println("手机号码必须由数字组成")
		}
	}

	fmt.Println("邮箱：")
	if input.Scan() {
		patient.Email = input.Text()
	}

	fmt.Println("症状：")
	if input.Scan() {
		patient.Symptom = input.Text()
	}

	fmt.Println("主治医生：")
	if input.Scan() {
		patient.Doctor = input.Text()
	}
	Patients = append(Patients, patient)
	log.Println("添加成功")
}

func staticInfo() {
	fmt.Printf("----------------------------\n" +
		"          新冠管理系统         \n" +
		"----------------------------\n")
	fmt.Printf("姓名\t\t\t性别\t\t年龄\t\t住址\t\t\t手机\t\t\t邮箱\t\t\t症状\t\t主治医生\n")
	sort.Slice(Patients, func(i, j int) bool {
		return Patients[i].Age > Patients[j].Age
	})
	for _, v := range Patients {
		fmt.Printf("%s\t\t%s\t%d\t%s\t\t\t%s\t\t%s\t\t%s\t\t%s\n",
			v.Name, v.Gender, v.Age, v.Address, v.Phone, v.Email, v.Symptom, v.Doctor)
	}
}

func addRecord() {

}
