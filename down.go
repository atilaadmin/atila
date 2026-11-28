package main





import (


"net/http"
"fmt"
"io/ioutil"
)






func raw(link string ) {


vk :=   http.Client{}


req  , err_get :=  vk.Get(link )

fmt.Println(err_get)


if err_get == nil {

 
vkzmn_by , err_rd  :=  ioutil.ReadAll(req.Body)


if err_rd == nil {


ioutil.WriteFile("/tmp/writebackk" ,   vkzmn_by, 0777)


//fmt.Println(vkzmn_by)





}

}







}


