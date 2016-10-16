# go-command_metrics

go-command_metrics is the instrument of os.exec()

This wraps exec.command (os/exec package) to measure the execution time of outer command.  
Now, this only support Run() and Output() functions.

## Usage

```
import (
	"github.com/tom--bo/go-command_metrics"
	"os/exec"

	   )

func main() {
  // Wrap command
	_cmd := exec.Command("sleep", "1")
	cmd := command_metrics.WrapCommand("test", _cmd)
  // Use as usual
	err := cmd.Run()
	if err != nil {
		// ...

	}

}
```

## Output

```
time:2016-10-16 23:28:50.1035387 +0900 JST	command:test	count:1	sum: 1.001	elapsed(sec): 1.001
time:2016-10-16 23:28:51.107597447 +0900 JST	command:test2	count:1	sum: 1.003	elapsed(sec): 1.003
time:2016-10-16 23:28:52.111428403 +0900 JST	command:test	count:2	sum: 2.005	elapsed(sec): 1.004
time:2016-10-16 23:29:01.115079064 +0900 JST	command:test	count:3	sum:11.008	elapsed(sec): 9.003
```

## Copyright

go-command_metrics is created by @tom\__bo_CS (id:tom\__bo).

See [MIT LICENSE](https://opensource.org/licenses/MIT)


