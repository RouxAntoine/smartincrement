#!/bin/sh

# num=$(pactl list | grep -A 2 'Sink #' | grep -i -B 1 'State: RUNNING' | grep Sink | sed -r 's/Sink #([0-9]*).*/\1/g')
# num=$(pactl list | grep -i 'destination #' | sed -r 's/Destination #([0-9]*).*/\1/g')
num=$( pactl list short sinks | sed -e 's,^\([0-9][0-9]*\)[^0-9].*,\1,' | head -n 1 )
initVal=$(pactl list sinks | grep '^[[:space:]]Volume[^:.]*:' | head -n 1 | sed -e 's,.* \([0-9][0-9]*\)%.*,\1,')

echo "$1"
echo "$num"
echo "$initVal"

if [ "$num" = "" ]
then
    num=0
fi


if [[ "$#" > 0 ]]
then
	# pactl -- set-sink-volume "$num" +2%		exemple with +x%
	if [[ "$1" = "up" || "$1" = "UP" ]]
	then
		pactl set-sink-mute "$num" false
		val=$(/bin/perso/smartincrement --db=/bin/perso/smartincrement.db -inc --init=$initVal --config=/bin/perso/smart.toml)
		pactl -- set-sink-volume "$num" "$val"%
	elif [[ $1 == "down" || $1 == "DOWN" ]]; 
	then
		pactl set-sink-mute $num false
		val=$(/bin/perso/smartincrement --db=/bin/perso/smartincrement.db -dec --init=$initVal --config=/bin/perso/smart.toml)
		pactl -- set-sink-volume $num "$val"%
	fi
else
	echo 'too few argument pass up or down'
fi	
