# Intelchat - Frontend

Aplicação com interface e comunicação de um cliente para o servidor de chat através de web socket. Foi utilizado framework react na versão 16.13.1.

----

## Procedimentos

Utilize os seguintes comandos dentro dessa pasta para:

* **Instalar dependências**: `npm i`

* **Executar o código em modo de debug**: `npm start`

* **Compilar**: `npm run build`

* **Servir arquivos compilados em determinada porta**: `serve -s build -l <PORTA>`
**É necessário ter o serve instalado, para instalar execute o comando `npm i -g serve`*

----

## Estrutura de arquivos

* **build/**

Pasta gerada ao executar o build, com arquivos estáticos que podem ser servidor pelo backend.

* **node_modules/**

Pasta gerada ao instalar dependências e que contém as mesmas.

* **public/** 

Arquivos públicos padrões do Reack.

* **src/**
  * **src/components/**
    * **src/components/ChatRoom**
    Componentes do React utilizados exclusivamente dentro da sala de chat.
    
    * **src/components/Template**
    Templates de componentes básicos como botões e caixas de texto.
    
    * **src/components/ChatRoom**
    Componentes do React utilizados exclusivamente na tela de acesso.

  * **src/configs/**
    * `configs.json`: arquivo de configurações;
    * `configs.template.json`: template de configurações possíveis com valores zerados;
    * `pt_BR.json`: arquivo com strings em português para exibição na interface.
  
  * **src/fonts/**

    Fontes externas utilizadas na aplicação.

  * **src/main/**
   
    Arquivos do frame principal da aplicação e de rotas.

  * **src/shared/**

    Funções e variáveis compartilhadas por diferentes classes. Contém o socket que se conecta ao servidor, sua função de conexão e erro eventos esperados.

  * **src/\*.js**

  Arquivos padrão para inicialização do React.

  * **src/index.css**

  Arquivo com estilos principais da aplicação. Contém fontes personalizadas e variáveis de cores utilizadas em diversos arquivos `.css`.

* **/\*.json**

Arquivos com dados da aplicação, scripts e dependências.

----

## Componentes

* **ChatRoom/AccessForm**

Formulário de acesso de novo usuário com entrada para nome e botão para conectar.

* **ChatRoom/UserAccess**

Tela de acesso de usuário que contém estados que lidam com as alterações do formulário, lida com mensagens de erro e envia tentativa de conexão com servidor.

* **Templates/Header**

Cabeçalho da aplicação. É exibido em todas as telas e posição do logo (nome da aplicação) possui responsividade.

* **Templates/Button**

Botão que recebe um handler para seu clique. Parâmetro off troca sua cor de fundo.

* **Templates/Input**

Input de texto. Recebe parâmetros de handlers para tecla pressionada (geralmente para tecla `Enter`) e alterações no seu valor, placeholder, tamanho máximo e um estado que define o seu valor.

* **ChatRoom/ChatHistory**

Histórico de mensagens do chat que faz mapeamente de mensagens e possui referência do fim das mesmasm.

* **ChatRoom/ChatInput**

Entrada de mensagens e botão de sair, similar ao formulário de acesso.

* **ChatRoom/ChatRoom**

Sala de chat que contém os frames de histórico, lista de usuários e input. Configura os handlers para mensagens do socket.

* **ChatRoom/Message**

Mensagem formatada para exibição no histórico.

* **ChatRoom/User**

Usuário listado formatado para exibição na lista.

* **ChatRoom/UsersList**

Lista de usuários. Essa lista possui responsividade e não é exibida em telas com largura menor que 768px.

----

## Rotas

* **/**

Tela de acesso de usuário, onde é possível digitar um apelido e entrar na sala de chat. Todas as rotas não configuradas são redirecionadas para essa rota, e cada vez que a mesma é acessada, é feita uma verificação para garantir que não há conexão com o web socket.

**Essa é a única rota acessível diretamente quando os arquivos de build estiverem servidos no backend*

* **/chat-room**

Sala de chat que é acessada após ter um resultado de acesso positivo recebido so servidor. É necessário ter uma conexão ativa com o web socket e, caso não haja, é feito um redirecionamento para a sala de acesso. O redirecionamento também ocorre em tempo real que o web socket perder sua conexão com o servidor. Ao sair da sala, a conexão com o websocket é fechada.

----

## Configurações

Configurações presentes nos arquivos configs.json.

* **serverAddress**: endreço do servidor de chat (string);
* **serverPort**: porta do servidor de chat (string);
* **maxMsgLength**: tamanho máximo da mensagem enviada no chat (inteiro);
* **maxNickLength**: tamanho máximo do apelido (inteiro).