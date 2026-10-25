# APRENDENDO GO

Aprender desenvolvimento de software na linguagem Go, iniciando pelos conceitos essenciais.

O fluxo básico do processo de compilação em Go:
- CÓDIGO -> COMPILADOR DO CÓDIGO -> LINGUAGEM DE MÁQUINA -> LINKING (LINKAGEM) -> EXECUTÁVEL (PROGRAMA)

Linguagem compilada, multiplataforma (GOOS e GOARCH) - permite compilar o código para diferentes sistemas operacionais 
e arquiteturas (compilação cruzada - cross-compilation), simplificando o processo de CI/CD. 
Porém nem todas as bibliotecas em Go suportam a compilação cruzada, pode ser necessário configurar dependencias 
nativas para plataformas diferentes, necessitando também fazer testes para validar diferentes compilações.

Não existe while em GO, apenar for, mas é possível fazer um for infinito.

### GOOS (Go Operating System): 
Variável de ambiente que especifica qual sistema operacional você está visando para a compilação.
Exemplos de valores comuns incluem:
- linux: Para distribuições Linux;
- windows: Para sistemas Windows;
- darwin: Para sistemas macOS.
### GOARCH (Go Architecture): 
Variável que define a arquitetura do processador alvo.
Exemplos de valores comuns incluem:
- amd64: Para arquitetura x86 de 64 bits;
- 386: Para arquitetura x86 de 32 bits;
- arm: Para processadores ARM, que são comuns em dispositivos móveis e embarcados.

### STRUCTS
No Go, uma struct (ou estrutura) é um tipo de dado composto que agrupa campos 
(também chamados de atributos ou propriedades) sob um único nome. Cada campo 
tem um nome e um tipo. As structs são semelhantes às classes em outras linguagens 
de programação orientadas a objetos, mas o Go não possui classes nem herança. 
Em vez disso, utiliza structs e interfaces para organizar e reutilizar o código.

As structs são úteis quando você deseja representar algo que possui múltiplos 
atributos relacionados. 
Elas permitem uma melhor organização dos dados, facilitando a leitura e 
a manutenção do código. 
- Além disso, ao usar structs, você pode:
- Agrupar dados relacionados sob um nome comum.
- Facilitar a passagem e manipulação de dados em funções.
Usar métodos associados à struct para encapsular comportamentos.

#### Componentes da struct
- *Declaração:* Você usa a palavra-chave type para definir uma nova struct 
e dar um nome a ela. Em seguida, define os campos que pertencem a essa struct, 
com seus respectivos tipos. No exemplo acima, temos a struct Pessoa com três 
campos: Nome, Idade e Cidade.
- *Instanciação:* Para usar uma struct, você precisa inicializá-la. Por exemplo, 
pessoa1 e pessoa2 são instâncias da struct Pessoa. Você pode inicializar 
uma struct especificando os valores de seus campos na ordem correta ou 
explicitamente nomeando os campos.
- *Acesso e Modificação:* Os campos de uma struct podem ser acessados e modificados 
usando a notação . (ponto). No exemplo, pessoa1.Nome acessa o campo Nome da struct pessoa1.

#### Go não tem herança: 
Ao contrário de linguagens como Java ou Python, Go não suporta herança. Em vez disso, 
você pode compor structs ou usar interfaces para compartilhar comportamento entre tipos diferentes.
#### Não há construtores explícitos: 
No Go, você não precisa definir explicitamente um construtor como em outras 
linguagens orientadas a objetos. Em vez disso, você pode criar funções que retornam 
uma nova instância de uma struct, como uma forma de construtor.


## Documentações úteis
https://golang.com.br/aprenda/go-para-iniciantes/
https://golang.com.br/cheatsheet/sintaxe-basica/

## Comandos utilizados

* Para rodar a aplicação

```bash
docker compose up
docker exec container_id go run main.go
```

* Para executar o código diretamente, sem usar o docker:

```bash
go run .
```

## Gerenciamento de dependências

* Para inicializar o módulo:

```bash
go mod init github.com/username/repo-name
```
