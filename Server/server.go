package main

import "Server/web"

/*
	程序入口
*/

func main() {
	webapp := web.Newwebapp()
	webapp.Prepare()
}



