# Command

Este patron de comportamiento consiste en convertir todas las peticiones o solicitudes
en objetos independientes que contienen toda la informacion sobre las mismas (una por cada objeto).
Permitiendo asi parametrizar las solicitudes, retrasar, encolar la ejecucion de las mismas
y soportar operaciones que no se puedan realizar.

[Mas informacion.](https://refactoring.guru/es/design-patterns/command)

## En el ejemplo:

Se trabaja con un televisor como dispositivo que tiene que atender las solicitudes
de encendido y apagado. Lo cual se lleva a cabo moviendolo a un objeto independiente que 
se encargue de esa operacion, que puede ser llamado desde cualquier otro lugar.