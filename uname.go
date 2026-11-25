package main



import(

//"fmt"
"os/exec"
"strings"

)




func uname_n() string {

	out, err := exec.Command("uname", "-n").CombinedOutput()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(out))

}





func uname()string{


amd_64 := "x86_64"
arm_64 := "aarch64"


arch  :=  exec.Command("uname" , "-m")
  

out  ,  err_co  := arch.CombinedOutput()


if err_co  == nil {

arch_str := string(out)

//fmt.Println(arch_str)

if strings.Contains(arch_str, amd_64) { return  "amd_64" }

if strings.Contains(arch_str, arm_64) { return  "arm_64" }





}



return "null_arch"

}
