///////////////////////////////////массивы "[]"//////////////////срезы карты
/////////////////////////массив это коллекция "однотипных" элементов "фиксированной" длинны
///массив ВСЕГДА фиксированного размера 
package main

import "fmt"

func main() {
    cities := [5]string{"New York", "Paris", "Berlin", "Madrid"} // cities := [...] если инициализ. массив так,то при выводе не 
    // будет пробела в конце
    fmt.Println("Cities:", cities)
    cities[1] = "Volgograd"  //так можем поменять значение в массиве. париж поменяем на волгоград
}
При выполнении приведенного выше кода получается следующий результат:
Cities: [New York Paris Berlin Madrid ]
Хотя массив содержит пять элементов, присваивать значения всем элементам необязательно. Как мы уже видели, последний элемент содержит 
пустую строку, так как это значение используется по умолчанию для строкового типа данных.


//////////////////////////////////////////////////////////////////////////////////////////////////
Еще один интересный способ инициализации массива — использовать многоточие и указать значение 
только для последнего элемента. Например, рассмотрим следующий код:
package main

import "fmt"

func main() {
    numbers := [...]int{99: -1}
    fmt.Println("First Position:", numbers[0])
    fmt.Println("Last Position:", numbers[99])
    fmt.Println("Length:", len(numbers))
}
Выполните этот код, и вы получите следующий результат:
First Position: 0
Last Position: -1
Length: 100

////////////////////////////////////////////Многомерные массивы///////////////////////////////////////////////////////////////////////
Go поддерживает многомерные массивы для работы со сложными структурами данных. Давайте создадим программу, в которой объявляется 
и инициализируется двумерный массив. Используйте следующий код:

package main

import "fmt"

func main() {
    var twoD [3][5]int
    for i := 0; i < 3; i++ {
        for j := 0; j < 5; j++ {
            twoD[i][j] = (i + 1) * (j + 1)
        }
        fmt.Println("Row", i, twoD[i])
    }
    fmt.Println("\nAll at once:", twoD)
}
При выполнении приведенной выше программы вы получите следующий результат:
Row 0 [1 2 3 4 5]
Row 1 [2 4 6 8 10]
Row 2 [3 6 9 12 15]

All at once: [[1 2 3 4 5] [2 4 6 8 10] [3 6 9 12 15]]

Вы объявили двумерный массив, указав, сколько элементов будет находиться во втором измерении, следующим образом:
 var twoD [3][5]int. Этот массив можно рассматривать как структуру данных со столбцами и строками, как матрицу или
  электронную таблицу. Изначально все элементы массива имеют значение по умолчанию (ноль). В цикле for мы инициализируем элементы
   массива с различными значениями в каждой строке. Наконец, мы выведем все значения в терминал.

   //////////////////////////трехмерный массив////////////////////////////
   package main

import "fmt"

func main() {
    var threeD [3][5][2]int
    for i := 0; i < 3; i++ {
        for j := 0; j < 5; j++ {
            for k := 0; k < 2; k++ {
                threeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
            }
        }
    }
    fmt.Println("\nAll at once:", threeD)
}
При выполнении приведенного выше кода получается следующий результат:
All at once: 
[
    [
        [1 2] [2 4] [3 6] [4 8] [5 10]
    ] 
    [
        [2 4] [4 8] [6 12] [8 16] [10 20]
    ] 
    [
        [3 6] [6 12] [9 18] [12 24] [15 30]
    ]
]

//////////////////////////////////////Объявление и инициализация среза////////////////////////////////////

package main

import "fmt"

func main() {
    months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October",
	 "November", "December"}
    fmt.Println(months)
    fmt.Println("Length:", len(months))
    fmt.Println("Capacity:", cap(months))
}

вывод:

[January February March April May June July August September October November December]
Length: 12
Capacity: 12

Обратите внимание, что пока срез не слишком отличается от массива. Срезы и массивы объявляются одинаковым образом. 
Для получения элементов среза можно использовать встроенные функции
 len() и cap(). С помощью этих функций мы подтвердим, что срез может иметь последовательность элементов из базового массива.

 /////////////////////////////////Добавление элементов в срез/////////////////////////////////////

 Для добавления элементов в срез в Go используется встроенная функция append(slice, element).
  Необходимо передать срез, который требуется изменить, и элемент, который нужно добавить в качестве значений функции. 
  Функция append возвращает новый срез, который нужно сохранить в переменной. Если вы изменяете срез, то новый срез можно 
  сохранить в той же переменной.

Давайте посмотрим, как выглядит процесс добавления в коде:

package main

import "fmt"

func main() {
    var numbers []int
    for i := 0; i < 10; i++ {
        numbers = append(numbers, i)
        fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
    }
}
При выполнении приведенного выше кода получается следующий результат:
0       cap=1   [0]
1       cap=2   [0 1]
2       cap=4   [0 1 2]
3       cap=4   [0 1 2 3]
4       cap=8   [0 1 2 3 4]
5       cap=8   [0 1 2 3 4 5]
6       cap=8   [0 1 2 3 4 5 6]
7       cap=8   [0 1 2 3 4 5 6 7]
8       cap=16  [0 1 2 3 4 5 6 7 8]
9       cap=16  [0 1 2 3 4 5 6 7 8 9]
Эти выходные данные интересны. Особенно с точки зрения того, что возвращает функция cap(). До третьей итерации все 
выглядит нормально, затем емкость изменяется на 4, хотя в срезе всего три элемента. В пятой итерации емкость снова 
изменяется на 8, а в девятой — на 16.


/////////////////////////////////////////карты//////////////////////////////////////////////////////////

Карта в Go по сути представляет собой хэш-таблицу, которая является коллекцией пар "ключ-значение". Все ключи в карте 
должны иметь один и тот же тип. То же относится и к значениям. Однако для ключей и значений можно использовать разные типы. 
Например, ключи могут быть числами, а значения — строками. Для доступа к определенному элементу карты используется его ключ.


Для объявления карты необходимо использовать ключевое слово map. Затем определяются типы ключей и значений следующим образом: map[T]T. 
Например, если вы хотите создать карту, содержащую возраст учащихся, можно использовать следующий код:

package main

import "fmt"

func main() {
    studentsAge := map[string]int{
        "john": 32,
        "bob":  31,
    }
    fmt.Println(studentsAge)
}
При выполнении приведенного выше кода получается следующий результат:

map[bob:31 john:32]

Если вы не хотите инициализировать карту с элементами, можно использовать встроенную функцию make() для создания среза, как в предыдущем разделе. 
Чтобы создать пустую карту, можно использовать следующий код:

studentsAge := make(map[string]int)
Карты являются динамическими. После создания карты можно добавлять и удалять элементы карты, а также обращаться к ним.

//////////////////////////добовление элементов в карту////////////////////////////////////////

Чтобы добавить элементы, не нужно использовать встроенную функцию, как в случае со срезами.
Работать с картами проще. Необходимо только определить ключ и значение. Если этой пары "ключ-значение" не существует, 
элемент добавляется в карту.

Давайте перепишем ранее использованный код и создадим карту с помощью функции make. Затем мы добавим элементы в карту. 
Можно использовать следующий код:

package main

import "fmt"

func main() {
    studentsAge := make(map[string]int)
    studentsAge["john"] = 32
    studentsAge["bob"] = 31
    fmt.Println(studentsAge)
}
После выполнения кода вы получите тот же результат, что и ранее:
map[bob:31 john:32]
Обратите внимание, что мы добавили элементы в карту, которая была инициализирована. 
Но если вы попытаетесь сделать то же самое с картой nil, то получите сообщение об ошибке. 


//////////////////////////////Доступ к элементам карты///////////////

Для доступа к элементам карты используется индекс (m[key]), как и для массивов и срезов. 
Ниже приведен простой пример, показывающий, как получить доступ к элементу:

package main

import "fmt"

func main() {
    studentsAge := make(map[string]int)
    studentsAge["john"] = 32
    studentsAge["bob"] = 31
    fmt.Println("Bob's age is", studentsAge["bob"])
}

////////////////////////////////////////////////////Удаление элементов в карте//////////////////////////

Чтобы удалить элемент из карты, используйте встроенную функцию delete(). 
package main

import "fmt"

func main() {
    studentsAge := make(map[string]int)
    studentsAge["john"] = 32
    studentsAge["bob"] = 31
    delete(studentsAge, "john")
    fmt.Println(studentsAge)
}
вывод
map[bob:31]
