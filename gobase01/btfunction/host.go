package btfunction

import (

	"os"
)

func HostNmae() string{

	h_name,_ := os.Hostname()

	return h_name
}


