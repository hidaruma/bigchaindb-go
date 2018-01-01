package multipipes

import (
	"runtime"
	"os"
	"fmt"
	"context"
	"time"
	"log"
	"os/signal"
	"syscall"
)

const (POISON_PILL string = "POISON_PILL",
	   // DEBUG TODO
	   	)

type Exception error

type PoisonPillException struct {
	*Exception
}

type Queue struct {

}

func Pipe(maxsize int) Queue {
	if maxsize == nil {
		maxsize = 0
	}
	var ctx context.Context
	ctx =
	return Queue(maxsize, ctx)
}

type Target struct {
	Timeout int
}


type Inqueue struct {
}

type Outqueue struct {

}

type Process struct {
}

func (p *Process) IsAlive() bool {

}

func (p *Process) Terminate() {
}

type Node struct {
	Target Target
	Timeout int
	AcceptTimeout bool
	Name string
	Inqueue Inqueue
	Outqueue Outqueue
	ProcessNamespace string
	NumberOfProcesses int
	Processes []Process
	ErrorChannel error
}

func (n *Node) init(target string, inqueue string, outqueue string, name string, timeout string, numberOfProcesses int, maxExecutionTime string, fractionOfCores int)  {
	n.Target = target
	n.Timeout = timeout
	var timeoutParam int
	timeoutParam = signature(n.Target).Parametars["timeout"]
	n.AcceptTimeout = timeoutParam
	n.Name = name
	n.Inqueue = inqueue
	n.Outqueue = outqueue
	n.ProcessNamespace = "pipeline"

	if (numberOfProcesses & numberOfProcesses <= 0) || (fractionOfCores & fractionOfCores <= 0) {
		log.Println(ValueError())
	}
	if fractionOfCores > 0 {
		n.NumberOfProcesses = runtime.NumCPU()  * fractionOfCores
	} else if numberOfProcesses > 0 {
		n.NumberOfProcesses = numberOfProcesses
	} else {
		n.NumberOfProcesses = 1
	}

	for i=0; i < n.NumberOfProcesses; i++ {
		n.Processes = append(n.Processes, Process(n.RunForever))
	}
}

func (n *Node) Log(args os.Args) {
	fmt.Print("[%s] %d> %s", os.Getpid(), n.Name, args[1])
}

func (n *Node) Start(errorChannel error) {
	n.ErrorChannel = errorChannel

	for _, process := range n.Processes {
		process.Start()
	}
}

func (n *Node) SafeRunForever() {
	err := n.RunForever()
	if err != nil {
		switch err.(type) {
		case KeyboardInterrupt:
		case Exception:
			n.ErrorChannel = err
		}
	}
}

func (n *Node) RunForever() {

	setproctitle(n.ProcessNamespace + ":" +n.Name)
	var d time.Time
	d = time.Now().Add(n.timeout * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	go n.Run(ctx)

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	}

}

func (n *Node) Run(ctx context.Context) {
	var arg interface{}
	var poisoned bool
	var timeout bool

	if n.Inqueue != nil {
		arg = n.Inqueue.Get(n.Timeout)
		if args == POISON_PILL {
			poisoned = true

			timeout = true
		}
	}
	var args []string
	switch arg.(type) {
	case []string:
		args = arg
	default:
		args = append(args, fmt.Sprintf("%v", arg))
	}
	var result Target
	if timeout {
		if n.AcceptTimeout
		args = nil
		result = n.Target(args, timeout)
		} else {
			result = nil
		}
	} else {
		result = n.Target(args)
	}

	if result != nil && n.outque != nil {
		for _, item := range result {
			n.Outqueue.Put(item)
		}
	}
	if poisoned {
		log.Println(PoisonPillException())
	}
}

func (n *Node) Join(timeout int) {
	for _, process := range n.Processes {
		process.Join(timeout)
	}
}

func (n *Node) Terminate() {
	for _, process := range n.Processes {
		process.Terminate()
	}
}

func (n *Node) Stop() {
	if n.Inqueue != nil {
		for i:= 0; i < n.NumberOfProcesses; i++ {
			n.Inqueue.Put(POISON_PILL)
		}
	}
}

func (n *Node) IsAlive() bool {
	for _, process := range n.Processes {
		if process.IsAlive() {
			return true
		}
	}
	return false
}

type Item struct {
}

type Pipeline struct {
	Items []Item{}
	Errors []error
	Nodes []*Node
	errorChannel
	RestartOnError bool
	ProcessNamespace string
}

func (pl *Pipeline) init(items []Item,a interface{}, restartOnError bool, processNamespace string) {
	pl.Items = items
	pl.Errors = nil
	pl.ErrorChannel = Pipe()
	if restartOnError != nil {
		pl.RestartOnError = restartOnError
	} else {
		pl.RestartOnError = false
	}
	if ProcessNamespace != nil {
		pl.ProcessNamespace = processNamespace
	} else {
		pl.ProcessNamespace = "pipeline"
	}
	threading.Thread().Start()

	n.Setup()
}

func (pl *Pipeline) Setup(indata , outdata ) {
	var itemsCopy []Item
	itemsCopy = pl.Items
	for _, item := range itemsCopy {
		item.ProcessNamespace = pl.ProcessNamespace
	}

	if indata != nil {
		itemsCopy = append(indata, itemsCopy)
	}
	if outdata != nil {
		itemsCopy = append(itemsCopy, outdata)
	}
}
func (pl *Pipeline) Connect(rest []*Node, pipe ) {
	if rest != nil {
		return pipe
	}
	var head *Node{}
	var tail []*Node{}
	head = rest[0]

}

func (pl *Pipeline) HandleError() {
	var exc error
	exc = n.ErrorChannel.Get()
	n.Errors = append(n.Errors, exc)
	os.Kill(os.Getpid(), syscall.SIGUSR1)
	}
}

func (pl *Pipeline) Restart() {
	pl.Stop()
	pl.Errors = nil
	pl.Start()
}

func (pl *Pipeline) Step() {
	for _, node := range pl.Nodes {
		node.Run()
	}
}

func (pl *Pipeline) Start() {
	for _, node := range pl.Nodes {
		node.Start(pl.ErrorChannel)
	}
}

func (pl *Pipeline) Join() {
	for _, node := range pl.Nodes {
		node.Join()
	}
}

func (pl *Pipeline) Terminate() {
	for _, node := range pl.Nodes {
		node.Terminate()
	}
}

func (pl *Pipeline) Stop(timeout int) {
	if timeout == nil {
		timeout = 30
	}
	for _, node := range pl.Nodes {
		node.Stop()

		err := node.Join(timeout)
		if err != nil {
			node.Terminate()
		}
	}
}

func (pl *Pipeline) IsAlive() bool {
	for _, node := range pl.Nodes {
		if node.IsAlive {
			return true
		}
	}
	return false
}

var LAST_ERROR string = nil

func ExceptionHandler(signum, frame) {

}