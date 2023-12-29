package route

import (
	"server/route/math"
	"server/route/todolist"
)

var RootRoute = map[string][]any{
	"math.fib":         {math.Fib},
	"todolist.add":     {todolist.Add},
	"todolist.bulkAdd": {todolist.BulkAdd},
	"todolist.get":     {todolist.Get},
	"todolist.getAll":  {todolist.GetAll},
	"todolist.update":  {todolist.Update},
}
