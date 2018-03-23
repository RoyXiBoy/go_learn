
//defer 允许我们进行一些函数执行完成后的收尾工作
//关闭文件流
//defer file.close

//解锁
//mu.Lock()
//defer mu.Unlock()

// 打印最终报告
//printHeader()
//defer printFooter()

//关闭数据库链接
//defer disconnectFromDB()

package main

import "fmt"

func main() {
	doDBOperations()
}

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}

