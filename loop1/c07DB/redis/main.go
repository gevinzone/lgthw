package main

func main() {
	client, err := Setup("127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	if err = Exec(client); err != nil {
		panic(err)
	}
	if err = Sort(client); err != nil {
		panic(err)
	}
}
