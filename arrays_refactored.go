package main

import "fmt"

func main() {
	// inicia o array e deixa o compilador decidir o tamanho
	DeploymentOptions := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	for index, option := range DeploymentOptions {
		fmt.Println(index, option)
	}
	/*
		0 R-pi
		1 AWS
		2 GCP
		3 Azure
	*/
}
