# APRENDENDO GO

Aprender desenvolvimento de software na linguagem Go, iniciando pelos conceitos essenciais.

O fluxo básico do processo de compilação em Go:
- CÓDIGO -> COMPILADOR DO CÓDIGO -> LINGUAGEM DE MÁQUINA -> LINKING (LINKAGEM) -> EXECUTÁVEL (PROGRAMA)

Linguagem compilada, multiplataforma (GOOS e GOARCH) - permite compilar o código para diferentes sistemas operacionais e arquiteturas (compilação cruzada - cross-compilation), simplificando o processo de CI/CD. Porém nem todas as bibliotecas em Go suportam a compilação cruzada, pode ser necessário configurar dependencias nativas para plataformas diferentes, necessitando também fazer testes para validar diferentes compilações.

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
go run main.go
```

## Gerenciamento de dependências

* Para inicializar o módulo:

```bash
go mod init github.com/username/repo-name
```
