package main

import "net/http"

type person struct {
	First string
}

func main() {

	// p1 := person{
	// 	First: "Nikhil",
	// }

	// p2 := person{
	// 	First: "Ishaan",
	// }

	// xp := []person{p1, p2}

	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("Print JSON", string(bs))

	// xp2 := []person{}

	// err = json.Unmarshal(bs, &xp2)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("Data in Go data structure ", xp2)

	http.HandleFunc("/encode", encoderFunction)
	http.HandleFunc("/decode", decoderFunction)
	http.ListenAndServe(":8080", nil)

}

func encoderFunction(w http.ResponseWriter, r *http.Request) {

}

func decoderFunction(w http.ResponseWriter, r *http.Request) {

}
