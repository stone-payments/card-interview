# Desafio de Desenho de Software: Módulo de Processamento de Transações

## 1. Contexto

Você foi contratado pela Stone Pagamentos para ajudar a pintar o mundo de verde. A Stone oferece um serviço de cartão de crédito com benefícios exclusivos para seus clientes e taxas reduzidas. O sucesso da empresa depende de um sistema de processamento de transações eficiente e confiável.

Seu papel como engenheiro de software é desenvolver um módulo central de processamento de transações. Esse módulo será responsável por validar e registrar as transações realizadas pelos clientes em tempo real. A confiabilidade e a precisão desse sistema são essenciais para evitar fraudes, garantir uma experiência fluida para os usuários e atender às normas do setor financeiro.

## 2. O Desafio Técnico

A equipe da Stone está enfrentando o seguinte desafio: processar um grande volume de transações, garantindo que apenas operações válidas sejam registradas e que os clientes sejam notificados imediatamente caso algo dê errado.

Para que uma transação seja aprovada, ela precisa passar pela validação de dois sistemas externos independentes:

1. **Serviço de Análise de Risco:** Avalia a probabilidade de a transação ser fraudulenta.
    
2. **Serviço de Limite de Crédito:** Verifica se o cliente possui saldo suficiente para a compra.
    

### Requisitos e Restrições

#### Requisitos Funcionais:

- O sistema deve expor uma API para receber requisições de transação.
    
- Para cada transação, o sistema deve consultar **ambos** os serviços (Análise de Risco e Limite de Crédito).
    
- A transação só será aprovada se **ambas** as validações externas retornarem sucesso.
    
- O sistema deve armazenar a requisição de entrada, as respostas de cada integração e o resultado final (aprovado/recusado) para fins de auditoria.
    

#### Requisitos Não-Funcionais (Constraints):

- **Performance:** A requisição completa, desde o recebimento até a resposta ao cliente, precisa ser resolvida em no máximo **1500 milissegundos (ms)**.
    
- **Latência das Integrações:** Cada uma das integrações externas (Risco e Limite) pode levar até **800ms** para responder.
    
- **Resiliência:** O sistema deve ser resiliente a falhas e timeouts das integrações externas.
    
- **Auditabilidade:** O mecanismo de armazenamento de dados para auditoria não deve comprometer o tempo de resposta da transação.
    
- **Escalabilidade:** A solução deve ser projetada para suportar um alto volume de requisições concorrentes.