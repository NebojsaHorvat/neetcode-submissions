type Solution struct{}
// 1 byte size | size bytes data | 1 byte size | ...
func (s *Solution) Encode(strs []string) string {
	ret := make([]byte,0,len(strs)*50)
	
	for _, str := range strs{
		strLen := byte(len(str))
		data := []byte(str)

		ret = append(ret, strLen)
		ret = append(ret, data...)
	}
	return string(ret)
}

func (s *Solution) Decode(encoded string) []string {
	data := []byte(encoded)
	ret := make([]string,0)
	for i:=0; i<len(data); {
		strLen := int(data[i])
		strData := string(data[i+1 : i+1+strLen])
		// fmt.Println(strLen)
		// fmt.Println(strData)
		i= i+1+strLen

		ret = append(ret,strData)
	}
	return ret
}
