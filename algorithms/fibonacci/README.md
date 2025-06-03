# 🧠 Passo a passo: Gerar os primeiros N termos da sequência de Fibonacci em Go

Este guia vai te ajudar a escrever um programa simples em Go que gera os primeiros **N** números da sequência de Fibonacci. Em vez de mostrar o código direto, vamos pensar na lógica por trás dele, passo a passo.

---

## 🔢 O que é a sequência de Fibonacci?

A sequência de Fibonacci começa com os números `0` e `1`.  
A partir daí, cada número seguinte é a **soma dos dois anteriores**.

**Exemplo:**  
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...

---

## 👣 Passo a passo para montar o programa

### 1. Ler o valor de N

Você precisa pedir ao usuário **quantos termos** da sequência ele quer ver.  
Use uma função que leia um número inteiro digitado no terminal.

---

### 2. Inicializar os dois primeiros termos

A sequência sempre começa com `0` e `1`.  
Então, se o usuário pedir pelo menos 1 termo, o primeiro número será 0.  
Se pedir 2 termos, os dois primeiros serão 0 e 1.

---

### 3. Criar uma estrutura para armazenar os números

Use um **slice de inteiros (`[]int`)** para guardar os valores da sequência.  
Esse slice deve ter tamanho `N`, pois você vai preencher ele com os N termos.

---

### 4. Preencher os próximos termos com um laço (loop)

A partir do **índice 2**, cada número é a soma dos dois anteriores:

```plaintext
F(n) = F(n-1) + F(n-2)
```
