package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	controlledIn := make(chan interface{})

	go func() {
		defer close(controlledIn)
		for data := range in {
			select {
			case <-done:
				return
			case controlledIn <- data:
			}
		}
	}()

	out := In(controlledIn)
	for _, stage := range stages {
		out = stage(out)
	}

	var result []interface{}
	for {
		select {
		case val, ok := <-out:
			if !ok {
				o := make(chan interface{})
				go func() {
					for _, v := range result {
						o <- v
					}
					close(o)
				}()
				return o
			}
			result = append(result, val)
		case <-done:
			empty := make(chan interface{})
			close(empty)
			return empty
		}
	}
}
