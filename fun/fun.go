// 1.все о функциях в GO

package fun 

import(
	"fmt"
)




func NameFunc(a int, b int) {   // NameFunc(входные параметры/аргументы функции) - сначала имя,а потом тип
	fmt.Println(a+b)                           //тело функции - тут пишим любой код,который будет выполняться при вызове данной функии
	 }
func main(){

  NameFunc(22,44)  //вызов функции - просто пишем NameFunc(параметры/аргументы функции; вводим их ) выведет в консоль "66"


}	
	/*  
	func Stroshka(s string) {   
	fmt.Println(s)
	}
	func main(){
		Stroshka("smt1") 
		Stroshka("smt2")
		Stroshka("smt3")     //будет выводить 3 раза строчку
}	
	*/

    /*
	func NameFunc1(a int, b string) {   
		 a = 2    как пример
		 b = "bit" как пример
		fmt.Println(a,b)
		}
		func main(){
			NameFunc1(a,b)  //выведет в консоль "2" и "bit"
		}
	
	*/

 




//==================================================================================================================
//2. Обнуление входных значений переменных

	func Obnulenie(c int) {   //обявили функцию; написали переменную и ее тип
    c = 0                     // объявили значение переменной
	}

	func main(){
		d := 10               //создание новой переменной и присваивание ей значение
		Obnulenie(d)          // при вызове функции Obnulenie с значением переменной d выдаст "10",а не "0"
		fmt.Println(d)  
	}

	//как это обойти:

 	func Obnulenie2(c *int) {   //обявили функцию,на вход она получит переменную с указателем "*"
		*c = 0                    
		}
	
	func main(){
			d := 10               
			Obnulenie2(&d)          // при вызове функции Obnulenie2 с значением переменной d выдаст "0"
			fmt.Println(d)  
		}

	//сложный для понимания пример:

 	func Obnulenie3(g *int) {   
		*g++                   
		}
	func main(){
			d := 10 
			c := &d              
			Obnulenie3(c)          
			fmt.Println(d)  //выведет в консоль "11"
		}

//==================================================================================================================
//3. Безопасный код: определенное и неопределенное поведение:

func devision(a,b int){			//функция с названием "деление"
	if b == 0 {					//проверяем внутри тела функции - не равен ли один из аргументов нулю,чтобы программа 
								//не выдавала ошибку,а выводила сообщение об ошибки
	fmt.Println("неизвестный результат") 	
	}else{
	fmt.Println(a/b)
	}
}

func main(){
	
	devision(6,0) //неизвестный результат
	devision(6,3) //2
	devision(6,2) //2
}

//==================================================================================================================
//4. Функция принимающая значение аргумента как две других функции

func f1(a int) {   
	fmt.Println(a)                  
}
func f2(b int) {   
	fmt.Println(b)
}  

// создаем функцию и аргумент с именем "first" и типом "func",которая на вход принимает число "int"  
func conclude(first func(int), second func(int),a,b int) {   
		   second(b)
		   first(a)               
		}              
func main(){
	conclude(f1,f2,10,20) //должны быть выведены цифры в обратном порядке. так сказано в условиях задания
	}
