package main

import "fmt"

func main() {
	opcion_menu_principal := 0

	menu_principal := ` Menú Principal
	1. Cambiar su contraseña
	2. Editar perfil
	3. Foro
	4. Finalizar programa`

	menu_perfil := `Menú Perfil
	1. Cambiar foto de perfil
	2. Cambiar nombre de usuario
	3. Editar biografía
	4. Volver al menú principal`

	menu_foro := ` Menú Foro
	1. Publicar nueva entrada en el foro
	2. Editar una entrada del foro
	3. Eliminar una entrada del foro
	4. Volver al menú principal`

	for {

		switch opcion_menu_principal {
		case 0:
			println(menu_principal)

			_, err := fmt.Scanf("%d", &opcion_menu_principal)
			if err != nil {
				panic(err)
			}

		case 1:
			println(menu_perfil)

			_, err := fmt.Scanf("%d", &opcion_menu_principal)
			if err != nil {
				panic(err)
			}

		case 2:
			println(menu_foro)

			_, err := fmt.Scanf("%d", &opcion_menu_principal)
			if err != nil {
				panic(err)
			}

		}
	}
}
