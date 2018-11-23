package main

import (
	"fmt"
	"os"
)

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	fmt.Print(
		" Буду работать и развиваться 24/7     \n" +
			"                                       \n" +
			"                     ,////             \n" +
			"                  ///////////.         \n" +
			"         &&&&& /////////////////       \n" +
			"     .&&&&&&&% /////////////////       \n" +
			"  %&&&&&&&&&&% /////////////////       \n" +
			" &&&&&&&&&&&&% /////////////////       \n" +
			" &&&&&&&&&&&&% /////////////////       \n" +
			" &&&&&&&&&&&&&  ////////////////       \n" +
			" &&&&&&&&&&&&&&&   /////////           \n" +
			" &&&&&&&&&&&&&&&&&&%   //  .,,,,,,,    \n" +
			" &&&&&&&&&&&&&&&&&&&&&&&  ,,,,,,,,,,   \n" +
			" .&&&&&&&&&&&&&&&&&&&&%   ,,,,,,,,,,   \n" +
			"     &&&&&&&&&&&&&&&      ,,,,,,,,,,   \n" +
			"        %&&&&&&&,           .,,,,,     \n" +
			"            *                          \n" +
			"                                       \n")
	a.Run(":8080")
}

// RPC сервер
// Разработать тестовый RPC сервер с одной моделью "Пользователь".
// Требования:
// 1. Модель должна иметь 3 поля: uuid, логин, дата регистрации;
// 2. Модель должна иметь 3 метода: добавить, получить и изменить (по каким полям выборка не имеет значения);
// 3. Для хранения использовать любую удобную бд;
// 4. RPC должен работать как json сервер (json-RPC);
// 5. Код должен быть залит на гит;
// 6. Код должен иметь хотя бы один тест.
