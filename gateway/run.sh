#/bin/bash
# sed 's/\r//g' run.sh | bash -
git pull
go build
kill `ps -ef | grep './gateway' | grep -v grep | awk '{print $2}'`
./gateway &