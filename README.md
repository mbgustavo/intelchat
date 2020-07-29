# Intelchat

Sala de chat que permite a comunicação de diversos usuários.

O desenvolvimento foi dividido em um ambiente backend e um frontend. Os ambientes se comunicam através de uma conexão por *web socket*, onde o frontend se conecta em um *pool* de clientes no backend ao entrar com sucesso na sala de chat, facilitando a comunicação em tempo real de mensagens e atualizações na sala. A conexão se mantém ativa enquanto o usuário estiver na sala.

O backend da aplicação foi desenvolvido em Go, utilizando a biblioteca externa gorilla/websocket.

O frontend foi desenvolvido utilizando o framework React, a partir da criação de sua aplicação padrão.

O código foi desenvolvido com nomes de variáveis, comentários e mensagens seguindo padrões em inglês, sendo apenas as mensagens exibidas ao usuário na navegação configuradas em português.

----

## Pré-requisitos

* É necessário possuir Go instalado. O último release foi realizado com a versão 1.14.2.
* É necessário ter node e npm instalados. O último release foi realizado com a versão 12.16.1 do node e 6.13.4 do npm.

----

## Procedimentos

Dentro de `scripts/` há scripts para comandos básicos que permitem realizar procedimento básicos para utilização do sistema. O conteúdo desses arquivos está comentado e pode ser utilizado para caso se deseje aplicar um comando individual. Os scripts devem ser executados dentro da sua pasta com o seguinte comando no bash: `sh <script>.sh`.

* **build**:

Compila o conteúdo do backend em Go, criando e colocando o executável na pasta backend/build juntamente com o arquivo de configurações para produção. Executa também o comando para realizar o build do código do frontend para frontend/build.

* **execute**:

Executa a aplicação caso o build já tenha sido realizado. Esse script apenas executa o código do backend, que deverá estar configurado para server os arquivos de build do frontend. Caso não se deseje manter a opção de servir os arquivos no backend, é possível rodar o frontend separadamente instalando o `serve` e executando o mesmo. Dentro do arquivo há instruções detalhadas de como realizar este procedimento.

* **install**:

Instala as dependências dos projetos. É recomendável rodar esse script antes de realizar o build, e obrigatório caso a instalação de dependências não tenha sido realizada anteriormente.

* **test**:

Executa testes unitários e realiza benchmark. Não é necessário realizar build para rodar esse comando.

----

## Eventos do Web Socket

### Servidor -> Cliente

* **Nome**

  `access`

  **Corpo**

  `{ "message": "Alfred acabou de entrar" }`

  **Descrição**

  Sinaliza entrada de novo usuário na sala

* **Nome**

  `access-result`

  **Corpo**

  `{ "result": true, "reason": "", "users" = [Will, Mike, Dustin] }`
  `{ "result": false, "reason": "room-full", "users" = [] }`

  **Descrição**

  Retorna resultado da tentativa de acesso ao usuário

* **Nome**

  `message`

  **Corpo**

  `{ "nickname": "Alfred", "message": "Boa tarde, meus caros." }`

  **Descrição**

  Mensagem enviada por um usuário repassada para os demais

* **Nome**

  `exit`

* **Corpo**

  `{ message: "Alfred saiu da sala" }`

  **Descrição**

  Sinaliza saída de usuário da sala

  
### Client -> Servidor

* **Nome**

  `access`

  **Corpo**

  `{ "Alfred" }`

  **Descrição**

  Tentativa de acesso de cliente com determinado apelido

* **Nome**

  `message`

  **Corpo**

  `{ "Boa tarde, meus caros." }`

  **Descrição**

  Mensagem enviada por um cliente

----

## Instruções de Uso

* Acesse a interface do chat acessando o endereço e a porta onde o serviço está rodando (Ex: `localhost:4000`)

[!UserAccess](https://user-images.githubusercontent.com/44649580/88800464-94c14500-d17e-11ea-9ab6-903eb9779f7a.png)

* Digite um apelido para entrar no sala de chat. O apelido não pode ser vazio e possui um limite de caracteres configurado em arquivos JSON (16 por padrão) e não podem haver usuários com o mesmo apelido conectados na sala.

* Aperte `Enter` ou o botão `Entrar`na página

* Possíveis erros ao acessar são mostrados logo abaixo do campo de apelido do usuário. Problemas de conexão com o servidor podem ser aferidos abrindo o console do navegador.

* Na tela inicial não há conexão com o websocket, a mesma só é estabelecida após o usuário entrar na sala.

* Dentro da sala de chat você poderá ver um frame principal onde as mensagens aparecem, e ao lado uma lista de usuários. Nessa lista de usuários, o seu apelido aparece destacado. Esse menu é responsivo e não é exibido para larguras de tela menores que 768px.

[!ChatRoom](https://user-images.githubusercontent.com/44649580/88800657-e1a51b80-d17e-11ea-98cc-096c02c3f595.png)

* Abaixo da tela com as mensagens há a entrada de mensagens para o chat que podem ser enviadas apertando `Enter` ou o botão `Enviar`. O botão `Sair` lança uma mensagem de confirmação para desconectar o usuário do chat e voltar para a tela de acesso. O botão de voltar do navegador possui o mesmo comportamento.

* Caso a conexão com o servidor caia, uma mensagem é exibida na tela para o usuário voltar para a tela de acesso.

* Não atualize a página ou sua conexão será perdida.

----

## Maiores no Desenvolvimento da Solução

As maiores dificuldades se concentraram no desenvolvimento do frontend, devido a minha especializaçãõ maior em backend, mas nada impediu de desenvolver uma aplicação funcional, entre as principais dificuldades, estão:

  - Definir a arquitetura do frontend com React, mantendo boas práticas de organização até então não muito aprofundadas;
  - Encontrar o css mais simples e eficaz para atender o planejado para a interface;
  - Simplificar comportamento de acesso, atentando para não mantermos uma conexão com websocket sem necessidade antes do usuário entrar nada sala;
  - Se despendeu um certo tempo tentando aplicar o histórico do chat com `display: flex` e `flex-direction: column-reverse`, pois o firefox, meu navegador padrão, possui um bug nesse tipo de visualização;
  - Dificuldades inerentes de usar uma linguagem e framework com pouca experiência (Go e React), mas que auxiliam muito a pensar fora da caixa para o desenvolvimento e a desenvolver habilidades.

----

## Mais Informações

Detalhes de estrutura e instruções mais detalhadas estão presentes dentro das pastas `frontend/` e `backend/`
