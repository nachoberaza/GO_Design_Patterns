# Strategy

Este es un patron que permite definir una familia de algoritmos, colocar cada uno de ellos en una clase separada y hacer
a sus objetos intarcambiables.

[Mas informacion.](https://refactoring.guru/es/design-patterns/strategy)

## En el ejemplo

Supongamos que estás creando una memoria caché. Debido a que está en la memoria, tiene un tamaño limitado.
Cuando alcanza su tamaño máximo, deben eliminarse algunas entradas para liberar espacio. Esto puede suceder a través de
varios algoritmos. Algunos de los algoritmos populares son:

Menos usada recientemente (LRU, por sus siglas en inglés): elimina una entrada que se ha utilizado menos recientemente.
Primera en entrar, primera en salir (FIFO): elimina una entrada que se creó primero.
Menos frecuentemente usada (LFU): elimina una entrada que se ha utilizado menos frecuentemente.
El problema está en cómo desacoplar nuestras clases caché de estos algoritmos para que podamos cambiar el algoritmo
durante el tiempo de ejecución. Además, la clase caché no debe cambiar cuando se añade un nuevo algoritmo.

Aquí es donde el patrón Strategy entra en escena. Sugiere la creación de una familia del algoritmo en la que cada
algoritmo tiene su propia clase. Cada una de estas clases sigue la misma interfaz, y esto hace que el algoritmo sea
intercambiable dentro de la familia. Digamos que el nombre de la interfaz común es evictionAlgo.

Ahora nuestra clase caché principal integrará la interfaz evictionAlgo. En lugar de implementar todos los tipos de
algoritmos de reemplazo en sí mismos, nuestra clase caché delegará todo ello a la interfaz evictionAlgo. Ya que
evictionAlgo es una interfaz, podemos cambiar el algoritmo durante el tiempo de ejecución a LRU, FIFO, LFU sin realizar
cambios en la clase caché.
