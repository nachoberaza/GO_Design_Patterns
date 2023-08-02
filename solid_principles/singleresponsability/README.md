# Principio de responsabilidad unica

Este principio dice que una clase debe tener una y solo una responsabilidad.

## Implementacion

En el ejemplo propuesto se tiene una estructura journal en donde se delega toda responsabilidad de persistencia a la estructura persistence.

Es decir que la estructura "persistence" tiene la unica responsabilidad de _persistir_ a journal.
