# Visitor

Este es un patron que permite separar algoritmos de los objetos sobre los que operan.

[Mas informacion.](https://refactoring.guru/es/design-patterns/visitor)

## En el ejemplo

El patrón Visitor permite añadir comportamientos a una estructura sin modificar dicha estructura. Digamos que tienes
encargado mantener una biblioteca que tiene distintas estructuras de forma, como:

- Cuadrado
- Círculo
- Triángulo

- Cada una de estas estructuras de forma implementa la interfaz común de forma.

Una vez que la gente de tu empresa empieza a utilizar tu estupenda biblioteca, te llueven las solicitudes de funciones.
Vamos a revisar algunas de las más simples: un equipo te pide que añadas el comportamiento getArea a las estructuras de
forma.

Existen muchas opciones para resolver este problema.

La primera opción que viene a la mente es añadir el método getArea directamente a la interfaz de forma e implementarla
después en cada estructura de forma. Ésta parece una solución a la que recurrir, pero tiene un precio. Como encargado
de mantener la biblioteca, no puedes arriesgarte a descomponer tu precioso código cada vez que alguien pide un nuevo
comportamiento. Sin embargo, quieres que otros equipos amplíen también tu biblioteca.

La segunda opción es que el equipo que solicita la función pueda implementar el comportamiento por sí mismo. No obstante,
esto no siempre es posible, ya que el comportamiento dependerá del código privado.

La tercera opción es resolver el problema anterior utilizando el patrón Visitor. Comenzamos definiendo una interfaz
visitante, así:

```go
    type visitor interface {
    visitForSquare(square)
    visitForCircle(circle)
    visitForTriangle(triangle)
    }
```

Las funciones visitForSquare(square), visitForCircle(circle), visitForTriangle(triangle) nos permitirán añadir
funcionalidades a cuadrados, círculos y triángulos respectivamente.

¿Te preguntas por qué no podemos tener un único método visit(shape) en la interfaz visitante? La razón es que
el lenguaje Go no soporta la sobrecarga de métodos, por lo que no puedes tener métodos con el mismo nombre pero
distintos parámetros.

Ahora, la segunda parte importante es añadir el método accept a la interfaz de forma.

func accept(v visitor)
Toda la estructura de forma tiene que definir este método, de forma parecida a ésta:

```go
    func (obj *square) accept(v visitor){
    v.visitForSquare(obj)
    }
```

Espera un momento, ¿no acabo de decir que no queremos modificar nuestras estructuras de forma existentes?
Lamentablemente, sí, cuando utilizamos el patrón Visitor, tenemos que alterar nuestras estructuras de forma.
Pero esta modificación se realizará solo una vez.

En caso de añadir otros comportamientos, como getNumSides, getMiddleCoordinates, utilizaremos la misma función
__accept(v visitor)__ sin realizar más cambios en las estructuras de forma.

Al final, las estructuras de forma solo tienen que modificarse una vez, y todas las futuras solicitudes
de comportamientos podrán gestionarse utilizando la misma función accept. Si el equipo solicita el comportamiento
__getArea__, podemos limitarnos a definir la implementación concreta de la interfaz visitante y escribir la lógica de
cálculo del área en esa implementación concreta.
