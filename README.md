# install 
go install github.com/mrco24/parameters@latest

# path setup
cp -r /root/go/bin/parameters /usr/local/bin

# use
parameters -l url.txt -o output.txt
