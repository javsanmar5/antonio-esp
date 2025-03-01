# ANTONIO.ESP - Tu lenguaje de programación

Programa en español, y no en cualquier español, sino en el español favorito de tu español favorito. aka: Toño el maestro

## Ejemplo de script

```
queVuelva random deLaMili randint

tiktok MiClase:
    carajote __init__(self):
        self.valor = randint(0, 10)

    carajote mostrar(self):
        print("El valor es:", self.valor)

carajote primera(x):
    porElCulo x eIgua Ji:
        queGuapoEre "Es verdadero"
    porElCuloQue x eIgua QueVa:
        queGuapoEre "Es falso"
    que:
        queGuapoEre "No es ni verdadero ni falso"

carajote otra():
    rojo:
        obj = MiClase()
        obj.mostrar()
        numero = 10
        porElCulo numero eIgua Nanai piola numero < 20:
            print("Número válido")
        que:
            print("Número pequeño")
        queGuapoEre "Todo correcto"
    incompetente ValueError:
        print("Error de valor")
    JAJAJAJA:
        print("Ejecución finalizada")

    contador = 2
    vivaVox contador > 0:
        print("Contador:", contador)
        contador -= 1


primera("VERDE")
otra()
```

## Palabras reservadas
| Python | ANTONIO.ESP|
|----------|-------------|
| False    | QueVa       |
| None     | Nanai       |
| True     | Ji          |
| and      | asopla      |
| break    | perroSanche |
| class    | tiktok      |
| continue | biblio      |
| def      | carajote    |
| del      | agenda2030  |
| elif     | porElCuloQue |
| else     | que         |
| except   | incompetente |
| finally  | JAJAJAJA    |
| for      | unaDuda     |
| from     | queVuelva   |
| if       | porElCulo   |
| import   | deLaMili    |
| in       | porfi       |
| is       | eIgua       |
| not      | mujer       |
| or       | piola       |
| pass     | niCaso      |
| return   | queGuapoEre |
| try      | rojo        |
| while    | vivaVox     |


## Guía de instalación y uso

### Requisitos

Antes de comenzar, asegúrate de tener instalados los siguientes requisitos en tu sistema:

- **Go** (Golang) instalado y configurado en tu PATH.
- **Python** instalado y accesible desde la terminal con el comando `python`.

Puedes verificar la instalación con:

```sh
python --version
go --version
```

### Instalación

Clona el repositorio del proyecto:

```sh
git clone https://github.com/javsanmar5-esp/antonio-esp.git
cd antonio-esp
```

Compila el proyecto con Go:

```sh
go build -o antonio
```

Esto generará un ejecutable llamado `antonio` en el directorio del proyecto.

## Uso

Para ejecutar el programa, usa el siguiente comando:

```sh
./antonio namefile.esp
```

Donde `namefile.esp` es el archivo que deseas procesar.

## Solución de Problemas

- Si el comando `go build` no funciona, asegúrate de que Go está correctamente instalado y configurado.
- Si `python` no es reconocido, intenta instalarlo o usar `python3`.
- Si el ejecutable `antonio` no se ejecuta, verifica los permisos:

  ```sh
  chmod +x antonio
  ```

## Contribuir

Si deseas contribuir al proyecto, puedes hacer un fork del repositorio, realizar cambios y enviar un pull request.

## Licencia
Este proyecto es código abierto y está bajo licencia MIT
