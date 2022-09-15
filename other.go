package main

type str struct {
	Result string
	Imei   string
}

//func main() {
//	//key := "E546C8DF278CD5931069B522E695D4F2"
//	//print("Enter a string: ")
//	//var str string
//	//_, _ = fmt.Scanln(&str)
//	//c, _ := aes.NewCipher([]byte(key))
//	//out := make([]byte, len(str))
//	//c.Encrypt(out, c)
//	jsonFile, _ := os.Open("test.json")
//	decoder := json.NewDecoder(jsonFile)
//	var items []str
//	_ = decoder.Decode(&items)
//	for _, item := range items {
//		if item.Result != "Başarılı" {
//			println("\"" + item.Imei + "\",")
//		}
//	}
//}
