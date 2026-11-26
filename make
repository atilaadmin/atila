export GOARCH=amd64
export CGO_ENABLED=0



garble -literals build 


rm -r bin ; mkdir bin
 
mv atila bin
 
