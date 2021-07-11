# gochannelchange
change the channel used by a goroutine

A go buffered channel cannot be resized but if we can switch to a new channel with a new size.

For example if we want to change dynamically the number of concurrent connections to a resource accessed through a goroutine using a buffered channel.

We start by creating a channel and storing its reference in a handle - a pointer to a pointer.
We pass the handle to the goroutine so the handle content can be changed.

Then we create another channel with the new required size and update the channel handle. We close the first channel to notify the goroutine to stop listening to the first channel and get the reference to the new one.

The main code can ssend data to the new channel, the goroutine will receive the data through the new channel.

Of course, in a real application we should use a channel or a waitgroup to properly wait for the go routine completion instead of the time.Sleep() call at the end of the program.

Note that we could also use a channel to send the channel reference to the goroutine, if using a handle seems too complicated.

