if [[ $(basename "$PWD") == "sh" ]]; then
    cd ..
elif [[ $(basename "$PWD") != "poker" ]]; then
    echo "please run this script in toktik project root directory"
    exit 1
fi

protoc -I ./proto ./proto/*.proto --go_out=./pkg --go-grpc_out=./pkg