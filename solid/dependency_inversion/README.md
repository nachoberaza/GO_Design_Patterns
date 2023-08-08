# Principio de inversion de dependencias

Este principio establece que los módulos de alto nivel no deben de depender de los de bajo nivel. En ambos casos deben 
depender de las abstracciones. Alto nivel se refiere a operaciones cuya naturaleza es más amplia o abarca un contexto 
más general y bajo nivel son componentes individuales más específicos. 

Este principio no tiene que ver necesariamente con el principio de _inversion de dependencias_.

## En la implementacion

Se tienen dos capas en donde se implementan un sistema de personas y relaciones. En la solucion se puede ver como existe
un metodo de bajo nivel para el agregado de parientes, el cual no es independiente de la capa de bajo nivel. Y un manera
de busqueda de parientes desde la capa superior, la cual no depende de la de menor nivel. Respetando asi el principio.
