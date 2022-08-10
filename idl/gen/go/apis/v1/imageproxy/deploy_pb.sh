#!/bin/bash

# set path and global
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE"
done
readonly DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null && pwd )"
readonly PROGRAM=$(basename $0)
readonly ARGS="$@"

usage() {
  cat <<- EOF
usage: $PROGRAM options

    Compile given proto file with protoc
    It generate
        {FILE}.pb.go
        {FILE}.pb.gw.go
        {FILE}.proto
        {FILE}_grpc.pb.go
    and git commit && push on own github repository

    OPTIONS:
       -f                 proto file containing the RPC rules.
       -h                 show this help


    Examples:
       Run:
       $PROGRAM -f my_proto.proto
EOF
}

cmdline() {
    local arg=
    for arg
    do
        local delim=""
        case "$arg" in
            #translate --gnu-long-options to -g (short options)
            --file)           args="${args}-f ";;
            --help)           args="${args}-h ";;
            --debug)          args="${args}-x ";;
            #pass through anything else
            *) [[ "${arg:0:1}" == "-" ]] || delim="\""
                args="${args}${delim}${arg}${delim} ";;
        esac
    done

    #Reset the positional parameters to the short options
    eval set -- $args

    while getopts "hx:f:" OPTION
    do
         case $OPTION in
         h)
             usage
             exit 0
             ;;
         x)
             echo "DEBUG MODE ON" 
             set -x
             ;;
         f)
             readonly FILE=$OPTARG
             ;;
        esac
    done

    if [[ ! -f $FILE ]]; then
        echo "You must provide --file proto file"
		exit 1
    fi
    return 0
}

main() {
	cmdline $ARGS
	NAME=$(basename $FILE)
	FILENAME=${NAME%.*}
	EXT=${NAME##*.}
	if [[ "$EXT" != "proto" ]]; then
		echo "file must be proto"
		exit 1
	fi
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false $FILE
	protoc --grpc-gateway_out ./ --grpc-gateway_opt paths=source_relative --grpc-gateway_opt allow_repeated_fields_in_body=true $FILE

	git commit ${FILENAME}* -m"auto deploy $$"
	git push
}

main
