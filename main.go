package main

import (
	"fmt"
)

func main() {
	// prod()
	test()
	// startServer()
}

func loadData() {
	fmt.Println("loading database")
	loadDB()
	fmt.Println(len(db), "entries")

	fmt.Println("preparing vectors")
	loadVectors()
	fmt.Println(len(vectors), "vectors")
}

func prod() {
	loadData()

	fmt.Println("http server ready")
	startServer()
}

func test() {
	// insert("Brasil tem 10.274 casos confirmados de sarampo")
	// insert("Surto de Sarampo no Brasil - Brasil Escola")
	insert("2 perdas")

	loadVectors()

	text := `Mais de 10 mil casos de sarampo foram confirmados no Brasil segundo o novo balanço divulgado nesta quarta-feira, 9, pelo Ministério da Saúde. De acordo com o levantamento, que compreende o período entre o início de 2018 e 8 de janeiro de 2019, foram 10.274 registros confirmados e dois Estados estão com surtos da doença: Amazonas (9.778 casos confirmados) e Roraima (355). Mas, de acordo com a pasta, o número de novos casos está caindo. O ministério informou que uma força-tarefa foi realizada no final do ano passado para concluir casos que estavam em investigação em Manaus, que tinha mais de 7 mil casos nessa situação. "Nas últimas semanas, houve diminuição na notificação de casos novos em Amazonas e em Roraima", diz a pasta.

	Ainda de acordo com o balanço, 12 pessoas morreram por causa da doença em três Estados: seis no Amazonas, quatro em Roraima e duas no Pará.

	O surto ocorreu no Brasil no ano passado e o vírus foi importado da Venezuela, onde a doença circula desde 2017. Casos isolados também foram registrados em outras partes do País, como Pará (61), Rio Grande do Sul (45), Rio de Janeiro (19), Pernambuco (4), Sergipe (4), São Paulo (3), Rondônia (2), Bahia (2) e Distrito Federal (1).

	Vacinação

	Os locais que registraram casos de sarampo receberam, segundo o ministério, 15,5 milhões de doses da vacina tríplice viral (contra sarampo, caxumba e rubéola) entre janeiro do ano passado e janeiro deste ano para intensificar campanhas, realizar vacinação de rotina e para ações de bloqueio. Paula Felix`

	matches := findMatches(text)
	fmt.Println(matches)
}
