package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// 1- Retornar a pasta atual

func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório: ", err)
		return ""
	}
	return dir
}

// 2- Listar arquivos e pastas
func listFilesAndDirectorys() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Erro ao listar arquivos: ", err)
		return
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// 3- Versão do SO
func getOsVersion() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "ver")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("uname", "-a")
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("sw_vers")
	} else {
		fmt.Println("SO não suportado!")
		return
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao executar comando: ", err)
		return
	}
	fmt.Println(string(out))

}

// 4- Configuração da máquina
func getSystemInfo() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "systeminfo")
	} else {
		fmt.Println("O Comando systeminfo não disponível nesse SO!")
		return
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao executar comando: ", err)
		return
	}
	fmt.Println(string(out))

}

// 5- Desligar o computador em 1 hora
func shutdownInOneHour() {
	cmd := exec.Command("shutdown", "/s", "/t", "3600")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao agendar o desligamento! ", err)
	}
}

// 6- Cancelar Desligamento
func cancelShutdown() {
	cmd := exec.Command("shutdown", "/a", "/t", "3600")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao cancelar o desligamento! ", err)
	}
}

func main() {
	fmt.Println("Pasta atual: ", getCurrentDirectory())
	fmt.Println("Arquivos e pastas: ")
	listFilesAndDirectorys()
	fmt.Println("Versão do SO: ")
	getOsVersion()
	fmt.Println("Configuração da máquina: ")
	getSystemInfo()
	// shutdownInOneHour() // Função que deliga o PC
	// cancelShutdown() // Função que cancela o desligamento do PC
}
