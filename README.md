## Preparando o ambiente:
O Go deve estar devidamente instalado: `go version`. 

Faça clone do projeto: `git clone https://github.com/stone-payments/card-interview.git -b <seu_nome>`

Ao final, abra um PR para validarmos o trabalho feito.

## Fase 1: Processador de Transações de Cartão de Crédito

### Contexto
Você foi contratado pela Stone Pagamento, para pintar o mundo de verde. A Stone oferece um serviço de cartão de crédito com benefícios exclusivos para seus clientes, como cashback em compras e taxas reduzidas. O sucesso da empresa depende de um sistema de processamento de transações eficiente e confiável.
Seu papel como engenheiro de software é desenvolver um módulo central de processamento de transações. Esse módulo será responsável por validar e registrar as transações realizadas pelos clientes em tempo real. A confiabilidade e a precisão desse sistema são essenciais para evitar fraudes, garantir uma experiência fluida para os usuários e atender às normas do setor financeiro.
A equipe da Stone está enfrentando o seguinte desafio: processar grandes volumes de transações, garantindo que apenas operações válidas sejam registradas e que os clientes sejam notificados imediatamente caso algo dê errado.

#### Requisitos

1. **Entrada da Transação:**
- A transação será representada por uma struct que contenha os dados abaixo

- Exemplo de JSON:


```json
{
    "card_number": "4111111111111111", 
    "amount": 100.50, 
    "currency": "USD", 
    "merchant": "Amazon", 
    "timestamp": "2025-01-17T10:00:00Z" 
}
```

2. **Validações:**

- O timestamp deve ser válido e não pode estar no futuro e deve ser informado do formado RFC-3339.
  - Mensagem de Erro se nao for RFC-3339: `timestamp not valid`
  - Mensagem de Erro se depois de hoje: `timestamp on the future`
- Para qualquer outro erro:
  - Mensagem de Erro: `invalid payload`

3. **Registro da Transação:**

- Use uma estrutura de memória como um slice ou mapa para simular o banco de dados.
- Armazene apenas as transações válidas.

4. **Resposta:**

- Para transações válidas, deve ser informado uma struct com as seguintes informações:
```json
{
	"status": "approved",   
	"authorize_id": "123e4567-e89b-12d3-a456-426614174000" 
}
```
- Para as transações rejeitadas: 
```json
{
	"status": "rejected", 
	"error": "invalid payload"
}
```

## Fase 2: Detecção de Transações Fraudulentas

### Contexto
Parabéns! O sistema que você desenvolveu o autorizador foi um sucesso. Os clientes ficaram satisfeitos, e a Stone conseguiu registrar milhares de transações sem problemas. No entanto, o aumento no volume de operações também chamou a atenção de fraudadores, que tentam explorar o sistema para realizar transações ilícitas.

Agora, seu próximo desafio é aprimorar o módulo de processamento de transações, adicionando um mecanismo de detecção de atividades suspeitas. Esse mecanismo deve identificar possíveis fraudes com base em regras simples e registrar alertas para a equipe de monitoramento.

### Novos Requisitos

1. **Detecção de Transações Suspeitas**
- Implemente uma validação adicional que marque transações como "suspeitas" com base nas seguintes condições:
  - Transações muito altas: Qualquer transação acima de **$10,000** deve ser marcada como suspeita: *high amount*
  - Compras fora do padrão: Se o mesmo número de cartão realizar mais de 5 transações em menos de 1 minuto, as transações adicionais devem ser marcadas como suspeitas. *not standard*

2. **Alteração na Resposta**

- Para transações suspeitas:
  - As transações ainda devem ser aprovadas, mas com um aviso.
  - A resposta deve incluir um campo indicando o status de "suspeita".
  - Exemplo de resposta:
```json
{
  "status": "approved_with_warning",
  "transaction_id": "123e4567-e89b-12d3-a456-426614174000",
  "warning": "transaction marked as suspicious: high amount"
}
``` 

3. **Registro de Alertas**
- As transações que contem risco, devem ser armazenadas em base de dados(slice ou map).

## Fase 3: Processamento em Alta Escala com Concorrência e Otimização de Memória

### Contexto
O sucesso da Stone está se expandindo globalmente, e o volume de transações aumentou exponencialmente. Agora, o sistema precisa processar milhares de transações por segundo de maneira eficiente, sem consumir recursos desnecessários.
A equipe identificou dois grandes desafios:
1. **Baixo desempenho**: O processamento atual não escala bem com o aumento no número de transações simultâneas.
2. **Consumo excessivo de memória**: O uso de estruturas de dados em memória cresceu muito, impactando os custos operacionais.
Sua missão é melhorar o sistema para lidar com transações em alta escala, utilizando os princípios de concorrência e otimização de memória.

### Novos Requisitos

1. **Persistência Otimizada**
- Foi identificado que não precisa armazenar todas as transações, somente as ultimas 10.000 transações.
- Substitua o armazenamento das transações válidas de memória por um buffer cíclico (circular buffer):
  - O buffer deve armazenar até 10.000 transações.
  - Quando o limite for atingido, sobrescreva as transações mais antigas.

2. **Fila de Processamento Assíncrona**
- As transações recebidas pelo sistema devem ser colocadas em uma **fila de trabalho**.
- O processamento deve ser realizado por **workers concorrentes**.
- Você deve permitir configurar o número de workers via uma variável de ambiente (`WORKER_COUNT`).

3. **Limite de Memória**
- A fila de trabalho deve ter um limite de tamanho (default: 1000 transações).
- Quando a fila estiver cheia:
  - O sistema deve retornar um erro ("Too Many Requests") para novas transações.
  - Exemplo de resposta:
```json
{
  "status": "rejected",
  "error": "Too many requests, try again later"
}
```

4. **Monitoramento de Performance**
- Adicione uma função para expor as seguintes métricas:
  - Número de transações processadas.
  - Número de transações rejeitadas (fila cheia).
  - Uso atual da fila (número de transações pendentes).
  - Número de workers ativos. 
Exemplo de resposta:

```json
{
  "processed_transactions": 105000,
  "rejected_transactions": 1500,
  "queue_usage": 450,
  "active_workers": 8
}
```