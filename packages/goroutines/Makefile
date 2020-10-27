#goroutines allow you to work on several tasks at once
#goroutines coordinate work using channels
#channel allow goroutines to send data to each other and sync

#concurrency = allows program to pause one task to work on another
#parallelism = running multiple tasks at the same time - need multi processors to do this
#computer can divide concurrent tasks among different processes to run  at same time

#other program languages do this via multi threading
#go use goroutines
#gorountines require less memory and less time to start up and stop - more goroutines can be run at once
#main is run using a goroutine

#go programs stop running once main goroutine ends
#thus you need to use channels so they can communicate and wait for each other
#go makes no guarentees about when it switches between goroutines or for how long - thus they run when most efficient
#if you need to control this use channels

#cannot use return value on a func thats called with a go statement
#Because you cannot rely on return response to be ready in time to use it

#channels allows you to communicate between goroutines
#means you can ensure recieving goroutine will not use data until sender has sent it
#each channel only has values of a particular type
#var myChannel chan float64 - channels that carries floats
#need to use make to create the channel
#var myChannel chan float64
#myChannel := make(chan float64)
#or myChannel := make(chan float64)
#use <- to send a value on a channel
#myChannel <- 3.14
#also use <- to receive the value fro  a channel
#<- myChannel

#Channels use blocking = pauses all operations in sending or recieving goroutines until data is send/recieve
#allow sync of actions - coordinate timing