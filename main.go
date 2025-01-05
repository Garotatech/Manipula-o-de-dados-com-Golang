package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type Produto struct {
	ImageLink string `json:"Image_link"`
	OfferLink string `json:"Offer_link"`
	Price     string `json:"Price"`
}

func main() {
	// Carregar o JSON
    // Semelhante a um Json.load() em Python
	dados := getJSON("data/api_response.json")
	if dados == nil {
		log.Fatalf("Erro ao carregar o JSON.")
	}

	// Obter produtos
	ofertas := getProdutos(dados)

	// Verificar resultado
	if len(ofertas) > 0 {
		fmt.Println("A lista com os resultados foi gerada com sucesso!")
		for _, oferta := range ofertas {
			fmt.Printf("Produto: %+v\n", oferta)
		}
	} else {
		fmt.Println("Nenhuma oferta foi encontrada.")
	}
}

// Lê e retorna os dados do JSON como um mapa
func getJSON(path string) map[string]interface{} {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo JSON: %v", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao ler o conteúdo do arquivo: %v", err)
	}

	var dados map[string]interface{}
	if err := json.Unmarshal(bytes, &dados); err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	return dados
}

// Extrai produtos do JSON usando regex e manipulação de dados
func getProdutos(dados map[string]interface{}) []Produto {
	ofertas := []Produto{}
	produtos := map[string]map[string]interface{}{}
	pattern := regexp.MustCompile(`sp-(\d+)`)

	// Organizar dados por ID do produto
	for chave, valor := range dados {
		if match := pattern.FindStringSubmatch(chave); match != nil {
			produtoID := match[1]
			if _, ok := produtos[produtoID]; !ok {
				produtos[produtoID] = map[string]interface{}{}
			}
			produtos[produtoID][chave] = valor
		}
	}

	// Processar cada produto
	for produtoID, atributos := range produtos {
		var imageLink, offerLink, price string

		for chave, valor := range atributos {
			// Validar os campos necessários
			if chave == "Product:sp-"+produtoID {
				// Extrair descrição
				if valMap, ok := valor.(map[string]interface{}); ok {
					if description, ok := valMap["description"].(string); ok {
						matches := regexp.MustCompile(`<img src="([^"]+)"`).FindStringSubmatch(description)
						if len(matches) > 1 && matches[1] != "" {
							imageLink = matches[1]
							offerLink = imageLink[:len(imageLink)-len("/image.jpg")] // Exemplo de ajuste
						}
					}
				}
			}

			// Obter preço
			if chave == "$Product:sp-"+produtoID+".priceRange.sellingPrice" {
				if valMap, ok := valor.(map[string]interface{}); ok {
					if highPrice, ok := valMap["highPrice"].(float64); ok {
						price = fmt.Sprintf("%.2f", highPrice)
					}
				}
			}
		}

		// Adicionar à lista de ofertas se todos os campos estiverem preenchidos
		if imageLink != "" && offerLink != "" && price != "" {
			ofertas = append(ofertas, Produto{
				ImageLink: imageLink,
				OfferLink: offerLink,
				Price:     price,
			})
		}
	}

	return ofertas
}
