package service

import (
	"client/service/opfunc"
	"flag"
	"log"
	"os"
)

func Run() {
	str := flag.String("str", "", "указание строки для поиска самой длинной подстроки без повторяющихся символов")
	urlReq := flag.String("url", "", "указание url для запроса")
	flag.Parse()
	if *str == "" || *urlReq == "" {
		log.Println("Обязательный флаг не указан")
		flag.Usage()
		os.Exit(1)
	}

	result, err := opfunc.GetLongestSubstring(*str, *urlReq)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Answer:%v -> %v", *str, result)

}
