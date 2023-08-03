package cmd

import (
	"log"
	"mangosteen/internal/router"
	jwt "mangosteen/internal/services"
)

/*
在go中 外部包的函数名首字母大写表示可以被外部调用
*/
func RunServer() {
	//database.GormConnect()
	//models.CreateUserTest()
	//defer database.GormClose()
	jwt.CareteJwt()
	r := router.New()
	/*
		监听0.0.0.0可以监听所有的网络接口
		如果监听127.0.0则只能在本机访问

		r.Run是不会退出的 除非程序出错
		因为r.Run内部使用了http.ListenAndServe
		而http.ListenAndServe内部使用了http.Server.ListenAndServe
		srv.Serve会一直阻塞执行
	*/
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln("r.Run err:", err) //严重错误，打印错误并退出程序 会调用os.Exit(1)
	}
	log.Println("r.Run的下面的代码不会执行")
}
