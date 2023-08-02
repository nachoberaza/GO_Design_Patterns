# Principio Open-Close

Este principio establece que una clase debe estar abierta para poder agregar nuevas funcionalidades pero cerrada para la
modificacion del codigo existente de la clase.

## Implementacion

En el codigo se puede ver que se tiene una estructura producto la cual posee un color y un tamaño. Dado un array de
productos, se quiere filtrar por sus propiedades. Esta tarea se delega en un estructura filter que se compone de dos
sub-estructuras que implementan la interfaz _specification_ que posee el metodo _*IsSatisfied*_ el cual determina si
un producto cumple tal especificacion.

En este ejemplo se puede ver que la estructura *ProductFilter* presenta el principio open-close ya que tiene la
capacidad de poder agregar funcionalidades de filtrado a medida que se requiera hacer un filtrado por algun otro
campo del producto. A su vez que la solucion presenta composicion dado que posee como propiedades, estructuras que
implementan las especificaciones.
