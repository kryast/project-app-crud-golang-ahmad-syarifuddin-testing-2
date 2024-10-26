package service

import (
	"fmt"

	model "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Model"
	utils "github.com/kryast/project-app-crud-golang-ahmad-syarifuddin/Utils"
)

func DeleteSiswa() {
	ReadSiswa()

	if len(model.DataSiswa) == 0 {
		utils.ClearScreen()
		fmt.Println("Data guru masih kosong!")
		BackCrud()
		return
	}

	fmt.Println("\n0. Kembali")
	var id int
	fmt.Print("Masukkan ID siswa yang ingin dihapus: ")
	fmt.Scan(&id)
	if id == 0 {
		utils.ClearScreen()
		Crud()
	} else if id == 99 {
		utils.ClearScreen()
		Home()
	}
	for i, siswa := range model.DataSiswa {
		if siswa.ID == id {
			model.DataSiswa = append(model.DataSiswa[:i], model.DataSiswa[i+1:]...)
			utils.ClearScreen()
			fmt.Println("Data siswa berhasil dihapus!")

			if err := utils.SaveSiswa("Model/data_siswa.json"); err != nil {
				fmt.Println("Error saving siswa data:", err)
			}
			return
		}
	}
	utils.ClearScreen()
	fmt.Println("ID tidak valid!")
}
