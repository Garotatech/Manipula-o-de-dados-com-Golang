# Extrator de Ofertas em Go

Este projeto tem como objetivo carregar um arquivo JSON com dados de produtos, extrair informações relevantes usando regex e manipulação de dados, e apresentar uma lista de ofertas. É uma implementação em Go que busca organizar e estruturar dados de produtos.

---

## 📚 Sobre o Projeto

O projeto realiza as seguintes tarefas:
- Carrega um arquivo JSON de dados brutos.
- Filtra os dados relevantes utilizando expressões regulares.
- Extrai informações essenciais de cada produto, incluindo:
  - **ImageLink**: Link da imagem do produto.
  - **OfferLink**: Link da oferta do produto.
  - **Price**: Preço do produto.
- Retorna uma lista estruturada das ofertas encontradas.

---

## 🚀 Tecnologias Utilizadas

- **Linguagem de Programação**: Go (Golang)
- **Pacotes Nativos**:
  - `encoding/json`: Para manipulação de JSON.
  - `io/ioutil`: Para leitura de arquivos.
  - `os`: Para manipulação de arquivos.
  - `regexp`: Para utilização de expressões regulares.
  - `fmt` e `log`: Para saída e logging de informações.

---

## ⚙️ Instalação e Configuração

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/seu-usuario/extrator-ofertas-go.git
   ```

2. **Navegue até o diretório do projeto**:
   ```bash
   cd extrator-ofertas-go
   ```

3. **Certifique-se de ter o Go instalado**.
   - Para verificar:
     ```bash
     go version
     ```

4. **Prepare o arquivo JSON**:
   - Certifique-se de que o arquivo `api_response.json` esteja no diretório `data/`.

5. **Execute o programa**:
   ```bash
   go run main.go
   ```

---

## 📂 Estrutura do Projeto

```
extrator-ofertas-go/
├── data/                     # Diretório de dados
│   └── api_response.json     # Arquivo JSON com os dados de entrada
├── main.go                   # Código principal
├── README.md                 # Documentação do projeto
```

---

## 📝 Como Funciona o Código

1. **Função `getJSON`**:
   - Lê o arquivo JSON no caminho especificado e o decodifica em um mapa de strings para interfaces.

2. **Função `getProdutos`**:
   - Utiliza regex para identificar chaves relacionadas a produtos.
   - Organiza os dados em um mapa de IDs de produtos para atributos.
   - Extrai links de imagem, links de oferta e preços, validando cada campo.
   - Retorna uma lista de estruturas `Produto` com os dados extraídos.

3. **Estrutura `Produto`**:
   - Define os campos extraídos do JSON:
     - `ImageLink`
     - `OfferLink`
     - `Price`

4. **Execução no `main`**:
   - Carrega o JSON com `getJSON`.
   - Processa os dados com `getProdutos`.
   - Exibe os resultados ou informa que nenhuma oferta foi encontrada.

---

## 📦 Exemplo de Saída

Ao executar o programa com um JSON válido, você verá a saída semelhante a:

```
A lista com os resultados foi gerada com sucesso!
Produto: {ImageLink:https://example.com/image.jpg OfferLink:https://example.com Price:99.99}
Produto: {ImageLink:https://example.com/other.jpg OfferLink:https://example.com/offer Price:149.99}
```

---

## 🛠️ Contribuindo

Contribuições são bem-vindas! Para contribuir:

1. Faça um fork do repositório.
2. Crie uma branch para sua feature:
   ```bash
   git checkout -b minha-feature
   ```
3. Commit suas mudanças:
   ```bash
   git commit -m "Adicionei uma nova feature"
   ```
4. Faça o push para a branch:
   ```bash
   git push origin minha-feature
   ```
5. Abra um Pull Request.

---

## 📄 Licença

Este projeto está sob a licença [MIT](LICENSE). Sinta-se à vontade para usar e modificar.

---

## 💬 Contato

- **GitHub**: [seu-usuario](https://github.com/seu-usuario)
- **Email**: seu-email@example.com

