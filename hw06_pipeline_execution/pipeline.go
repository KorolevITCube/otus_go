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

	// аналогично костыль, чтобы сохранить промежуточные результаты и
	// в случае преждевременного закрытия пайплайна - выдать пустой канал
	// кажется, что такое поведение все же не целевое и возврат значений, которые пайплайн успел обработать - нормально
	var result []interface{}
	for {
		select {
		case val, ok := <-out:
			if !ok {
				o := make(chan interface{}, len(out))
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
			// откровенно костыль, чтобы вычитать все значения и горутины закрылись
			go func() {
				for v := range out {
					_ = v
				}
			}()
			empty := make(chan interface{})
			close(empty)
			return empty
		}
	}
}
