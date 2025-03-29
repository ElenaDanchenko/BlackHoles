package main

import (
	"combinations"
	"fmt"
	"time"
)

// представим граф, как структуру данных типа map: ключи - вершины, значения -  смежные вершины ~ список смежности
type Graph map[int][]int

var bhSize = 0 // глубина черной дыры

// Searcher - сохраняет весь путь, который он прошел до черной дыры.
// всегда в начале - вершина старта, в конце - черная дыра
type Searcher struct {
	l []int // список посещенных вершин
}

// Поисковики (пути), завершившие свой путь в черной дыре
var listOfSearchers []Searcher

type BlackHole struct { // черная дыра
	nodes []int
}

// список найденных черных дыр
var listOfBlackHoles []BlackHole

var nodesCounter = 0

var graphSizeOne = Graph{
	1: {},
}

var graphSizeThreeNotLoop = Graph{
	1: {2, 3},
	2: {3},
	3: {},
}

var graphSizeThreeIsLoop = Graph{
	1: {2},
	2: {3},
	3: {1},
}

var graphSizeTwentyWithoutBlackHole = Graph{
	1:  {3, 5},
	2:  {1, 4},
	3:  {2, 6},
	4:  {1},
	5:  {7, 8},
	6:  {10},
	7:  {8, 9},
	8:  {10},
	9:  {11, 12},
	10: {13},
	11: {14},
	12: {15},
	13: {16},
	14: {17},
	15: {18},
	16: {19},
	17: {20},
	18: {17},
	19: {20},
	20: {18},
}

var graphSizeTwentyWithBlackHole = Graph{
	1:  {},
	2:  {1},
	3:  {1},
	4:  {2, 3},
	5:  {3},
	6:  {3},
	7:  {4},
	8:  {5},
	9:  {6},
	10: {5},
	11: {7, 8},
	12: {9},
	13: {9},
	14: {10},
	15: {10},
	16: {11},
	17: {15},
	18: {14},
	19: {20},
	20: {15},
}

var graphSizeTwentyThreeWithBlackHole = Graph{
	1:  {},
	2:  {1},
	3:  {1},
	4:  {2, 3},
	5:  {3},
	6:  {3},
	7:  {4},
	8:  {5},
	9:  {6},
	10: {5},
	11: {7, 8},
	12: {9},
	13: {9},
	14: {10},
	15: {10},
	16: {11},
	17: {15},
	18: {14},
	19: {20},
	20: {15},
	21: {1},
	22: {1},
	23: {1},
}

var graphByExample = Graph{
	1:  {9},
	2:  {1, 3, 10, 11},
	3:  {},
	4:  {3, 5},
	5:  {6, 11, 12},
	6:  {7, 8, 12},
	7:  {},
	8:  {7, 13},
	9:  {8, 10, 12},
	10: {12},
	11: {4, 12},
	12: {},
	13: {12},
}

var graphSizeHundred = Graph{
	1:   {2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	2:   {12, 13, 14},
	3:   {15, 16, 17},
	4:   {18, 19, 20},
	5:   {21, 22, 23},
	6:   {24, 25, 26},
	7:   {27, 28, 29},
	8:   {30, 31, 32},
	9:   {33, 34, 35},
	10:  {36, 37, 38},
	11:  {39, 40, 41},
	12:  {42, 43, 44},
	13:  {45, 46, 47},
	14:  {48, 49, 50},
	15:  {51, 52, 53},
	16:  {54, 55, 55},
	17:  {57, 58, 59},
	18:  {60, 61, 62},
	19:  {63, 64, 65},
	20:  {66, 67, 68},
	21:  {70, 71, 72},
	22:  {73, 74, 75},
	23:  {76, 77, 78},
	24:  {79, 80, 81},
	25:  {82, 83, 84},
	26:  {85, 86, 87},
	27:  {88, 89, 90},
	28:  {42, 43, 44},
	29:  {45, 46, 47},
	30:  {48, 49, 50},
	31:  {51, 52, 53},
	32:  {54, 55, 55},
	33:  {57, 58, 59},
	34:  {60, 61, 62},
	35:  {63, 64, 65},
	36:  {66, 67, 68},
	37:  {70, 71, 72},
	38:  {73, 74, 75},
	39:  {76, 77, 78},
	40:  {79, 80, 81},
	41:  {82, 83, 84},
	42:  {91},
	43:  {91},
	44:  {91},
	45:  {91},
	46:  {91},
	47:  {91},
	48:  {91},
	49:  {91},
	50:  {91},
	51:  {91},
	52:  {92},
	53:  {92},
	54:  {92},
	55:  {92},
	56:  {92},
	57:  {92},
	58:  {92},
	59:  {92},
	60:  {92},
	61:  {92},
	62:  {93},
	63:  {93},
	64:  {93},
	65:  {93},
	66:  {93},
	67:  {93},
	68:  {93},
	69:  {93},
	70:  {93},
	71:  {93},
	72:  {94},
	73:  {94},
	74:  {94},
	75:  {94},
	76:  {94},
	77:  {94},
	78:  {94},
	79:  {94},
	80:  {94},
	81:  {94},
	82:  {95, 96, 97},
	83:  {95, 97, 98},
	84:  {95, 98, 99},
	85:  {95, 99, 100},
	86:  {95, 96, 97},
	87:  {95, 97, 98},
	88:  {95, 98, 99},
	89:  {95, 99, 100},
	90:  {95},
	91:  {},
	92:  {},
	93:  {},
	94:  {1, 11, 21, 31, 41, 51, 61, 71, 81, 91},
	95:  {2, 12, 22, 32, 42, 52, 62, 72, 82, 92},
	96:  {3, 13, 23, 33, 43, 53, 63, 73, 83, 93},
	97:  {4, 14, 24, 34, 44, 54, 64, 74, 84, 94},
	98:  {5, 15, 25, 35, 45, 55, 65, 75, 85, 95},
	99:  {6, 16, 26, 36, 46, 56, 66, 76, 86, 96},
	100: {7, 17, 27, 37, 47, 57, 67, 77, 87, 97},
}

func contains(arr []int, value int) bool {
	for _, z := range arr {
		if z == value {
			return true
		}
	}
	return false
}

// func containsBH -  проверяет, что “черные дыры” не добавляются повторно
// arr - список BlackHole, value - конкретная дыра
// func containsBH(arr chan []BlackHole, value BlackHole) bool {
func containsBH(arr []BlackHole, value BlackHole) bool {
	for _, z := range arr {
		if len(z.nodes) == len(value.nodes) {
			allelemequal := true
			for i := 0; i < len(z.nodes); i++ {
				allelemequal = allelemequal && z.nodes[i] == value.nodes[i]
			}
			if allelemequal {
				return true
			}
		}
	}
	return false
}

// func CopySearcher - создается копия текущего поисковика, чтобы у каждой вершины был свой поисковик.
func CopySearcher(searcher1 Searcher) Searcher {
	var searcher2 Searcher

	for _, v := range searcher1.l {
		searcher2.l = append(searcher2.l, v)
	}

	return searcher2
}

// func Inject
// добавляет вершину в список посещенных вершин в searcher
// добавляет searcher завершившие свой путь в чд в список listOfSearchers

func Inject(graph Graph, searcher Searcher, v int) {

	searcher.l = append(searcher.l, v) // добавляем в список l пройденных текущим поисковиком вершин текущую вершину

	if len(searcher.l) > bhSize {
		return
	}

	// Проверяем есть ли исходящие ребра у вершины

	if len(graph[v]) == 0 {
		listOfSearchers = append(listOfSearchers, searcher) // если нет ребер, то найдена черная дыра  - сохранить текущий поисковик
	} else { // Исходящие ребра есть

		if len(graph[v]) == 1 {
			if !contains(searcher.l, graph[v][0]) {
				Inject(graph, searcher, graph[v][0])
			}
		} else {
			for _, h := range graph[v] { // проход по смежным вершинам h для данной v
				if !contains(searcher.l, h) { // если поисковик не посещал такую вершину ранее проверяем ее рекурсивно
					Inject(graph, CopySearcher(searcher), h) // отправить копию поисковика в следующую вершину, т.к. для каждого пути свой поисковик
				}
				// не выполнение условия говорит о том что вершина была уже посещена этим поисковиком -> поисковик не важен и не запоминается. с ним не работаем
			}
		}
	}

}

func SearchBlackHoles(graph Graph, size int) {
	nodesCounter = len(graph)
	bhSize = size

	// цикл создает уникального поисковика для каждой вершины графа и запускает поиск черной дыры, начиная с этой вершины
	for key, _ := range graph {
		fmt.Print("Построение путей из вершины ", key, " Осталось вершин: ", nodesCounter, "\n")
		var searcher Searcher // уникальный поисковик создается для каждой вершины
		// Функция Inject переносит поисковика от вершины к вершине, добавляя новую вершину в его список l
		// Ей мы передаем обрабатываемый graph - граф, searcher - новый поисковик, key - текущая вершина, для которой создан поисковик
		Inject(graph, searcher, key)
		nodesCounter--
	}

	fmt.Print("Завершено построение путей\n")

	searchersByEnds := make(map[int][]int) // поисковики, разделенные по их конечным точкам: ключи - конечные точки(чд глубины 1), значения - оставшиеся части путей

	// анализ списка поисковиков listOfSearchers
	// el - очередной поисковик, el.l -  его путь
	for _, el := range listOfSearchers {
		// есть ли коли-во вершин в текущем поисковике равно требуемой глубине
		if len(el.l) == size {
			// предположительно выхода из данного набора узлов нет
			noWayToOut := true
			// перебор всех вершин пути
			for _, node := range el.l {
				// перебор смежных вершин
				for _, toNode := range graph[node] {
					noWayToOut = noWayToOut && contains(el.l, toNode) // если из рассматриваемого узла можно попасть только в узлы el.l
				}
			}

			// если нет пути из рассматриваемого набора узлов el.l наружу, то найдена черная дыра заданного размера

			if noWayToOut {
				var newBlackHole BlackHole
				// список вершин заполняем из содержимого поисковика
				newBlackHole.nodes = el.l
				// проверяем не добавлена ли эта вершина ранее
				if !containsBH(listOfBlackHoles, newBlackHole) {
					listOfBlackHoles = append(listOfBlackHoles, newBlackHole)
				}
			}
			// если длина пути el.l меньше size
		} else if len(el.l) < size { // поиск черных дыр, которые образованы вершинами, не входящими в 1 путь (слабосвзяность), создаются ключи - последние элементы
			// путей и значения - все остальные элементы путей
			for index, node := range el.l { // перечисление всех вершин коротких поисковиков под ключами, которыми заканчиваются их пути el.l[len(el.l)-1] - последний элемент
				if index < len(el.l)-1 && !contains(searchersByEnds[el.l[len(el.l)-1]], node) {
					searchersByEnds[el.l[len(el.l)-1]] = append(searchersByEnds[el.l[len(el.l)-1]], node)
				}
			}
		}
	}

	// завершение анализа более коротких путей
	// endPoint - последняя вершина короткого пути, nodes - остальные вершины
	for endPoint, nodes := range searchersByEnds {
		allVariants := combinations.Combinations(nodes, size-1) // все сочетания элементов, кроме последнего, далее ко всем сочетаниям будет добавлен endPoint
		fmt.Print("Начало анализа ", len(allVariants), " вариантов для узла ", endPoint, "\n")
		for _, variant := range allVariants {
			BHcandidate := variant
			BHcandidate = append(BHcandidate, endPoint)
			noWayToOut := true
			for _, node := range BHcandidate {
				for _, toNode := range graph[node] {
					noWayToOut = noWayToOut && contains(BHcandidate, toNode)
				}
			}

			if noWayToOut {
				var newBlackHole BlackHole
				newBlackHole.nodes = BHcandidate
				if len(newBlackHole.nodes) == size && !containsBH(listOfBlackHoles, newBlackHole) {
					listOfBlackHoles = append(listOfBlackHoles, newBlackHole)
				}
			}

		}
	}

}

func main() {

	size := 2

	start := time.Now()

	//SearchBlackHoles(graphSizeOne, size)
	//SearchBlackHoles(graphSizeThreeNotLoop, size)
	//SearchBlackHoles(graphSizeThreeIsLoop, size)
	//SearchBlackHoles(graphSizeTwentyWithoutBlackHole, size)
	//SearchBlackHoles(graphSizeTwentyWithBlackHole, size)
	//SearchBlackHoles(graphSizeTwentyThreeWithBlackHole, size)
	//SearchBlackHoles(graphByExample, size)
	SearchBlackHoles(graphSizeHundred, size)

	duration := time.Since(start)

	fmt.Println("Функция выполнилась за:", duration)
	fmt.Println("BlackHoles размера", size, listOfBlackHoles)
}
