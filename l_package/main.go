package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	_ "l_package/pk1"
//)
//
////import (
////	"fmt"
////	p "l_package/person"
////	"l_package/pk1"
////	"l_package/utils"
////	"l_package/utils/timeutils"
////)
//
//// 如果叫init的话，只会被默认调用，不能主动调用
//func init(){
//	fmt.Println("main init()")
//}
//func init1(){
//	fmt.Println("main init1()")
//}
//
//func main() {
//	//utils.Count()
//	//timeutils.PrintTime()
//	//pk1.MyTest1()
//	//utils.MyTest2()
//	//
//	//p1 := p.Person{Name: "王二狗", Sex: "男"}
//	//fmt.Println(p1)
//
//	//fmt.Println("main")
//	//utils.Count()
//	//pk1.MyTest1()
//
//	// init 不能主动调用
//	//init() // undefined: init
//	//init1()
//
//	db, err := sql.Open("mysql", "root:toorex@tcp(127.0.0.1:3306)/testgo_qianfeng?charset=utf8")
//	if err != nil {
//		fmt.Println("错误信息：", err)
//		return
//	}
//	fmt.Println("连接成功：", db)
//
//}
