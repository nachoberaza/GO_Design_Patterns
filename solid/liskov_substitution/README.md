# Principio de sustitucion de Liskov

Este principio establece que si una funcion recibe una interfaz y espera un tipo *_T_* quien implementa a esa interfaz, 
cualquier estructura desprendida de _*T*_ que implemente la misma interfaz puede ser usada por esa funcion.

Es decir que si la clase _B_ es una subclase de la clase _A_, se deberia poder pasar un objeto de la clase _B_ a cualquier 
método que espere un objeto de la clase _A_ y el método no debería dar ningún resultado extraño en ese caso.

## En la implementacion

Se tienen dos tipos que implementan la misma interfaz _sizeable_. Por lo que segun el principio de sustitucion de Liskov
ambos pueden ser usados por el metodo _*UseIt()*_, asi como se hace en el metodo main.
