package main

import (
	"fmt"
	"os"

	"github.com/yxxhero/dmidecode"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	dmi, err := dmidecode.New()
	checkError(err)

	// 获取BIOS信息
	infos, err := dmi.BIOS()
	checkError(err)

	for i := range infos {
		fmt.Println(infos[i])
	}

	// 获取所有信息
	infos_all, err := dmi.ALL()
	checkError(err)
	infos_all.Println()
}
