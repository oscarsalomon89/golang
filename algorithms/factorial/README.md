# 🧠 Lógica para Calcular o Fatorial de um Número N

O **fatorial** de um número é o produto de todos os números inteiros positivos de **1 até esse número**.  
É representado com o símbolo **!**.  
Por exemplo:  
`5! = 5 × 4 × 3 × 2 × 1 = 120`

---

## 🔍 Passo a passo da lógica

1. **Receber o número N**

   - O usuário digita um número, por exemplo, `N = 5`.

2. **Verificar se N é 0 ou 1**

   - Por definição:
     - `0! = 1`
     - `1! = 1`
   - Então, se `N` for 0 ou 1, o resultado já é **1**.

3. **Começar com um resultado inicial igual a 1**

   - Criamos uma variável chamada `resultado` e começamos com o valor 1.

4. **Multiplicar todos os números de 2 até N**

   - Por exemplo:
     - Se N = 5, fazemos:
       - `resultado = 1 * 2 = 2`
       - `resultado = 2 * 3 = 6`
       - `resultado = 6 * 4 = 24`
       - `resultado = 24 * 5 = 120`
   - No final, temos o fatorial calculado!

5. **Mostrar o resultado**
   - Depois do cálculo, exibimos o valor final do fatorial.

---

## 🧮 Exemplo com N = 5

Vamos seguir os passos como uma história:

- O número dado é 5
- Começamos com `resultado = 1`
- Multiplicamos por 2 → resultado = 2
- Multiplicamos por 3 → resultado = 6
- Multiplicamos por 4 → resultado = 24
- Multiplicamos por 5 → resultado = 120
- Resultado final: **120**

---

## ⚠️ Observação

Se o número N for **negativo**, o fatorial **não está definido**.  
Nesse caso, o programa deve mostrar uma mensagem de erro, como:

> "Fatorial não está definido para números negativos."

---

## ✅ Dica

Você também pode calcular fatorial de forma **recursiva**, mas neste exemplo usamos a forma **iterativa** (com `for`).
