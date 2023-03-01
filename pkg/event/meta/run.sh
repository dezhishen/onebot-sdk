MYDIR=`dirname $0`
INCLUDE=$GOPATH/include
for file in ./*
do
    if [ "${file##*.}"x = "proto"x ]; then
        protoc --proto_path=$INCLUDE \
            --proto_path=$MYDIR/../../model \
            --proto_path=$MYDIR \
            --go_out=plugins=grpc:./ \
        $MYDIR/$file
    fi
done   
mv ./github.com/dezhishen/onebot-sdk/pkg/event/*/*.pb.go ./
rm -rf ./github.com