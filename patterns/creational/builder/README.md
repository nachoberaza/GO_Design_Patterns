# Builder

El patron builder consiste en quitar todo el codigo que tenga que ver con la construccion
de un objeto, de su propia clase y colocarlo dentro de un objeto independiente que hara de
_*"constructor"*_. 

Este patron organiza la construccion de objetos en una serie de pasos ejecutados por el o los objetos constructor. 
Estos pasos pueden no ser todos los que se ejecutan dependiendo la configuracion del objeto final que se quiera. 

## Clase directora

Es una clase encargada simplemente de coordinar los objetos constructores. Es decir, darle un orden en particular o configuraciones particulares a los objetos resultantes.

Esta clase no es siempre necesaria pero si permite abstraer mucho mas el proceso de construccion.

[Mas Informacion.](https://refactoring.guru/es/design-patterns/builder)