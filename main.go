package main

import "fmt"

type User struct {

}
func main()  {
	fmt.Println("准备爬取知乎")
	if Loginzhihu(){
		gethome()
	}else {
		fmt.Println("登录知乎失败")
	}


}
