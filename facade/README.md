**Exemplo conceitual**

É fácil subestimar as complexidades que acontecem nos bastidores quando você pede uma pizza com cartão de crédito. Existem dezenas de subsistemas que atuam nesse processo. Aqui está apenas uma lista deles:

  - Checar conta
 - Verificar o PIN de segurança
 - Saldo do crédito/débito
 - Fazer uma entrada no livro-caixa
 - Enviar notificação

Em um sistema complexo como este, é fácil se perder e quebrar coisas se você estiver fazendo algo errado. É por isso que existe o conceito do padrão Facade: uma coisa que permite o cliente trabalhar com dezenas de componentes usando uma interface simples. O cliente só precisa inserir os dados do cartão, o PIN de segurança, o valor a pagar, e o tipo de operação. O Facade direciona comunicações adicionais com vários componentes sem expor o cliente a complexidades internas.

**Fontes:**

https://refactoring.guru/pt-br/design-patterns/facade

https://refactoring.guru/pt-br/design-patterns/facade/go/example

https://golangbyexample.com/facade-design-pattern-in-golang/
