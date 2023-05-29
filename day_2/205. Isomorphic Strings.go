package day_2

func IsIsomorphic(s string, t string) bool {
	cstrs := make(map[string]string)
	tstrs := make(map[string]string)

	for i := 0; i < len(s); i++ {
		chs := string(s[i])
		cht := string(t[i])
		if cstrs[chs] == "" && tstrs[cht] == "" {
			cstrs[chs] = cht
			tstrs[cht] = chs
		} else if cstrs[chs] != cht && tstrs[cht] != chs {
			return false
		}
	}
	return true
}
