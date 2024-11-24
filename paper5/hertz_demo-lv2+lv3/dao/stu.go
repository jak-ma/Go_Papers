package dao

import (
	"encoding/json"
	"fmt"
	"go_papers/paper5/hertz_demo-lv2+lv3/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"time"
)

// 中转数据映射
var Data = map[string]model.Student{}

// 读取yaml文件初始化 data

func LoadDataFromFile() {
	file, err := os.OpenFile("D:\\GO\\gopath\\src\\go_papers\\paper5\\hertz_demo-lv2+lv3\\dao\\stu.yaml", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	decoder.Decode(&Data)
}

// 将 Data 写入 YAML 文件

func SaveDataToFile() {
	file, err := os.OpenFile("D:\\GO\\gopath\\src\\go_papers\\paper5\\hertz_demo-lv2+lv3\\dao\\stu.yaml", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()
	encoder.Encode(Data)
}

// 读取json文件初始化data
func LoadDataFromJson() {
	file, err := os.OpenFile("D:\\GO\\gopath\\src\\go_papers\\paper5\\hertz_demo-lv2+lv3\\dao\\stu.json", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//decoder := json.NewDecoder(file)
	//decoder.Decode(&Data)
	r, _ := ioutil.ReadAll(file)
	json.Unmarshal(r, &Data)
}

// 将Data写入json文件
func SaveDataToJson() {
	file, err := os.OpenFile("D:\\GO\\gopath\\src\\go_papers\\paper5\\hertz_demo-lv2+lv3\\dao\\stu.json", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	p, _ := json.MarshalIndent(Data, "", "   ")
	file.Write(p)
	//encoder := json.NewEncoder(file)
	//encoder.SetIndent("", "    ")
	//encoder.Encode(Data)
}

// 添加学生
func AddStudent(sex, name, id, comefrom string) {
	t := time.Now()
	tTos := t.Format("2006-01-02 15:04:05")
	Data[id] = model.Student{name, id, comefrom, sex, tTos}
	SaveDataToFile()
	//SaveDataToJson()
}

// 根据id判断是否已经存在
func SearchId(id string) bool {
	_, exist := Data[id]
	if !exist {
		return false
	} else {
		return true
	}
}

// 修改学生信息
func ProfileStuInfro(id, newcomefrom string) bool {
	student, exist := Data[id]
	if !exist {
		return false
	}
	student.Comefrom = newcomefrom
	t := time.Now()
	student.Optime = t.Format("2006-01-02 15:04:05")
	Data[id] = student
	SaveDataToFile()
	//SaveDataToJson()
	return true
}
