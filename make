export GOARCH=amd64
export CGO_ENABLED=0



garble -literals -tiny build 


rm -r bin ; mkdir bin
 

upx atila

mv atila bin
 
