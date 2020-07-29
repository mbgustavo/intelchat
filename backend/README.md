# Intelchat - Backend

Aplicação que controla as conexões por web socket em um pool de clientes e mantém a comunicação livre de condições de corrida. A aplicação também pode servir os arquivos de distribuição do frontend e executar a aplicação toda por si só.

Para realizar o upgrade de uma requisição http para uma conexão web socket, foi utilizada a biblioteca gorilla/websocket (https://github.com/gorilla/websocket). Uma biblioteca amplamente utilizada e com boas referências de utilização.

A aplicação possui cobertura por testes unitários.

----

## Procedimentos

Utilize os seguintes comandos dentro dessa pasta para:

* **Instalar dependências**: `go mod download`

* **Executar o código**: `go run main.go`

* **Compilar**: `go build`

* **Executar o código compilado**: Abara o executável gerado na instalação e tenha certeza de ter o arquivo `configs/configs.json` no mesmo local de onde for feita a execução.

* **Rodar testes**: `go test ./...`
**Logs de warning aparecem nos testes devido a situações de erro forçadas executadas, que em execução normal devem ser barradas no frontend*

* **Rodar testes com benchmarking**: `go test ./... -bench *`

* **Ter acesso a documentação de pacotes gerada automaticamente**:
  * Tenha certeza de que o seu repositório local está dentro do workspace do Go;
  * Tenha o godoc instalado no seu dispositivo;
  * Execute godoc -http=:<_PORT_>;
  * Acesse localhost:<_PORT_>.
  * Alternativamento, procure por esse repositório em godoc.org

----

## Estrutura de arquivos

* **build/**:
Pasta gerada ao executar o build pelos scripts fornecidos, onde haverá um executável e uma pasta config com uma cópia do arquivo em `configs/configs.prod.json`.

* **configs/**:
  * `configs.json`: arquivo efetivamente lido na execução do código, recomendável manter configurações que facilitem a depuração;
  * `configs.prod.json`: arquivo com configurações de produção, deve ser utilizado com o executável compilado ao executar o script de build;
  * `configs.template.json`: template de configurações possíveis com valores zerados.

Detalhes sobre as configurações estão descritos na sessão de configurações aqui presente.

* **log/**:

Estando ativada a configuração `logFiles`, arquivos de logs são geradas e colocados nessa pasta, separados por dia.

* **pkg/**
  * **pkg/logger/**:

  Pacote utilizado para logar mensagens que servem tanto para registro como depuração. Nesse pacote é possível diferenciar logs de mensagens normais de erros e *warnings*. Erros são logados para comportamento inesperados da própria aplicação, enquanto warnings são indicativos de problemas de protocolo com alguma cliente, por exemplo. Há também uma função para logs de debug, que apenas aparecem caso a configuração `verboseLogs` esteja ativa. Caso a configuração `logFiles` esteja desativada, os logs são impressos na saída padrão do sistema operacional.

  * **pkg/websocket/**:
      * `client.go`: loop de conexão do cliente, que segue aguardando mensagens do cliente para lidar com as mesmas de forma adequada;
      * `handlers.go`: handlers para diferentes eventos que chegam nas mensagens de um cliente;
      * `pool.go`: controle do pool de conexões, adiciona e remove clientes do pool e realiza broadcast de mensagens em uma thread dedicada, mantendo o acesso ao sockets `thread*safe`
      * `types.go`: tipos e constantes utilizados no pacote;
      * `validator.go`: validadores para clientes e pool, prevenindo possíveis erros e mantendo as regras e limites da aplicação;
      * `websocket.go`: onde se realiza o upgrade para websocket e inicia o loop de leitura dos clientes.

* **main.go**:

Rotina de entrada do sistema

* **go.mod/go.sum**:

Arquivos com informações de pacotes externos.

----

## Configurações

Configurações presentes nos arquivos configs.json. É sempre necessário haver um arquivo de configuração onde a aplicação for executada.

* **serverPort**: porta de execução do aplicação (string);
* **webSocket**: configurações websocket;
* **webSocket/readBufferSize**: tamanho do buffer de leitura do websocket (inteiro);
* **webSocket/writeBufferSize**: tamanho do buffer de escrita do websocket (inteiro);
* **logFiles**: ativa log em arquivos em vez de usar a saída padrão do sistema (booleano);
* **verboseLogs**: ativa mais logs para auxiliar na depuração (booleano);
* **maxClients**: capacidade de clientes no pool de conexões (inteiro);
* **maxNickLength**: tamanho máximo de apelido (inteiro);
* **serveFiles**: serve arquivos de front (booleano);
* **staticFilesPath**: path dos arquivos estáticos do front a serem servidos (string).