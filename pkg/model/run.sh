MYDIR=`dirname $0`
INCLUDE=$GOPATH/include
for file in ./*
do
    if [ "${file##*.}"x = "proto"x ]; then
        echo "$file"
        protoc --proto_path=$INCLUDE \
        --proto_path=$MYDIR \
        --go_out=plugins=grpc:./  \
        $file

    fi
done   

mv ./github.com/dezhishen/onebot-sdk/pkg/model/*.pb.go ./
rm -rf ./github.com