# Factory

Este patron consiste en que a la hora de construir un objeto, en lugar de utilizar un _new_ se invoque un nuevo
metodo _factory_. El cual internamente llama al *new* y retorna __"productos"__ que son los objetos buscados. 
Esto con el fin de poder abstraer el proceso de creacion de una clase a cualquier subclase que quiera sobreescribir dicho metodo. 

En resumen, es un patron que proporciona una interfaz para crear objetos
en una superclase, mientras que permite a las subclases alterar el tipo de objetos que se crean.

Un factory es cualquier estructura que se encarga de la creacion de un objeto (puede ser una struct dedicada o simplemente
un metodod).

[Mas informacion.](https://refactoring.guru/es/design-patterns/factory-method)