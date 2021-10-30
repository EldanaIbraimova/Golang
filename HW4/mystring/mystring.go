package mystring

import "strings"

func CheckWorker(worker string) bool{
	workers := "Aliya, Eldana, Laura, Aishuak, Arman, Assylzhan"
	return strings.Contains(workers, worker)
}