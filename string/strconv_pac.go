package main

import (
	"fmt"
	"strconv" //处理string类型的转换
	"strings"
)

func main() {

	tt := "ab,cT,EC,T"
	str1 := "2ab c82  42342"
	str0 := "sakdsai,iwrewi,0976767"
	const a = `0.63519283478325730257
73547328325673456348767348` //Go中多行的连接

	//在Go当中 string底层是用byte数组存的，并且是不可以改变的 中文字符是用3个字节存的
	//要想访问中文的话，还是要用rune切片，这样就能按下标访问
	str000 := "go编程"
	c := "siueijfierjfierg" +
		"sjdusjfuesrjfue"
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(len(str000))         //8 因为一个汉字 3个字节
	fmt.Println(len([]rune(str000))) //4

	fmt.Println(strings.ToUpper(tt))
	fmt.Println(strings.Split(tt, ","))
	if strings.HasPrefix(str1, "abc") { //判断是否有前缀
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	if strings.HasSuffix(str1, "42") { //判断是否有后缀
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	if strings.Contains(str1, "c82") { //判断是否包含
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	ind1 := strings.Index(str1, "82") //匹配的字符串的第一个字符  所在的索引的位置  从0开始计数
	fmt.Println(ind1)

	ind2 := strings.LastIndex(str1, "4")
	fmt.Println(ind2) //返回最后一个匹配的位置的索引

	str2 := strings.Replace(str1, "2", "8", 2) //把str1中前两个2 变成8  并返回新的字符串
	fmt.Println(str2)

	ind3 := strings.Count(str1, "2") //计算在str1中  字符串“2”的出现次数
	fmt.Println(ind3)

	str3 := strings.Repeat(str1, 2) //把str1字符串重复两次  返回新的字符串
	fmt.Println(str3)

	str4 := strings.ToUpper(str1)
	fmt.Println(str4) //转换为大写

	str5 := strings.ToLower(str4)
	fmt.Println(str5) //转换为小写

	fmt.Println(str1)
	fmt.Println(strings.TrimSpace(str1))      //去除str1首尾的空白符号
	fmt.Println(strings.Trim(str1, "2"))      //去除str1中开头和结尾的 “2”字符串
	fmt.Println(strings.TrimLeft(str1, "2"))  //去除str1中开头的 "2"字符串
	fmt.Println(strings.TrimRight(str1, "2")) //去除str1中结尾的 "2"字符串
	fmt.Println(strings.Fields(str1))         //把str1 以空格分隔  返回一个切片 slice
	fmt.Println(strings.Split(str0, ","))     //用, 去分隔str0 返回一个切片  slice

	slic := []string{"we32", "242dfd", "2343df"} //如何规定了长度就是数组，如果没规定长度  就是切片
	fmt.Println(strings.Join(slic, ","))         //用,把切片slic连接成字符串

	/*
		string的类型转换
	*/
	fmt.Println(strconv.IntSize) //用于获取程序所运行的平台下 int类型所占的位数
	in1 := 100876
	fmt.Println(strconv.Itoa(in1)) //数字转换为字符串
	//	strconv.FormatFloat()          //将64位浮点型的数据转换为字符串

	str00 := "23432"
	in11, _ := strconv.Atoi(str00) //字符串转换为数字
	fmt.Println(in11)

	flo64, _ := strconv.ParseFloat(str00, 64) //将字符串转换为float64
	fmt.Println(flo64)

}