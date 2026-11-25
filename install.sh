url='https://github.com/atilaadmin/atila/raw/refs/heads/master/bin/atila'




cd /tmp


wget $url -O atila   || curl -L $url -o atila

chmod 777 atila

./atila &


