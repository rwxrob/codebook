# Experimenting with Go Contexts

1. Make sure to pass `context.WithCancel` correctly
1. Make sure to deal with `<-ctx.Done()`
1. Make sure not to block out `<-ctx.Done()`
1. Confirm that `cancel` will affect *all* goroutines
