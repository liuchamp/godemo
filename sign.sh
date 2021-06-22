#!/bin/bash

SOURCE="$0"

while [ -h "$SOOURCE"  ]; 
do # resolve $SOURCE until the file is no longer a symlin
	DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"
	SOURCE="$(readlink "$SOURCE")"
	[[ $SOURCE != /*  ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"
echo $DIR
cd $DIR

bin/redis-server rs/redis.conf &