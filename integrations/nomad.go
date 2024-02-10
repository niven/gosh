package integrations

type NomadJobInfo struct {
	Region     string      `json:"region"`
	Namespace  string      `json:"namespace"`
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Status     string      `json:"status"`
	TaskGroups []TaskGroup `json:"taskgroups"`
}

type TaskGroup struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name string            `json:"name"`
	Env  map[string]string `json:"env"`
}
