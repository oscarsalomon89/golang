# 🧠 Lógica para Somar Números Pares até N

Imagine que você recebe um número qualquer, por exemplo, **N = 10**, e o seu objetivo é somar **todos os números pares** de 1 até esse número.  
Os números pares são aqueles que **dividem por 2 e o resto dá 0** (ou seja, múltiplos de 2): 2, 4, 6, 8, 10...

---

## 🔍 Passo a passo da lógica

1. **Receber o número N**

   - Alguém te dá um número final. É o limite até onde você pode contar.
   - Exemplo: N = 10.

2. **Começar do número 2**

   - Sabemos que o primeiro número par depois de 1 é o 2. Então a nossa contagem começa ali.

3. **Ir somando os pares um por um até chegar no N**

   - A ideia é:
     - Começar em 2
     - Depois ir para 4
     - Depois 6
     - Depois 8
     - E depois 10 (se N for 10)
   - A cada passo, somamos o número na nossa variável de **soma total**.

4. **Parar quando passar do N**

   - Assim que o número que estamos somando for maior do que N, a contagem para.

5. **Mostrar o resultado da soma**
   - No final, mostramos a soma total dos pares que encontramos nesse intervalo.

---

## 🧮 Exemplo com N = 10

Vamos seguir os passos como uma história:

- O número dado é 10.
- Começamos a somar do 2.
- Somamos 2 → soma = 2
- Somamos 4 → soma = 6
- Somamos 6 → soma = 12
- Somamos 8 → soma = 20
- Somamos 10 → soma = 30
- O próximo seria 12, mas já passou de 10, então paramos.
- Resultado final: **30**

---

## 💡 Dica

Você poderia fazer isso de outra forma também, por exemplo, verificando se cada número entre 1 e N é par, mas como queremos apenas pares, ir de 2 em 2 já é mais direto e eficiente!
