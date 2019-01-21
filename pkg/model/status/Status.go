package status

type Status string

var enums []string

func New(str string) Status {

	for _, s := range enums {
		if s == str {
			return Status(s)
		}
	}

	panic("Unknown Status: " + str)
}

func (kind Status) String() string {
	return string(kind)
}

func List() []string {
	return enums
}

func ListGeneric() []interface{} {

	result := make([]interface{}, len(enums))

	for i, s := range enums {
		result[i] = s
	}

	return result
}

func createStatus(s string) Status {

	enums = append(enums, s)

	return Status(s)
}

var (
	Created  = createStatus("created")
	Restarting  = createStatus("restarting")
	Running  = createStatus("running")
	Removing  = createStatus("removing")
	Paused  = createStatus("paused")
	Exited  = createStatus("exited")
	Dead  = createStatus("dead")
)
