# State

Este patron permite a un objeto alterar su comportamiento cuando su estado interno cambia. Simula que el objeto cambia
su _clase_.

## En el ejemplo

Se usa una maquina expendedora que solo tiene un tipo de item y que solo tiene 4 estados:

- hasItem (tieneelArtículo)
- noItem (notieneelArtículo)
- itemRequested (artículoSolicitado)
- hasMoney (tieneDinero)

Y solo tiene 4 operaciones:

- Seleccionar el artículo
- Añadir el artículo
- Insertar el dinero
- Dispensar el artículo

El patrón de diseño State debe utilizarse cuando el objeto puede estar en varios estados diferentes y, dependiendo de la solicitud entrante, el objeto necesite cambiar su estado actual.
En nuestro ejemplo, una máquina expendedora puede encontrarse en muchos estados diferentes y estos estados cambiarán constantemente de uno a otro. Digamos que la máquina expendedora se encuentra en itemRequested. Una vez que se realiza la acción “Insertar dinero”, la máquina pasa al estado hasMoney.
Dependiendo de su estado actual, la máquina se puede comportar de forma diferente ante las mismas solicitudes. Por ejemplo, si un usuario quiere comprar un artículo, la máquina procederá a ello si se encuentra en hasItemState, o la rechazará si está en noItemState.
El código de la máquina expendedora no está contaminada con esta lógica; todo el código dependiente del estado se aloja en implantaciones de estado respectivas.

[Mas informacion.](https://refactoring.guru/es/design-patterns/state)
