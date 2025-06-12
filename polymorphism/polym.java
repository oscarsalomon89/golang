// Definimos una interfaz "Forma"
interface Forma {
    double area();
    double perimeter();
}

// Implementamos la interfaz en la clase "Cuadrado"
class Cuadrado implements Forma {
    private double lado;

    public Cuadrado(double lado) {
        this.lado = lado;
    }

    @Override
    public double area() {
        return lado * lado;
    }

    @Override
    public double perimeter() {
        return 4 * lado;
    }
}

// Implementamos la interfaz en la clase "Círculo"
class Circulo implements Forma {
    private double radio;

    public Circulo(double radio) {
        this.radio = radio;
    }

    @Override
    public double area() {
        return Math.PI * radio * radio;
    }

    @Override
    public double perimeter() {
        return 2 * Math.PI * radio;
    }
}

// Clase principal
public class PolimorfismoEjemplo {
    public static void imprimirArea(Forma forma) {
        System.out.println("Área: " + forma.area());
    }

    public static void main(String[] args) {
        Forma cuadrado = new Cuadrado(4);
        Forma circulo = new Circulo(5);

        imprimirArea(cuadrado);
        imprimirArea(circulo);
    }
}
