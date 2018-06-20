package client


func CheckDate(tmp string) string {
	var res string
	if len(tmp) == 1 {
		res = "0" + tmp
	} else {
		res = tmp
	}
	return res
}

