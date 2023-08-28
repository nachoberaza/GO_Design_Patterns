# Chain of responsability

Este patron se basa en manejar una plantear el sistema como un conjunto de manejadores que
dependiendo de la peticion que se le hace, cada uno de estos manejadores puede decidir si procesar la peticion
y pasarla al siguiente o directamente descartarla.

[Mas informacion.](https://refactoring.guru/es/design-patterns/chain-of-responsibility)

## En el ejemplo:

Cuando llega un paciente, primero pasa por recepción,
después por la consulta del médico y después por la sala de curas,
y por último por la caja (etcétera). El paciente pasa por una cadena de departamentos y cada uno de ellos
envía al paciente un poco más allá en la cadena, una vez que se complete su función.

El paciente pasa primero por la recepción. Después, en base al estado del paciente,	desde recepción lo envían
al siguiente manejador de la cadena.
