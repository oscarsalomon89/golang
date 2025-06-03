# 📘 Cálculo simples do Máximo Comum Divisor (MDC)

---

## 🧠 O que é o MDC?

O **Máximo Comum Divisor (MDC)** entre dois números é o **maior número** que divide **exatamente** ambos, sem deixar resto.

### Exemplo:

Para os números 48 e 18:

- Os divisores de 48 são: 1, 2, 3, 4, 6, 8, 12, 16, 24, 48
- Os divisores de 18 são: 1, 2, 3, 6, 9, 18
- O número **6** aparece nas duas listas e é o **maior divisor comum** → então **MDC = 6**

---

## 🔍 Passo a passo da lógica para resolver com programação (sem código)

1. **Receber os dois números inteiros positivos**: vamos chamá-los de `a` e `b`.

2. **Descobrir qual é o menor dos dois**:

   - Porque não faz sentido testar divisores maiores que o menor número.

3. **Criar uma estrutura de repetição**:

   - Que vá de `1` até o menor dos dois números.

4. **Verificar a cada passo**:

   - Se o número atual (`i`) divide exatamente `a` e `b`, ou seja, se o resto da divisão é zero nos dois casos.

5. **Guardar o número que foi divisor comum**:

   - Toda vez que `i` divide `a` e `b`, ele é um **candidato** a ser o MDC.
   - Mas seguimos testando, porque pode existir um número maior que também seja divisor comum.

6. **Ao final do laço**:
   - O **último número** que foi divisor comum será o **maior** entre todos → o MDC.

---

## 🧮 Exemplo com os valores 48 e 18

Vamos simular os passos com os números 48 e 18:

1. O menor número é 18
2. Verificamos todos os números de 1 até 18
3. Testamos um por um:

   - 1 divide os dois → guardamos
   - 2 divide os dois → guardamos
   - 3 divide os dois → guardamos
   - ...
   - 6 divide os dois → guardamos
   - 7 não divide → ignoramos
   - ...
   - 18 só divide ele mesmo → não serve

4. O último número que foi divisor comum foi o **6**  
   → **MDC = 6**

---

## ✅ Conclusão

Essa abordagem é ótima para:

- Aprender a lógica de forma clara e sequencial
- Visualizar o funcionamento do MDC
- Entender a estrutura de repetição e condição de múltiplos

❗ **Importante**: Essa não é a forma mais rápida de resolver, mas é a mais didática para quem está começando.
