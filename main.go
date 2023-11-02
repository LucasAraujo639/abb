package main

import (
	"fmt"
	TDAabb "tdas/diccionario"
	TDAPila "tdas/pila"
)

type funcCmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

// Apilar nodos todo hacia la izquierda y si no busco a la derecha y todo hacia la izquierda devuelta
func apilarTodoIzquierda(nodo *nodoAbb[int, string], mipila TDAPila.Pila[nodoAbb[int, string]]) {
	for nodo != nil {
		mipila.Apilar(*nodo)
		nodo = nodo.izquierdo
	}
}
func funcionComparacion(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
func main() {
	mipila := TDAPila.CrearPilaDinamica[nodoAbb[int, string]]()

	arbol := TDAabb.Crearabb[int, string](funcionComparacion)
	arbol.Guardar(5, "nodo raiz")
	arbol.Guardar(3, "Hijo Izquierdo 1")
	arbol.Guardar(2, "Hijo Izquierdo 2")

	iterador := TDAabb.IteradorRango[int, string](10, 39)

	if iterador.HaySiguiente() {
		fmt.Println(iterador.VerActual())
		iterador.Siguiente()
	}
	apilarTodoIzquierda(iterador.abb.raiz, mipila)
	for !mipila.EstaVacia() {
		fmt.Println(mipila.Desapilar())
	}

}
