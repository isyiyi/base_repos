package stuManagerPro


// 本文件创建三个结构体
// 1. CommonInfo		通用信息
// 2. PeopleInfo		用户信息
// 3. TypeInfo			类型信息（作为一个字典）
// 4. 返回信息

type CommonInfo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Gender bool `json:"gender"`
	Address string `json:"address"`
	Age int `json:"age"`
	Type int `json:"type"`
}

type PeopleInfo struct {
	CommonInfo
	Major string `json:"major"`
	School string `json:"school"`
}

type ResultInfo struct {
	code int
	err error
	info string
}

var typeInfo map[int]string = make(map[int]string)