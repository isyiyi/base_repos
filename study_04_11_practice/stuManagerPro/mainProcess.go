package stuManagerPro

import (
	"fmt"
	"log"
)

// 整个程序的主处理文件

// 主界面
func MainUi() {
	var option int
	var peopleInfos []PeopleInfo
	for {
		fmt.Printf("1. 添加类型\n2. 修改类型\n3. 删除类型\n4.保存类型到本地\n" +
			"5. 添加学生\n6.退出系统\n7. 保存结构体到本地\n")
		_, err := fmt.Scan(&option)
		if err != nil {
			log.Fatalln("程序出现错误，已终止")
		}
		switch option {
		case 1, 2:
			var typeId int
			var typeStr string
			_, err = fmt.Scan(&typeId)
			_, err = fmt.Scan(&typeStr)
			var res ResultInfo = operateType(typeId, typeStr, 1)
			if res.code == 200 {
				fmt.Println("成功")
			} else {
				fmt.Println(res.info)
			}
			break
		case 3:
			var typeId int
			_, err = fmt.Scan(&typeId)
			var res ResultInfo = operateType(typeId, "", 3)
			if res.code == 200 {
				fmt.Println("成功")
			} else {
				fmt.Println(res.info)
			}
			break
		case 4:
			s, _ := mapConvJson()
			res := writeStrFile(s)
			fmt.Println(res.info)
			break
		case 5:
			var peopleInfo PeopleInfo
			fmt.Scan(&peopleInfo.Id)
			fmt.Scan(&peopleInfo.Name)
			fmt.Scan(&peopleInfo.Address)
			fmt.Scan(&peopleInfo.Age)
			fmt.Scan(&peopleInfo.Type)
			fmt.Scan(&peopleInfo.Major)
			fmt.Scan(&peopleInfo.School)
			var gender string
			fmt.Scan(&gender)
			if gender == "男" {
				peopleInfo.Gender = true
			}else{
				peopleInfo.Gender = false
			}
			peopleInfos = append(peopleInfos, peopleInfo)
			break
		case 7:
			s, _ := structConvJson(peopleInfos)
			writeStrFile(s)
		}
		if option == 6 {
			fmt.Println("系统关闭")
			break
		}
	}
}

// 增加、修改、删除用户类型
func operateType(typeId int, typeStr string, operateType int) ResultInfo {
	var res ResultInfo
	// 增加
	// 增加和修改的逻辑可以写在一块：直接根据key操作，不存在就增加没存在就修改
	if operateType == 1 || operateType == 2 {
		typeInfo[typeId] = typeStr
		res = ResultInfo{200, nil, fmt.Sprintf("类型id：%d对应的类型被新建或修改", typeId)}
		return res
	} else { // 删除
		_, ok := typeInfo[typeId]
		if ok {
			delete(typeInfo, typeId)
			res = ResultInfo{200, nil, fmt.Sprintf("类型id：%d已经被删除", typeId)}
		} else {
			res = ResultInfo{404, nil, fmt.Sprintf("类型id：%d没有被找到", typeId)}
		}
		return res
	}
}
