/**
* @Author : henry
* @Data: 2020-07-23 19:18
* @Note: map 的初始化、 增删改查
**/

package main

import "fmt"

func main() {
	////  先声明再初始化最后赋值
	//// 先声明一个字典（map）名字叫做henry。key所对应的数据类型是string，value对应的类型也是字符串
	//var henry map[string]string
	//fmt.Printf("判断henry字段是否为空:[%v]\n",henry == nil)
	//fmt.Printf("第一次查看henry字典的值：[%v]\n",henry)
	//
	//henry = make(map[string]string)
	//fmt.Printf("再次判断henry字典知否为空:[%v]\n",henry == nil)
	//fmt.Printf("第二次查看henry字典的值：[%v]\n",henry)
	//
	//henry["name"] = "亨利"
	//fmt.Printf("yingzhengjie字典的类型为:[%v]\n",reflect.TypeOf(henry))
	//fmt.Printf("第三次查看henry字典的值:[%v]\n",henry)

	////初始化之后再赋值
	//henry := make(map[string]int)
	//henry["yzj"] = 25
	//fmt.Println(henry)

	////直接初始化赋值
	//henry := map[string]int{
	//	"亨利":18,
	//	"饼干":20,
	//}
	//fmt.Println(henry)

	////字典的赋值操作
	//henry := make(map[int]string)
	//letter := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	//for k, v := range letter {
	//	henry[7-k] = v
	//}
	//fmt.Println(henry)

	////字典的删除操作
	//henry := make(map[int]string)
	//letter := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	//for k, v := range letter {
	//	henry[k] = v
	//}
	//fmt.Println(henry)
	//
	//for i := 0; i < 4; i++ {
	//	delete(henry, i)
	//}
	//fmt.Println(henry)
	//
	//for k := range henry {
	//	delete(henry, k)
	//}
	//fmt.Println(henry)
	//fmt.Println(henry == nil)

	////字典的修改操作
	//henry := make(map[int]string)
	//letter := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	//
	//for k, v := range letter {
	//	henry[k] = v
	//}
	//fmt.Printf("修改之前的样子：【%v】\n", henry)
	//yzj := henry
	//yzj[0] = "亨利"
	//yzj[1] = "henry"
	//fmt.Printf("修改之后的样子：【%v】\n", henry)

	//// 字典的查询方式
	//henry := make(map[int]string)
	//letter := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	//for k, v := range letter {
	//	henry[k] = v
	//}
	//fmt.Println(henry)
	//for i, j := range henry {
	//	fmt.Println("key = ", i, "value = ", j)
	//}
	//for i := range henry {
	//	fmt.Println(i)
	//}

	// 修改原字典的Value
	type Student struct {
		ID   int
		NAME string
	}
	dict := make(map[int]*Student)
	dict[1] = &Student{
		ID:   100,
		NAME: "henry",
	}
	dict[2] = &Student{
		ID:   200,
		NAME: "亨利",
	}

	fmt.Println(dict[1])
	s := dict[1]
	s.ID = 100000
	fmt.Println(dict)
	fmt.Println(dict[1])

	////判断map键值是否存在
	//henry := make(map[int]string)
	//letter := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	//for k, v := range letter {
	//	henry[k] = v
	//}
	//fmt.Printf("字典中的值为:[%v]\n", henry)
	//if v, ok := henry[1]; ok {
	//	fmt.Println("存在key = ", v)
	//} else {
	//	fmt.Println("没有找到key = ", v)
	//}
	//
	//v, ok := henry[1]
	//if ok {
	//	fmt.Println("再一次确认，已经存在key = ", v)
	//} else {
	//	fmt.Println("再一次确认，没有找到key = ", v)
	//}

	//// map的排序
	//var ProgramingLanguage = map[string]int{
	//	"Java":              0,
	//	"C":                 1,
	//	"C++":               2,
	//	"Python":            3,
	//	"C#":                4,
	//	"PHP":               5,
	//	"JavaScript":        6,
	//	"Visual Basic.NET":  7,
	//	"Perl":              8,
	//	"Assembly language": 9,
	//	"Ruby":              10,
	//}
	//var SortString []string
	//for k := range ProgramingLanguage {
	//	SortString = append(SortString, k)
	//}
	//sort.Strings(SortString)
	//for _, k := range SortString {
	//	fmt.Println("Key:", k, "Value:", ProgramingLanguage[k])
	//}

	//// map的嵌套
	//type EmployeeInformation map[string]int
	//StaffQuarters := make(map[string]EmployeeInformation)
	//EmployeeNumber := make(EmployeeInformation)
	//EmployeeNumber["henry"] = 23
	//EmployeeNumber["bingan"] = 24
	//
	//StaffQuarters["888"] = EmployeeNumber
	//StaffQuarters["999"] = EmployeeInformation{"babinxin":25,"zhouzhiruo":26,}
	//fmt.Println(StaffQuarters)

	//// map的嵌套用法展示
	//Province := make(map[string]map[string][]string) // 定义省字典
	//City := make(map[string][]string)                // 定义市区的字典
	//Scenery := make(map[string][]string)             // 定义景区的字典
	//Scenery["西安"] = []string{"秦始皇兵马俑", "大雁塔", "大唐芙蓉园", "华清池",
	//	"黄巢堡国家森林公园", "西安碑林博物馆", "骊山国家森林公园", "西安城墙",
	//	"秦始皇陵", "翠华山",}
	//
	//Scenery["安康"] = []string{"中坝大峡谷", "香溪洞", "安康双龙生态旅游度假区",
	//	"瀛湖生态旅游景区", "简车湾休闲景区", "南宫山", "天书峡景区", "汉江燕翔洞景区",
	//	"飞渡峡-黄安坝景区", "千层河",}
	//
	//City["西安市区"] = []string{"新城区", "碑林区", "莲湖区", "灞桥区", "未央区", "雁塔区",
	//	"阎良区", "临潼区", "长安区", "高陵区", "咸阳区"}
	//City["安康市区"] = []string{"汉滨区", "汉阴县", "石泉县", "宁陕县", "紫阳县", "岚皋县",
	//	"平利县", "镇平县", "旬阳县", "白河县",}
	//
	//Province["陕西"] = City
	//City["西安景区"] = Scenery["西安"]
	//City["安康景区"] = Scenery["安康"]
	//
	//for k, v := range Province {
	//	fmt.Println(k, v)
	//	for x, y := range v {
	//		fmt.Println(x,y)
	//	}
	//}

	//// map的挖坑
	//Captial := make(map[string]map[string]int)
	//Area := make(map[string]int)
	//Area["大兴区"] = 100
	//Captial["北京"] = Area
	//County := make(map[string]int)
	//County["密云县"] = 200
	//fmt.Println(Captial["北京"]["大兴区"])
	//Captial["北京"] = County
	//fmt.Println(Captial["北京"]["大兴区"])
	//fmt.Println(Captial["北京"]["密云县"])

	//// map的填坑
	//Captial := make(map[string]map[string]int)
	////Area := make(map[string]int)
	////Area["大兴区"] = 100
	////Captial["北京"] = Area
	//
	//if _, ok := Captial["北京"]; ok {
	//	Captial["北京"]["密云县"] = 200
	//} else {
	//	County := make(map[string]int)
	//	Captial["北京"] = County
	//	County["密云县"] = 200
	//}
	//
	//fmt.Println(Captial["北京"]["大兴区"])
	//fmt.Println(Captial["北京"]["密云县"])

	//// 定义一个集合的思想
	//set := make(map[string]bool)
	//set["henry"] = true
	//if set["henry"] {
	//	fmt.Printf("已经存在名为yizhengjie的key，其值为【%v】\n", set["henry"])
	//} else {
	//	fmt.Println("该变量不存在！")
	//}
	//
	//if set["yzj"] {
	//	fmt.Println("已经存在该变量")
	//} else {
	//	fmt.Printf("不存在名为[yzj]的key ，其值为【%v】\n",set["yzj"])
	//}
	//fmt.Println(set)

	// map的寻址
	type NameType struct {
		Blog string
	}
	BlogAddress := make(map[string]*NameType)
	BlogAddress["henry"] = &NameType{"http://www.baidu.com"}

	fmt.Println(BlogAddress)
	fmt.Println(BlogAddress["henry"])
}
