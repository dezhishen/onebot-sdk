MYDIR=C:/github/onebot-sdk/pkg/proto
INCLUDE=C:/Users/11795/go/include
for file in ./*
do
    if [ "${file##*.}"x = "proto"x ]; then
        echo "$file" 
        # echo "protoc --proto_path=$INCLUDE --proto_path=$MYDIR --go_out=plugins=grpc:\model\  $MYDIR/$file"
        protoc --proto_path=$INCLUDE --proto_path=$MYDIR --go_out=plugins=grpc:../model  $MYDIR/$file
    fi
done   