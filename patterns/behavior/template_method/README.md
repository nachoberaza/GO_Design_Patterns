# Template method

Este es un patron que define un _"esqueleto"_ de un algoritmo en la superclase pero permitiendo que las subclases lo sobrescriban
parcialmente sin cambiar la estructura original.

[Mas informacion.](https://refactoring.guru/es/design-patterns/template-method)

## En el ejemplo

Consideremos el ejemplo de la funcionalidad de la contraseña de un solo uso (OTP, por sus siglas en inglés).
Hay varias formas de entregarle la OTP a un usuario (SMS, correo electrónico, etc.). Pero, independientemente de si se
trata de un SMS o un correo electrónico, el proceso OTP es el mismo:

Generar un número de n dígitos al azar.
Guardar el número en caché para su posterior verificación.
Preparar el contenido.
Enviar la notificación.

Todos los nuevos tipos de OTP que se introduzcan en el futuro, probablemente pasen por todos estos pasos.

De modo que tenemos un escenario en el que los pasos de una operación concreta son los mismos, pero la implementación de
estos pasos puede variar. Ésta es una situación adecuada para considerar utilizar el patrón Template Method.

Primero, definimos un algoritmo plantilla base que consista en un número fijo de métodos. Éste será nuestro método
plantilla. Después, implementaremos cada uno de los métodos de los pasos, pero dejaremos el método plantilla sin variar.
